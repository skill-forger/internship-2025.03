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

// Select finds and returns all not deleted tag models
func (r *repository) Select() ([]*model.Tag, error) {
	var result []*model.Tag

	query := r.db.Model(&model.Tag{}).Find(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// SelectPost finds and returns all not deleted published post models by tag id
func (r *repository) SelectPost(id int) ([]*model.Post, error) {
	var result []*model.Post

	query := r.db.Model(&model.Post{}).
		Joins("JOIN post_tag ON posts.id =post_tag.post_id").
		Where("post_tag.tag_id = ?", id).
		Where("posts.deleted_at IS NULL").
		Where("posts.is_published = ?", true).
		Find(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// SelectByPost finds and returns all not deleted tag models by post id
func (r *repository) SelectByPost(id int) ([]*model.Tag, error) {
	var result []*model.Tag

	query := r.db.Model(&model.Tag{}).
		Joins("JOIN post_tag ON tags.id = post_tag.tag_id").
		Where("post_tag.post_id = ?", id).
		Where("tags.deleted_at IS NULL").
		Find(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// SelectUserByPost finds and returns the not deleted user model by post id
func (r *repository) SelectUserByPost(id int) (*model.User, error) {
	var result *model.User

	query := r.db.Model(&model.User{}).
		Joins("JOIN posts ON users.id = posts.user_id").
		Where("users.deleted_at IS NULL").
		Where("posts.id = ?", id).
		First(&result)

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

// Delete tags that are not used in any post
func (r *repository) Delete(id int) error {
	return r.db.Delete(&model.Tag{}, id).Error
}

// HasPosts check if the tag is used in any post
func (r *repository) HasPosts(id int) (bool, error) {
	var count int64
	err := r.db.Table("post_tag").Where("tag_id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil

}
