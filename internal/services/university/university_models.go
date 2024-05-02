package university

import "go.mongodb.org/mongo-driver/bson/primitive"

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type University struct {
	ID               primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Name             string               `json:"name" bson:"name" validate:"required"`
	Location         string               `json:"location" bson:"location" validate:"required"`
	NumberOfStudents int                  `json:"numberOfStudents" bson:"numberOfStudents" validate:"required"`
	Website          string               `json:"website" bson:"website" validate:"required"`
	Halls            []primitive.ObjectID `json:"halls" bson:"halls" validate:"required"`
	CoverImage       string               `json:"coverImage" bson:"coverImage" validate:"required"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ CREATE OBJECTS -----------------------------------------
// ---------------------------------------------------------------------------------------------------
type CreateUniversityRequest struct {
	Name             string   `json:"name" binding:"required"`
	Location         string   `json:"location" binding:"required"`
	NumberOfStudents int      `json:"numberOfStudents" binding:"required"`
	Website          string   `json:"website" binding:"required"`
	Halls            []string `json:"halls" binding:"required"`
	CoverImage       string   `json:"coverImage" binding:"required"`
}

// ---------------------------------------------------------------------------------------------------
// ----------------------------------------- RESPONSE OBJECTS ----------------------------------------
// ---------------------------------------------------------------------------------------------------
type UniversityResponse struct {
	ID               string   `json:"id,omitempty"`
	Name             string   `json:"name"`
	Location         string   `json:"location"`
	NumberOfStudents int      `json:"numberOfStudents"`
	Website          string   `json:"website"`
	Halls            []string `json:"halls"`
	CoverImage       string   `json:"coverImage"`
}
