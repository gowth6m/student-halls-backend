package university

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertCreateUniversityRequestToUniversity(req CreateUniversityRequest) (University, error) {

	return University{
		ID:               primitive.NewObjectID(),
		Name:             req.Name,
		Location:         req.Location,
		Website:          req.Website,
		Halls:            []primitive.ObjectID{},
		CoverImage:       req.CoverImage,
		NumberOfStudents: req.NumberOfStudents,
	}, nil
}

func ConvertUniversityToUniversityResponse(university University) UniversityResponse {
	halls := []string{}
	for _, hall := range university.Halls {
		halls = append(halls, hall.Hex())
	}

	return UniversityResponse{
		ID:               university.ID.Hex(),
		Name:             university.Name,
		Location:         university.Location,
		Website:          university.Website,
		Halls:            halls,
		CoverImage:       university.CoverImage,
		NumberOfStudents: university.NumberOfStudents,
	}
}
