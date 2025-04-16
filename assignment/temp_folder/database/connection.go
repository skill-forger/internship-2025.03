package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connection represents the database connection
type Connection interface {
	DataSourceName() string
	Open() (*gorm.DB, error)
	Close() error
	Instance() (*gorm.DB, error)
	Ping() error
}

// connection is an implementation of the database Connection
type connection struct {
	dsn      string
	config   *Config
	instance *gorm.DB
}

// NewConnection creates and returns a database connection instance
func NewConnection(dsn string, config *Config) Connection {
	if config == nil {
		config = newDefaultConfig()
	}

	if config.Config == nil {
		config.Config = newGormConfig()
	}

	return &connection{dsn: dsn, config: config}
}

// DataSourceName returns the data source connection string
func (c *connection) DataSourceName() string {
	return c.dsn
}

// Open initializes a new database client
func (c *connection) Open() (*gorm.DB, error) {
	if c.config == nil || c.config.Config == nil {
		return nil, ErrMissingConfig
	}

	var err error
	c.instance, err = gorm.Open(mysql.Open(c.dsn), c.config.Config)
	if nil != err {
		return nil, err
	}

	instanceDb, err := c.instance.DB()
	if nil != err {
		return nil, err
	}

	if c.config.MaxOpenConnections > 0 {
		instanceDb.SetMaxOpenConns(c.config.MaxOpenConnections)
	}

	if c.config.MaxIdleConnections > 0 {
		instanceDb.SetMaxIdleConns(c.config.MaxIdleConnections)
	}

	if c.config.ConnectionMaxTime > 0 {
		instanceDb.SetConnMaxLifetime(c.config.ConnectionMaxTime)
	}

	if c.config.ConnectionIdleTime > 0 {
		instanceDb.SetConnMaxIdleTime(c.config.ConnectionIdleTime)
	}

	return c.instance, nil
}

// Close closes the current database client
func (c *connection) Close() error {
	if c.instance == nil {
		return ErrUninitializedDatabase
	}

	gormDb, err := c.instance.DB()
	if err != nil {
		return err
	}

	return gormDb.Close()
}

// Instance return the current instance of the Mongo database client
func (c *connection) Instance() (*gorm.DB, error) {
	if c.instance == nil {
		return nil, ErrUninitializedDatabase
	}

	return c.instance, nil
}

// Ping verifies if the current database client is active and healthy
func (c *connection) Ping() error {
	instance, err := c.Instance()
	if err != nil {
		return err
	}

	gormDb, err := instance.DB()
	if err != nil {
		return err
	}

	return gormDb.Ping()
}
