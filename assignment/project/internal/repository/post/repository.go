package post

import (
	"golang-project/internal/model"
	repo "golang-project/internal/repository"

	"gorm.io/gorm"
)

// repository represents the implementation of repository.Post
type repository struct {
	db *gorm.DB
}

// NewRepository returns a new implementation if repository.Post
func NewRepository(db *gorm.DB) repo.Post {
	return &repository{db: db}
}

func (r *repository) Insert(post *model.Post) (*model.Post, error) {
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

// GetTagsByPostID retrieves all tags associated with a post
func (r *repository) GetTagsByPostID(postID int) ([]*model.Tag, error) {
	var tags []*model.Tag

	err := r.db.Table("tags").
		Joins("JOIN post_tag ON tags.id = post_tag.tag_id").
		Where("post_tag.post_id = ?", postID).
		Find(&tags).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}

// AddPostTag associates a tag with a post
func (r *repository) AddPostTag(postID int, tagID int) error {
	postTag := &model.PostTag{
		PostID: postID,
		TagID:  tagID,
	}

	return r.db.Create(postTag).Error
}

// InsertManyPostTags batch insert post_tag
func (r *repository) InsertManyPostTags(postTags []*model.PostTag) error {
	return r.db.Create(&postTags).Error
}

// FindSlugsLike retrieves all slugs that start with the given baseSlug
func (r *repository) FindSlugsLike(baseSlug string) ([]string, error) {
	var slugs []string
	pattern := baseSlug + "%"
	err := r.db.Model(&model.Post{}).Where("slug LIKE ?", pattern).Pluck("slug", &slugs).Error
	if err != nil {
		return nil, err
	}
	return slugs, nil
}
