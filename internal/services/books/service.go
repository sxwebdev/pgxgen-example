package books

import (
	"github.com/sxwebdev/pgxgen-example/internal/config"
	"github.com/sxwebdev/pgxgen-example/internal/store"
	"github.com/tkcrm/modules/logger"
)

type Service struct {
	logger logger.Logger
	config *config.Config
	store  store.IStore
}

func New(l logger.Logger, c *config.Config, st store.IStore) *Service {
	return &Service{
		logger: l,
		config: c,
		store:  st,
	}
}
