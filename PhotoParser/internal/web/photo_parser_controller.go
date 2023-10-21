package web

import (
	"PhotoParser/internal/service"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type PhotoParserController struct {
	service *service.PhotoParserService
}

func NewPhotoParserController(service *service.PhotoParserService) *PhotoParserController {
	return &PhotoParserController{
		service: service,
	}
}

func (c *PhotoParserController) ParsePhotoHandler(w http.ResponseWriter, r *http.Request) {
	photo, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("read %d bytes\n", len(photo))
	products := c.service.ParsePhoto(photo)
	responseData, err := json.Marshal(products)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	numBytes, err := w.Write(responseData)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("sent %d bytes\n", numBytes)
}
