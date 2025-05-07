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
	var count int64
	err := r.db.Model(&model.FollowUser{}).Where("user_id = ? AND follow_user_id = ?", userID, followUserID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
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

// Follow adds a follow relationship between user and followUser
func (r *repository) Follow(follow *model.FollowUser) error {
	return r.db.Create(follow).Error
}

// Unfollow removes a follow relationship between user and followUser
func (r *repository) Unfollow(userID, followUserID int) error {
	return r.db.Where("user_id = ? AND follow_user_id = ?", userID, followUserID).Delete(&model.FollowUser{}).Error
}

// SelectFavouritePosts returns all posts that the given user has marked as favourite
func (r *repository) SelectFavouritePosts(userID int) ([]*model.Post, error) {
	var posts []*model.Post
	err := r.db.Model(&model.Post{}).
		Preload("User").
		Preload("Tags", "deleted_at IS NULL").
		Joins("JOIN favorite_post ON posts.id = favorite_post.post_id").
		Where("favorite_post.user_id = ? AND posts.is_published = ?", userID, true).
		Find(&posts).Error
	return posts, err
}