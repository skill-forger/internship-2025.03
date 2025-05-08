package user

import (
	"gorm.io/gorm"

	"golang-project/internal/model"
	repo "golang-project/internal/repository"
)

// repository represents the implementation of repository.User
type repository struct {
	db *gorm.DB
}

// NewRepository returns a new implementation of repository.User
func NewRepository(db *gorm.DB) repo.User {
	return &repository{db: db}
}

// Read finds and returns the user model by email
func (r *repository) Read(id int) (*model.User, error) {
	var result *model.User

	query := r.db.Model(&model.User{}).First(&result, "`id` = ?", id)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// ReadByEmail finds and returns the user model by email
func (r *repository) ReadByEmail(email string) (*model.User, error) {
	var result *model.User

	query := r.db.Model(&model.User{}).First(&result, "email = ?", email)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// Insert performs insert action into user table
func (r *repository) Insert(o *model.User) (*model.User, error) {
	if err := r.db.Create(o).Error; err != nil {
		return nil, err
	}
	return o, nil
}

// Update performs update action into user table
func (r *repository) Update(o *model.User, updates map[string]interface{}) (*model.User, error) {
	query := r.db.Model(&model.User{}).Where("id = ?", o.ID).Updates(updates)

	if err := query.Error; err != nil {
		return nil, err
	}

	return o, nil
}

func (r *repository) ReadOwnPosts(id int, isPublishedFilter *bool) ([]*model.Post, error) {
	var posts []*model.Post

	// Start building the query
	query := r.db.Where("user_id = ?", id)

	// Add filter by publication status if provided
	if isPublishedFilter != nil {
		query = query.Where("is_published = ?", *isPublishedFilter)
	}

	// Execute the query with preloads
	if err := query.
		Preload("Tags").
		Preload("User").
		Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}
