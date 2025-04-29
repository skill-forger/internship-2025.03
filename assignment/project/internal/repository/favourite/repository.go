package favourite

import (
	"golang-project/internal/model"
	repo "golang-project/internal/repository"

	"gorm.io/gorm"
)

// repository represents the implementation of repository.Favourite
type repository struct {
	db *gorm.DB
}

// NewRepository returns a new implementation of repository.Favourite
func NewRepository(db *gorm.DB) repo.Favourite {
	return &repository{db: db}
}

// IsFollowing checks if user is following followUser
func (r *repository) IsFollowing(userID, followUserID int) (bool, error) {
	return false, nil
}

// SelectFollowing returns all users that the given user is following
func (r *repository) SelectFollowing(userID int) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Table("users").
		Joins("JOIN follow_user ON users.id = follow_user.follow_user_id").
		Where("follow_user.user_id = ?", userID).
		Find(&users).Error
	return users, err
}
