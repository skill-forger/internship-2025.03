package tag

import (
	"github.com/labstack/echo/v4"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	"gorm.io/gorm"
	"net/http"
)

// repository represents the implementation of repository.Tag
type repository struct {
	db *gorm.DB
}

func (r *repository) GetPostByTagId(id int) ([]*model.Post, error) {
	var results []*model.Post

	query := r.db.Model(&model.Post{}).
		Joins("JOIN post_tags ON posts.id = post_tags.post_id").
		Where("post_tags.tag_id = ?", id).
		Find(&results)

	if err := query.Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (r *repository) Delete(id int) error {
	// Check if tag exists
	tag, err := r.Read(id)

	if err != nil {
		return err
	}

	// Check if tag has any posts
	var count int64
	if err := r.db.Model(&model.Post{}).Joins("JOIN post_tags ON posts.id = post_tags.post_id").Where("post_tags.tag_id = ?", id).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "cannot delete tag that contains posts")
	}

	// Delete the tag
	return r.db.Delete(tag).Error
}

// NewRepository returns a new implementation of repository.Tag
func NewRepository(db *gorm.DB) repo.Tag {
	return &repository{db: db}
}

func (r *repository) ReadAll() ([]*model.Tag, error) {
	var result []*model.Tag

	query := r.db.Model(&model.Tag{}).Find(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil

}

// Read finds and returns the tag model by id
func (r *repository) Read(id int) (*model.Tag, error) {
	var result *model.Tag

	query := r.db.Model(&model.Tag{}).First(&result, "`id` = ?", id)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// ReadByName finds and returns the tag model by name
func (r *repository) ReadByName(name string) (*model.Tag, error) {
	var result *model.Tag

	query := r.db.Model(&model.Tag{}).First(&result, "`name` = ?", name)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// Insert performs insert action into tag table
func (r *repository) Insert(o *model.Tag) (*model.Tag, error) {
	if err := r.db.Create(o).Error; err != nil {
		return nil, err
	}

	return o, nil
}

// Update performs update action into tag table
func (r *repository) Update(o *model.Tag) (*model.Tag, error) {
	if err := r.db.Save(o).Error; err != nil {
		return nil, err
	}

	return o, nil
}
