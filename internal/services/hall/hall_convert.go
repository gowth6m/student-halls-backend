package hall

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertCreateHallRequestToHall(req CreateHallRequest) (Hall, error) {
	universityID, err := primitive.ObjectIDFromHex(req.University)
	if err != nil {
		return Hall{}, err
	}

	return Hall{
		ID:         primitive.NewObjectID(),
		Name:       req.Name,
		University: universityID,
		Location:   req.Location,
		PostalCode: req.PostalCode,
		CoverImage: req.CoverImage,
		Reviews:    []primitive.ObjectID{},
	}, nil
}

func ConvertHallToHallResponse(hall Hall) HallResponse {
	reviews := []string{}
	for _, review := range hall.Reviews {
		reviews = append(reviews, review.Hex())
	}

	return HallResponse{
		ID:         hall.ID.Hex(),
		Name:       hall.Name,
		University: hall.University.Hex(),
		Location:   hall.Location,
		PostalCode: hall.PostalCode,
		CoverImage: hall.CoverImage,
		Reviews:    reviews,
	}
}
