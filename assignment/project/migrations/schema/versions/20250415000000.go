package versions

import (
	"golang-project/internal/model"

	"gorm.io/gorm"
)

func Migrate20250415000000(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
		&model.Tag{},
	)
}

func Rollback20250415000000(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&model.User{},
		&model.Post{},
		&model.Comment{},
		&model.Tag{},
	)
}
