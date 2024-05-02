package shop

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"student-halls.com/internal/db"
)

type ShopRepository interface {
	CreateShop(c context.Context, shop Shop) (Shop, error)
	GetAllShops(c context.Context) ([]Shop, error)
}

type repositoryImpl struct {
	collection *mongo.Collection
}

func NewShopRepository() ShopRepository {
	collection := db.Client.Database(db.DATABASE_NAME).Collection(db.COLLECTION_SHOPS)
	return &repositoryImpl{collection: collection}
}

func (r *repositoryImpl) CreateShop(c context.Context, shop Shop) (Shop, error) {
	result, err := r.collection.InsertOne(c, shop)
	if err != nil {
		return Shop{}, err
	}

	_, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return Shop{}, errors.New("failed to convert inserted ID to ObjectID")
	}

	return shop, nil
}

func (r *repositoryImpl) GetAllShops(c context.Context) ([]Shop, error) {
	cursor, err := r.collection.Find(c, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var shops []Shop
	if err := cursor.All(c, &shops); err != nil {
		return nil, err
	}

	return shops, nil
}
