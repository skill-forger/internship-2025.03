package versions

import (
	"gorm.io/gorm"

	"golang-project/internal/model"
)

// Migrate20250301000000 performs migration logic for version 20250301000000
func Migrate20250301000000(db *gorm.DB) error {
	type User struct {
		model.BaseModel
		FirstName    string `gorm:"TYPE:VARCHAR(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci;NOT NULL"`
		LastName     string `gorm:"TYPE:VARCHAR(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci;NOT NULL"`
		Email        string `gorm:"TYPE:VARCHAR(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci;NOT NULL;uniqueIndex"`
		Password     string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci;NOT NULL"`
		Pseudonym    string `gorm:"TYPE:VARCHAR(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci;NOT NULL;uniqueIndex"`
		ProfileImage string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci"`
		Biography    string `gorm:"TYPE:TEXT CHARACTER SET utf8 COLLATE utf8_unicode_ci"`
	}

	return db.AutoMigrate(&User{})
}

// Rollback20250301000000 performs rollback logic for version 20250301000000
func Rollback20250301000000(db *gorm.DB) error {
	type User struct {
		model.BaseModel
	}

	return db.Migrator().DropTable(&User{})
}
