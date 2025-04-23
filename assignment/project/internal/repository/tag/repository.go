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

// Select finds and returns all tag models
func (r *repository) Select() ([]*model.Tag, error) {
	var result []*model.Tag

	query := r.db.Model(&model.Tag{}).Find(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// Read finds and returns the tag model by id
func (r *repository) Read(id int) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.First(&tag, id).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// Insert performs insert action into tag table
func (r *repository) Insert(tag *model.Tag) error {
	return r.db.Create(tag).Error
}
