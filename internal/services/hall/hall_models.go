package hall

import "go.mongodb.org/mongo-driver/bson/primitive"

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type Hall struct {
	ID         primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Name       string               `json:"name" bson:"name" validate:"required"`
	University primitive.ObjectID   `json:"university" bson:"university" validate:"required"`
	Location   string               `json:"location" bson:"location" validate:"required"`
	PostalCode string               `json:"postalCode" bson:"postalCode" validate:"required"`
	CoverImage string               `json:"coverImage" bson:"coverImage" validate:"required"`
	Reviews    []primitive.ObjectID `json:"reviews" bson:"reviews" validate:"required"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ CREATE OBJECTS -----------------------------------------
// ---------------------------------------------------------------------------------------------------
type CreateHallRequest struct {
	Name       string   `json:"name" binding:"required"`
	University string   `json:"universityId" binding:"required"`
	Location   string   `json:"location" binding:"required"`
	PostalCode string   `json:"postalCode" binding:"required"`
	CoverImage string   `json:"coverImage" binding:"required"`
	Reviews    []string `json:"reviews" binding:"required"`
}

// ---------------------------------------------------------------------------------------------------
// ----------------------------------------- RESPONSE OBJECTS ----------------------------------------
// ---------------------------------------------------------------------------------------------------
type HallResponse struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name"`
	University string   `json:"university"`
	Location   string   `json:"location"`
	PostalCode string   `json:"postalCode"`
	CoverImage string   `json:"coverImage"`
	Reviews    []string `json:"reviews"`
}
