package schema

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"golang-project/migrations/schema/versions"
)

// NewMigration returns new gorm schema migration instance
func NewMigration(db *gorm.DB) *gormigrate.Gormigrate {
	option := gormigrate.DefaultOptions
	option.TableName = "schema_migrations"

	return gormigrate.New(db, option, []*gormigrate.Migration{
		{
			ID:       "20250301000000",
			Migrate:  versions.Migrate20250301000000,
			Rollback: versions.Rollback20250301000000,
		},

		{
			ID:       "20250416000000",
			Migrate:  versions.Migrate20250416000000,
			Rollback: versions.Rollback20250416000000,
		},

		{
			ID:       "20250423000000",
			Migrate:  versions.Migrate20250423000000,
			Rollback: versions.Rollback20250423000000,
		},
	})
}
