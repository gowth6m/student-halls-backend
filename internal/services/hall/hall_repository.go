package hall

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"student-halls.com/internal/db"
)

type HallRepository interface {
	CreateHall(c context.Context, hall Hall) (Hall, error)
	GetAllHalls(c context.Context) ([]Hall, error)
	GetHallByID(c context.Context, id string) (*Hall, error)
	GetHallsByUniversityID(c context.Context, universityID string) ([]Hall, error)
}

type repositoryImpl struct {
	collection *mongo.Collection
}

func NewHallRepository() HallRepository {
	collection := db.Client.Database(db.DATABASE_NAME).Collection(db.COLLECTION_HALL)
	return &repositoryImpl{collection: collection}
}

func (r *repositoryImpl) CreateHall(c context.Context, hall Hall) (Hall, error) {
	result, err := r.collection.InsertOne(c, hall)
	if err != nil {
		return Hall{}, err
	}

	_, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return Hall{}, errors.New("failed to convert inserted ID to ObjectID")
	}

	return hall, nil
}

func (r *repositoryImpl) GetAllHalls(c context.Context) ([]Hall, error) {
	cursor, err := r.collection.Find(c, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var halls []Hall
	if err := cursor.All(c, &halls); err != nil {
		return nil, err
	}

	return halls, nil
}

func (r *repositoryImpl) GetHallByID(c context.Context, id string) (*Hall, error) {
	hallID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": hallID}
	var hall Hall
	err = r.collection.FindOne(c, filter).
		Decode(&hall)
	if err != nil {
		return nil, err
	}

	return &hall, nil
}

func (r *repositoryImpl) GetHallsByUniversityID(c context.Context, universityID string) ([]Hall, error) {
	uID, err := primitive.ObjectIDFromHex(universityID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"university": uID}
	cursor, err := r.collection.Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var halls []Hall
	if err := cursor.All(c, &halls); err != nil {
		return nil, err
	}

	return halls, nil
}
