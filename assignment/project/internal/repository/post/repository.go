package post

import (
	"gorm.io/gorm"

	"golang-project/internal/model"
	repo "golang-project/internal/repository"
)

// repository represents the implementation of repository.Post
type repository struct {
	db *gorm.DB
}

// NewRepository returns a new implementation of repository.Post
func NewRepository(db *gorm.DB) repo.Post {
	return &repository{db: db}
}

// Read finds and returns the post model by id
func (r *repository) Read(id int) (*model.Post, error) {
	var result model.Post

	query := r.db.Model(&model.Post{}).
		Preload("User").
		Where("id = ? AND is_published = ?", id, true).
		First(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	// Get tags for the post
	tags, err := r.GetTags(id)
	if err != nil {
		return nil, err
	}
	result.Tags = tags

	return &result, nil
}

// GetTags retrieves all tags associated with a post
func (r *repository) GetTags(postID int) ([]*model.Tag, error) {
	var tags []*model.Tag

	err := r.db.Table("tags").
		Joins("JOIN post_tag ON tags.id = post_tag.tag_id").
		Where("post_tag.post_id = ? AND tags.deleted_at IS NULL", postID).
		Find(&tags).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}
