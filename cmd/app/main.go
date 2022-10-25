package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/sxwebdev/pgxgen-example/internal/config"
	"github.com/sxwebdev/pgxgen-example/internal/server"
	"github.com/tkcrm/modules/cfg"
	"github.com/tkcrm/modules/logger"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)
	defer stop()

	// Load configuration
	var config config.Config
	if err := cfg.LoadConfig(&config); err != nil {
		logger.New().Fatalf("could not load configuration: %v", err)
	}

	// Init logger
	logger := logger.New(
		logger.WithLogLevel(logger.LogLevel(config.LogLevel)),
		logger.WithAppName(config.AppName),
	)
	defer logger.Sync()

	// Init Server
	srv, err := server.New(ctx, logger, &config)
	if err != nil {
		logger.Fatalf("init server error: %v", err)
	}

	// Start server
	if err := srv.Start(ctx); err != nil {
		logger.Fatalf("start server error: %v", err)
	}
}
