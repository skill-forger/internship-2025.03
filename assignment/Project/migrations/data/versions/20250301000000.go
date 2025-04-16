package versions

import (
	"gorm.io/gorm"

	"golang-project/internal/model"
	"golang-project/util/hashing"
)

// Migrate20250301000000 performs migration logic for version 20250301000000
func Migrate20250301000000(db *gorm.DB) error {
	demoPassword, err := hashing.NewBcrypt().Generate([]byte("demouser@123"))
	if err != nil {
		return err
	}

	type User struct {
		model.BaseModel
		FirstName    string
		LastName     string
		Email        string
		Password     string
		Pseudonym    string
		ProfileImage string
		Biography    string
	}

	data := &User{
		FirstName:    "demo",
		LastName:     "user",
		Email:        "user@demo.com",
		Password:     string(demoPassword),
		Pseudonym:    "demo_user",
		ProfileImage: "demo user profile image",
		Biography:    "this is the demo user biography",
	}

	return db.Model(&User{}).Create(data).Error
}

// Rollback20250301000000 performs rollback logic for version 20250301000000
func Rollback20250301000000(db *gorm.DB) error {
	type User struct {
		model.BaseModel
	}

	return db.Unscoped().Delete(&User{}, map[string]any{"email": "user@demo.com"}).Error
}
