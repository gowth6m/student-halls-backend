package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ CREATE OBJECTS -----------------------------------------
// ---------------------------------------------------------------------------------------------------
type CreateProductRequest struct {
	Name        string                 `json:"name" binding:"required"`
	Description string                 `json:"description"`
	SellerId    string                 `json:"sellerId" binding:"required"`
	Variants    []CreateVariantRequest `json:"variants" binding:"required"`
	Category    string                 `json:"category" binding:"required"`
}

type CreateVariantRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
}

// ---------------------------------------------------------------------------------------------------
// ----------------------------------------- RESPONSE OBJECTS ----------------------------------------
// ---------------------------------------------------------------------------------------------------
type ProductResponse struct {
	ID          string            `json:"id,omitempty"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Seller      string            `json:"seller"`
	Variants    []VariantResponse `json:"variants"`
	Category    string            `json:"category"`
	CreatedAt   string            `json:"createdAt,omitempty"`
	UpdatedAt   string            `json:"updatedAt,omitempty"`
}

type VariantResponse struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CreatedAt   string  `json:"createdAt,omitempty"`
	UpdatedAt   string  `json:"updatedAt,omitempty"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Seller      primitive.ObjectID `json:"seller" bson:"seller"`
	Variants    []Variant          `json:"variants,omitempty" bson:"variants,omitempty"`
	Category    primitive.ObjectID `json:"category,omitempty" bson:"category,omitempty"`
	CreatedAt   primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Variant struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64            `json:"price" bson:"price"`
	Stock       int                `json:"stock" bson:"stock"`
	CreatedAt   primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type Category struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}
