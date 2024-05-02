package product

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ConvertCreateProductRequestToProduct converts CreateProductRequest to Product, including ID conversions.
func ConvertCreateProductRequestToProduct(req CreateProductRequest) (Product, error) {
	sellerID, err := primitive.ObjectIDFromHex(req.SellerId)
	if err != nil {
		return Product{}, err // Seller ID conversion failed
	}

	categoryID, err := primitive.ObjectIDFromHex(req.Category)
	if err != nil {
		return Product{}, err // Category ID conversion failed
	}

	var variants []Variant
	for _, v := range req.Variants {
		variant, err := ConvertCreateVariantRequestToVariant(v)
		if err != nil {
			return Product{}, err // Variant conversion failed
		}
		variants = append(variants, variant)
	}

	return Product{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		Seller:      sellerID,
		Variants:    variants,
		Category:    categoryID,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}, nil
}

// ConvertCreateVariantRequestToVariant converts CreateVariantRequest to Variant.
func ConvertCreateVariantRequestToVariant(req CreateVariantRequest) (Variant, error) {
	return Variant{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}, nil
}

// Convert Product to Productapi.
func ConvertProductToProductResponse(prod Product) ProductResponse {
	var variants []VariantResponse
	for _, v := range prod.Variants {
		variants = append(variants, VariantResponse{
			ID:          v.ID.Hex(),
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
			CreatedAt:   v.CreatedAt.Time().Format(time.RFC3339),
			UpdatedAt:   v.UpdatedAt.Time().Format(time.RFC3339),
		})
	}

	return ProductResponse{
		ID:          prod.ID.Hex(),
		Name:        prod.Name,
		Description: prod.Description,
		Seller:      prod.Seller.Hex(),
		Variants:    variants,
		Category:    prod.Category.Hex(),
		CreatedAt:   prod.CreatedAt.Time().Format(time.RFC3339),
		UpdatedAt:   prod.UpdatedAt.Time().Format(time.RFC3339),
	}
}
