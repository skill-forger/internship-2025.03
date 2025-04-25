package comment

import (
	repo "golang-project/internal/repository"
	"gorm.io/gorm"
)

// repository represents the implementation of repository.Comment
type repository struct {
	db *gorm.DB
}

// NewRepository returns a new implementation of repository.Comment
func NewRepository(db *gorm.DB) repo.Comment {
	return &repository{db: db}
}
