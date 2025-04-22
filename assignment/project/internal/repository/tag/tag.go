package tag

import (
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	"gorm.io/gorm"
)

// repository represents the implementation of repository.Tag
type repository struct {
	db *gorm.DB
}

// NewRepository returns a new implementation of repository.Tag
func NewRepository(db *gorm.DB) repo.Tag {
	return &repository{db: db}
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
	if err := r.db.Create(tag).Error; err != nil {
		return err
	}
	return nil
}

// Delete performs delete action into tag table
func (r *repository) Delete(id int) error {
	if err := r.db.Delete(&model.Tag{}, id).Error; err != nil {
		return err
	}
	return nil
}

// HasPost checks if a tag is associated with any post
func (r *repository) HasPost(id int) (bool, error) {
	var count int64
	err := r.db.Table("post_tag").Where("tag_id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
