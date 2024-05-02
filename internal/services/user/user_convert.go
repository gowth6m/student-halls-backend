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

	var universityID primitive.ObjectID
	if req.University != "" {
		universityID, err = primitive.ObjectIDFromHex(req.University)
		if err != nil {
			return User{}, err
		}
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
	}, nil
}

func ConvertUserToUserResponse(user User) UserResponse {
	return UserResponse{
		ID:          user.ID.Hex(),
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		University:  user.University.Hex(),
		YearOfStudy: user.YearOfStudy,
	}
}
