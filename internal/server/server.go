package server

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sxwebdev/pgxgen-example/internal/config"
	"github.com/sxwebdev/pgxgen-example/internal/handler"
	"github.com/sxwebdev/pgxgen-example/internal/store"
	"github.com/tkcrm/modules/db/postgres"
	"github.com/tkcrm/modules/logger"
)

type Server struct {
	logger  logger.Logger
	config  *config.Config
	db      *postgres.PostgreSQL
	store   store.IStore
	handler *handler.Handler
}

func New(ctx context.Context, logger logger.Logger, config *config.Config) (*Server, error) {
	s := &Server{
		logger: logger,
		config: config,
	}

	// Connect to database
	conn, err := postgres.New(ctx, config.Postgres, logger)
	if err != nil {
		return nil, errors.Wrap(err, "error initializing db")
	}
	s.db = conn
	s.store = store.New(s.db.DB)

	// init handker
	s.handler = handler.New(logger, config, s.store)

	return s, nil
}

func (s *Server) Start(ctx context.Context) error {
	defer s.stop()

	s.logger.Info("server successfully started")

	<-ctx.Done()

	return nil
}

func (s *Server) stop() {
	s.db.Close()
	s.logger.Info("server stopped")
}
