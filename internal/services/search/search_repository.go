package search

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"student-halls.com/internal/db"
	"student-halls.com/internal/services/hall"
	"student-halls.com/internal/services/university"
)

type SearchCriteria struct {
	Query string
}

// SearchResult defines the structure for holding search results.
type SearchResult struct {
	Universities []university.University
	Halls        []hall.Hall
}

type SearchRepository interface {
	Search(c context.Context, criteria SearchCriteria) (SearchResult, error)
}

type searchRepositoryImpl struct {
	universityCollection *mongo.Collection
	hallCollection       *mongo.Collection
}

func NewSearchRepository() SearchRepository {
	return &searchRepositoryImpl{
		universityCollection: (*mongo.Collection)(db.Client.Database(db.DATABASE_NAME).Collection(db.COLLECTION_UNIVERSITY)),
		hallCollection:       (*mongo.Collection)(db.Client.Database(db.DATABASE_NAME).Collection(db.COLLECTION_HALL)),
	}
}

func (s *searchRepositoryImpl) Search(c context.Context, criteria SearchCriteria) (SearchResult, error) {
	var result SearchResult

	// Search in Universities
	universityFilter := bson.M{"name": bson.M{"$regex": criteria.Query, "$options": "i"}}
	universityCursor, err := s.universityCollection.Find(c, universityFilter)
	if err != nil {
		return SearchResult{}, err
	}
	defer universityCursor.Close(c)
	for universityCursor.Next(c) {
		var university university.University
		if err := universityCursor.Decode(&university); err != nil {
			continue
		}
		result.Universities = append(result.Universities, university)
	}

	// Search in Halls
	hallFilter := bson.M{"name": bson.M{"$regex": criteria.Query, "$options": "i"}}
	hallCursor, err := s.hallCollection.Find(c, hallFilter)
	if err != nil {
		return SearchResult{}, err
	}
	defer hallCursor.Close(c)
	for hallCursor.Next(c) {
		var hall hall.Hall
		if err := hallCursor.Decode(&hall); err != nil {
			continue
		}
		result.Halls = append(result.Halls, hall)
	}

	return result, nil
}
