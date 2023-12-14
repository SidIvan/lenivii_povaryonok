package service

import (
	"PhotoParser/internal/dto"
	"bytes"
	"log"
	"math/rand"
	"os/exec"
	"strings"
)

var mockProducts = []string{"яблоко", "банан", "мясо", "яйцо", "макароны", "мука", "молоко", "огурец", "сыр", "масло"}
var products = map[string]string{
	"0": "яблоко",
	"1": "банан",
	"2": "виноград",
	"3": "апельсин",
	"4": "ананас",
	"5": "арбуз",
}

type PhotoParserService struct {
}

func NewPhotoParserService() *PhotoParserService {
	return &PhotoParserService{}
}

func mockResult() []dto.Product {
	photoProducts := make(map[string]struct{})
	numProducts := rand.Intn(11)
	for len(photoProducts) != numProducts {
		photoProducts[mockProducts[rand.Intn(10)]] = struct{}{}
	}
	result := make([]dto.Product, 0)
	for product := range photoProducts {
		result = append(result, dto.Product{Name: product})
	}
	return result
}

func (s *PhotoParserService) ParsePhoto(photo []byte) []dto.Product {
	//"../yolov7/detect.py", "--weights", "weights.pt", "--img-size", "640", "--save-txt", "--nosave", "--source", "dist/fruits.jpg"
	cmd := exec.Command("python3", "../yolov7/detect.py", "--weights", "cmd/photo_parser/weights.pt", "--img-size", "640", "--nosave")
	cmd.Stdin = bytes.NewReader(photo)
	res := bytes.NewBufferString("")
	cmd.Stdout = res
	errs := bytes.NewBufferString("")
	cmd.Stderr = errs
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Model's stdout:")
	log.Println(res.String())
	log.Println("Model's stderr:")
	log.Println(errs.String())
	result := make([]dto.Product, 0)
	for _, productNo := range strings.Split(res.String(), " ") {
		result = append(result, dto.Product{Name: products[productNo]})
	}
	return result
}
