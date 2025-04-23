package versions

import (
	"gorm.io/gorm"
)

func Migrate20250423000000(db *gorm.DB) error {
	// Rename column IsPublic to IsPublished in the posts table
	return db.Migrator().RenameColumn("posts", "is_public", "is_published")
}

func Rollback20250423000000(db *gorm.DB) error {
	// Rename column IsPublished back to IsPublic in the posts table
	return db.Migrator().RenameColumn("posts", "is_published", "is_public")
}
