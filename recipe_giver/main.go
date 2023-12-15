package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"recipe_giver/internal/config"
	"recipe_giver/internal/http_server"
	"strconv"
	"syscall"
	"time"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "cfg", ".config.yaml", "set config path")
	flag.Parse()
	cleanedFilePath := filepath.Clean(filePath)
	if err := config.ValidateConfigPath(cleanedFilePath); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	cfg, err := config.NewConfig(cleanedFilePath)
	if err != nil {
		fmt.Println(fmt.Errorf("fatal: init config %w", err))
		os.Exit(1)
	}
	server := http_server.NewServer(strconv.Itoa(cfg.Http.Port))
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	errWg, errCtx := errgroup.WithContext(ctx)
	errWg.Go(func() error {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})
	errWg.Go(func() error {
		<-errCtx.Done()
		ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		return server.Shutdown(ctxWithTimeOut)
	})
	errorInWait := errWg.Wait()
	if errors.Is(errorInWait, context.DeadlineExceeded) || errorInWait == nil {
		fmt.Println("Gracefully quit server")
	} else if errorInWait != nil {
		fmt.Println(errorInWait.Error())
	}
}
