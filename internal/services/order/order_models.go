package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ProductID  primitive.ObjectID `json:"productId" bson:"productId"`
	VariantID  primitive.ObjectID `json:"variantId" bson:"variantId"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	UnitPrice  float64            `json:"unitPrice" bson:"unitPrice"`
	TotalPrice float64            `json:"totalPrice" bson:"totalPrice"`
}

type Order struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID     primitive.ObjectID `json:"userId" bson:"userId"`
	Items      []OrderItem        `json:"items" bson:"items"`
	TotalPrice float64            `json:"totalPrice" bson:"totalPrice"`
	Status     string             `json:"status" bson:"status"`
	CreatedAt  primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
