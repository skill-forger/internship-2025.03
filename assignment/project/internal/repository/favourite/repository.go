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

// Follow adds a follow relationship between user and followUser
func (r *repository) Follow(userID, followUserID int) error {
	follow := &model.FollowUser{
		UserID:       userID,
		FollowUserID: followUserID,
	}
	return r.db.Create(follow).Error
}

// Unfollow removes a follow relationship between user and followUser
func (r *repository) Unfollow(userID, followUserID int) error {
	return r.db.Where("user_id = ? AND follow_user_id = ?", userID, followUserID).Delete(&model.FollowUser{}).Error
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

// SelectFollowed returns all users that the given user is following
func (r *repository) SelectFollowed(userID int) ([]*model.User, error) {
	var users []*model.User
	err := r.db.Table("users").
		Joins("JOIN follow_user ON users.id = follow_user.follow_user_id").
		Where("follow_user.user_id = ?", userID).
		Find(&users).Error
	return users, err
}

// SelectFollowedPosts retrieves all posts from a list of user IDs
func (r *repository) SelectFollowedPosts(userIDs []int) ([]*model.Post, error) {
	var posts []*model.Post
	err := r.db.Where("user_id IN ?", userIDs).Where("is_published = ?", true).Find(&posts).Error
	return posts, err
}

// AddFavorite adds a post to a user's favorites
func (r *repository) AddFavorite(userID, postID int) error {
	favorite := &model.FavoritePost{
		UserID: userID,
		PostID: postID,
	}
	return r.db.Create(favorite).Error
}

// RemoveFavorite removes a post from a user's favorites
func (r *repository) RemoveFavorite(userID, postID int) error {
	return r.db.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&model.FavoritePost{}).Error
}

// IsFavorite checks if a post is in a user's favorites
func (r *repository) IsFavorite(userID, postID int) (bool, error) {
	var count int64
	err := r.db.Model(&model.FavoritePost{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// SelectFavorites returns all posts that a user has favorited
func (r *repository) SelectFavorites(userID int) ([]*model.Post, error) {
	var posts []*model.Post
	err := r.db.Table("posts").
		Joins("JOIN favorite_post ON posts.id = favorite_post.post_id").
		Where("favorite_post.user_id = ?", userID).
		Find(&posts).Error
	return posts, err
}
