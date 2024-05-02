package shop

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"student-halls.com/internal/services/product"
)

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ CREATE OBJECTS -----------------------------------------
// ---------------------------------------------------------------------------------------------------
type CreateShopRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	OwnerID     string `json:"ownerId" binding:"required"`
	Address     string `json:"address"`
}

// ---------------------------------------------------------------------------------------------------
// ----------------------------------------- RESPONSE OBJECTS ----------------------------------------
// ---------------------------------------------------------------------------------------------------
type ShopResponse struct {
	ID        string                    `json:"id,omitempty"`
	Name      string                    `json:"name"`
	OwnerID   string                    `json:"ownerId"`
	Address   string                    `json:"address"`
	Products  []product.ProductResponse `json:"products"`
	CreatedAt string                    `json:"createdAt,omitempty"`
	UpdatedAt string                    `json:"updatedAt,omitempty"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type Shop struct {
	ID          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description,omitempty" bson:"description,omitempty"`
	OwnerID     primitive.ObjectID   `json:"ownerId" bson:"ownerId"`
	Products    []primitive.ObjectID `json:"products,omitempty" bson:"products,omitempty"`
	Address     string               `json:"address,omitempty" bson:"address,omitempty"`
	CreatedAt   primitive.DateTime   `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   primitive.DateTime   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
