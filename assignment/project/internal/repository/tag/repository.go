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

// Select finds and returns all not deleted tag models by tag IDs
// If the id parameter is nil, it will return all tags

func (r *repository) Select(id []int) ([]*model.Tag, error) {
	var result []*model.Tag

	query := r.db.Model(&model.Tag{})

	if id != nil {
		query.Where("id IN (?)", id)
	}

	query.Find(&result)

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

// SelectPost finds and returns all post_tag models by post ids
func (r *repository) SelectPostTag(id []int) ([]*model.PostTag, error) {
	var result []*model.PostTag

	query := r.db.Table("post_tag").
		Where("post_id IN (?)", id).
		Find(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// SelectUser finds and returns the not deleted user models user ids
func (r *repository) SelectUser(ids []int) ([]*model.User, error) {
	var result []*model.User

	query := r.db.Model(&model.User{}).
		Where("id IN (?)", ids).
		Find(&result)

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
