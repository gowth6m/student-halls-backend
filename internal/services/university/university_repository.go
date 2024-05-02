package university

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"student-halls.com/internal/db"
)

type UniversityRepository interface {
	CreateUniversity(c context.Context, university University) (University, error)
	GetAllUniversities(c context.Context) ([]University, error)
	GetUniversityByID(c context.Context, id string) (*University, error)
}

type repositoryImpl struct {
	collection *mongo.Collection
}

func NewUniversityRepository() UniversityRepository {
	collection := db.Client.Database(db.DATABASE_NAME).Collection(db.COLLECTION_UNIVERSITY)
	return &repositoryImpl{collection: collection}
}

func (r *repositoryImpl) CreateUniversity(c context.Context, university University) (University, error) {
	filter := bson.M{"$or": []bson.M{{"name": university.Name}, {"location": university.Location}}}

	var existingUniversity University
	err := r.collection.FindOne(c, filter).Decode(&existingUniversity)
	if err == nil {
		return University{}, errors.New("a university with this name or location already exists")
	} else if err != mongo.ErrNoDocuments {
		return University{}, err
	}

	result, err := r.collection.InsertOne(c, university)
	if err != nil {
		return University{}, err
	}

	_, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return University{}, errors.New("failed to convert inserted ID to ObjectID")
	}

	return university, nil
}

func (r *repositoryImpl) GetAllUniversities(c context.Context) ([]University, error) {
	cursor, err := r.collection.Find(c, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var universities []University
	if err := cursor.All(c, &universities); err != nil {
		return nil, err
	}

	return universities, nil
}

func (r *repositoryImpl) GetUniversityByID(c context.Context, id string) (*University, error) {
	filter := bson.M{"_id": id}

	var university University
	err := r.collection.FindOne(c, filter).Decode(&university)
	if err != nil {
		return nil, err
	}

	return &university, nil
}
