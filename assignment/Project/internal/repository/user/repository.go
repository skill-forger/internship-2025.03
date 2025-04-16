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

	query := r.db.Model(&model.User{}).First(&result, "`email` = ?", email)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// Insert performs insert action into user table
func (r *repository) Insert(o *model.User) (*model.User, error) {
	return nil, nil
}

// Update performs update action into user table
func (r *repository) Update(o *model.User) (*model.User, error) {
	return nil, nil
}
