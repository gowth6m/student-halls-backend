package shop

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"student-halls.com/internal/services/product"
	"time"
)

func ConvertCreateShopToShop(req CreateShopRequest) (Shop, error) {
	ownerID, err := primitive.ObjectIDFromHex(req.OwnerID)
	if err != nil {
		return Shop{}, err
	}

	return Shop{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     ownerID,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}, nil
}

func ConvertShopToShopResponse(shop Shop) ShopResponse {
	return ShopResponse{
		ID:        shop.ID.Hex(),
		Name:      shop.Name,
		OwnerID:   shop.OwnerID.Hex(),
		Products:  []product.ProductResponse{},
		Address:   shop.Address,
		CreatedAt: shop.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt: shop.UpdatedAt.Time().Format(time.RFC3339),
	}
}
