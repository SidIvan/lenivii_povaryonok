package service

import (
	"PhotoParser/internal/dto"
	"math/rand"
)

var products = []string{"яблоко", "банан", "мясо", "яйцо", "макароны", "мука", "молоко", "огурец", "сыр", "масло"}

type PhotoParserService struct {
}

func NewPhotoParserService() *PhotoParserService {
	return &PhotoParserService{}
}

func (s *PhotoParserService) ParsePhoto(photo []byte) []dto.Product {
	photoProducts := make(map[string]struct{})
	numProducts := rand.Intn(11)
	for len(photoProducts) != numProducts {
		photoProducts[products[rand.Intn(10)]] = struct{}{}
	}
	result := make([]dto.Product, 0)
	for product := range photoProducts {
		result = append(result, dto.Product{Name: product})
	}
	return result
}
