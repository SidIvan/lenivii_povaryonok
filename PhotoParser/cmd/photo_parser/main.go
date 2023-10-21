package main

import (
	"PhotoParser/internal/app"
	"PhotoParser/internal/config"
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "cmd/photo_parser/mock.yaml", "mock config set")
	flag.Parse()
	cfg := config.NewPhotoParserConfig(cfgPath)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	app := app.NewPhotoParserApp(cfg)
	app.Start()
	log.Printf("Server started on port %d\n", cfg.Port)
	<-ctx.Done()
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(cfg.GracefulTimeoutSec)*time.Second)
	defer cancel()
	app.Stop(ctx)
}
