package database

import (
	"errors"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ErrMissingConfig         = errors.New("database config is missing")
	ErrUninitializedDatabase = errors.New("database instance is not initialized")
)

// Config represents the database connection basic configuration
type Config struct {
	*gorm.Config
	ConnectionMaxTime  time.Duration
	ConnectionIdleTime time.Duration
	MaxIdleConnections int
	MaxOpenConnections int
}

// newDefaultConfig returns a default configuration for database Connection
func newDefaultConfig() *Config {
	return &Config{
		Config:             newGormConfig(),
		MaxIdleConnections: 2,
		MaxOpenConnections: 4,
	}
}

// newGormConfig returns the default Gorm configuration for database Connection
func newGormConfig() *gorm.Config {
	return &gorm.Config{Logger: logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Duration(200) * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
			ParameterizedQueries:      true,
		},
	)}
}
