package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/tkcrm/modules/pkg/db/postgres"
	"github.com/tkcrm/modules/pkg/logger"
)

// Config ...
type Config struct {
	AppName  string `default:"queues"`
	LogLevel string `default:"info"`
	Postgres postgres.Config
}

func (c *Config) Validate() error {
	// PostgreSQL
	if err := c.Postgres.Validate(); err != nil {
		return err
	}

	return validation.ValidateStruct(
		c,
		validation.Field(&c.AppName, validation.Required),
		validation.Field(&c.LogLevel, validation.Required, validation.In(logger.GetAllLevels()...)),
	)
}
