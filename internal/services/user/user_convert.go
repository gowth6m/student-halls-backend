package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func ConvertCreateUserRequestToUser(req CreateUserRequest) (User, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	var universityID *primitive.ObjectID
	if req.University != nil && *req.University != "" {
		id, err := primitive.ObjectIDFromHex(*req.University)
		if err != nil {
			return User{}, err
		}
		universityID = &id
	}

	return User{
		ID:          primitive.NewObjectID(),
		Username:    req.Username,
		Password:    string(hashed),
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		University:  universityID,
		YearOfStudy: req.YearOfStudy,
		UserType:    "student",
	}, nil
}

func ConvertUserToUserResponse(user User) UserResponse {
	universityHex := ""
	if user.University != nil {
		universityHex = user.University.Hex()
	}

	return UserResponse{
		ID:          user.ID.Hex(),
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		University:  &universityHex,
		YearOfStudy: user.YearOfStudy,
		UserType:    user.UserType,
		UserImg:     user.UserImg,
	}
}
