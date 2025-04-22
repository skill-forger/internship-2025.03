package tag

import (
	"gorm.io/gorm"

	"golang-project/internal/model"
	repo "golang-project/internal/repository"
)

// repository represents the implementation of repository.Tag
type repository struct {
	db *gorm.DB
}

// NewRepository returns a new implementation of repository.Tag
func NewRepository(db *gorm.DB) repo.Tag {
	return &repository{db: db}
}

// ReadAllTags finds and returns all tag models
func (r *repository) ReadAll() ([]*model.Tag, error) {
	var result []*model.Tag

	query := r.db.Model(&model.Tag{}).Find(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}
