package data

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"golang-project/migrations/data/versions"
)

// NewMigration returns new gorm data migration instance
func NewMigration(db *gorm.DB) *gormigrate.Gormigrate {
	option := gormigrate.DefaultOptions
	option.TableName = "data_migrations"

	return gormigrate.New(db, option, []*gormigrate.Migration{
		{
			ID:       "20250301000000",
			Migrate:  versions.Migrate20250301000000,
			Rollback: versions.Rollback20250301000000,
		},
	})
}
