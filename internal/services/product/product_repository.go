package product

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"student-halls.com/internal/db"
)

type ProductRepository interface {
	CreateProduct(c context.Context, product Product) (Product, error)
	CreateManyProduct(c context.Context, products []Product) ([]Product, error)
	GetAllProducts(c context.Context) ([]Product, error)
}

type repositoryImpl struct {
	collection *mongo.Collection
}

// NewProductRepository creates a new instance of a MongoDB-based product repository.
func NewProductRepository() ProductRepository {
	collection := db.Client.Database(db.DATABASE_NAME).Collection(db.COLLECTION_PRODUCTS)
	return &repositoryImpl{collection: collection}
}

// CreateProduct inserts a new product into the database.
func (r *repositoryImpl) CreateProduct(c context.Context, product Product) (Product, error) {
	result, err := r.collection.InsertOne(c, product)
	if err != nil {
		return Product{}, err
	}

	_, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return Product{}, errors.New("failed to convert inserted ID to ObjectID")
	}

	return product, nil
}

// CreateManyProduct inserts multiple products into the database.
func (r *repositoryImpl) CreateManyProduct(c context.Context, products []Product) ([]Product, error) {
	var interfaceSlice []interface{} = make([]interface{}, len(products))
	for i, v := range products {
		interfaceSlice[i] = v
	}

	result, err := r.collection.InsertMany(c, interfaceSlice)
	if err != nil {
		return nil, err
	}

	for i, id := range result.InsertedIDs {
		if oid, ok := id.(primitive.ObjectID); ok {
			products[i].ID = oid
		}
	}

	return products, nil
}

// GetAllProducts retrieves all products from the database.
func (r *repositoryImpl) GetAllProducts(c context.Context) ([]Product, error) {
	cursor, err := r.collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var products []Product
	if err := cursor.All(c, &products); err != nil {
		return nil, err
	}

	return products, nil
}
