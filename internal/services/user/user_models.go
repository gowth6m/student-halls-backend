package user

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type Session struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ CREATE OBJECTS -----------------------------------------
// ---------------------------------------------------------------------------------------------------
type CreateUserRequest struct {
	Username    string  `json:"username" binding:"required"`
	Password    string  `json:"password" binding:"required"`
	Email       string  `json:"email" binding:"required"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	University  *string `json:"university"`
	YearOfStudy *int    `json:"yearOfStudy"`
}

// ---------------------------------------------------------------------------------------------------
// ----------------------------------------- RESPONSE OBJECTS ----------------------------------------
// ---------------------------------------------------------------------------------------------------
type UserResponse struct {
	ID          string  `json:"id,omitempty"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	FirstName   string  `json:"firstName,omitempty"`
	LastName    string  `json:"lastName,omitempty"`
	University  *string `json:"university,omitempty"`
	YearOfStudy *int    `json:"yearOfStudy,omitempty"`
	UserType    string  `json:"userType,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  UserResponse
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type User struct {
	ID          primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Username    string              `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Password    string              `json:"-" bson:"password,omitempty" validate:"required"`
	Email       string              `json:"email" bson:"email" validate:"required,email"`
	FirstName   string              `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName    string              `json:"lastName,omitempty" bson:"lastName,omitempty"`
	University  *primitive.ObjectID `json:"university,omitempty" bson:"university,omitempty"`
	YearOfStudy *int                `json:"yearOfStudy,omitempty" bson:"yearOfStudy,omitempty"`
	UserType    string              `json:"userType,omitempty" bson:"userType,omitempty"`
	CreatedAt   primitive.DateTime  `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   primitive.DateTime  `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
