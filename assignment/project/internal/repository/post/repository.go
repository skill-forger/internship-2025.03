package post

import (
	"fmt"

	"gorm.io/gorm"

	"golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	"golang-project/util/pagination"
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
		Preload("Tags", "deleted_at IS NULL").
		Where("id = ? AND is_published = ?", id, true).
		First(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	return &result, nil
}

// Select retrieves all posts from the database with optional filters
func (r *repository) Select(filter *contract.ListPostRequest) ([]*model.Post, error) {
	var posts []*model.Post

	query := r.db.Model(&model.Post{}).
		Preload("User").
		Preload("Tags", "deleted_at IS NULL").
		Where("is_published = ?", true)

	// Apply filters if provided
	if filter != nil {
		if filter.Title != "" {
			title := fmt.Sprintf("%%%s%%", filter.Title)
			query = query.Where("title LIKE ?", title)
		}
		if filter.Pseudonym != "" {
			pseudonym := fmt.Sprintf("%%%s%%", filter.Pseudonym)
			query = query.Joins("JOIN users ON posts.user_id = users.id").
				Where("users.pseudonym LIKE ?", pseudonym)
		}
		if filter.Tag != "" {
			query = query.Joins("JOIN post_tag ON posts.id = post_tag.post_id").
				Joins("JOIN tags ON post_tag.tag_id = tags.id").
				Where("tags.name = ? AND tags.deleted_at IS NULL", filter.Tag)
		}
	}

	// Apply pagination
	if filter != nil && filter.Page > 0 && filter.PageSize > 0 {
		offset := pagination.CalculateOffset(filter.Page, filter.PageSize)
		query = query.Offset(offset).Limit(filter.PageSize)
	}

	err := query.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}
