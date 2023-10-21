package app

import (
	"PhotoParser/internal/config"
	"PhotoParser/internal/logging"
	"PhotoParser/internal/service"
	"PhotoParser/internal/web"
	"context"
	"fmt"
	"log"
	"net/http"
)

type PhotoParserApp struct {
	cfg    *config.PhotoParserConfig
	server *http.Server
}

func NewPhotoParserApp(cfg *config.PhotoParserConfig) *PhotoParserApp {
	return &PhotoParserApp{
		cfg: cfg,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port),
			Handler: initApi(cfg),
		},
	}
}

func initApi(cfg *config.PhotoParserConfig) http.Handler {
	postMux := http.NewServeMux()
	photoParserController := web.NewPhotoParserController(service.NewPhotoParserService())
	postMux.HandleFunc("/parsePhoto", photoParserController.ParsePhotoHandler)
	return logging.Logging(postMux)
}

func (a *PhotoParserApp) Start() {
	go func() {
		err := a.server.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()
}

func (a *PhotoParserApp) Stop(ctx context.Context) {
	a.server.Shutdown(ctx)
}
