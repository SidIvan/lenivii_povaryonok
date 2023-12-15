package http_server

import (
	"net/http"
	"recipe_giver/internal/handlers"
)

func NewServer(port string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/get_ingredients_list", handlers.GetIngredients)
	mux.HandleFunc("/status", handlers.GetStatus)
	mux.HandleFunc("/get_recipe", handlers.GetRecipe)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	return server
}
