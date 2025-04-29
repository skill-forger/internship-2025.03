package repository

import (
	"golang-project/internal/model"
)

// User represents the repository actions to the user table
type User interface {
	Read(int) (*model.User, error)
	Insert(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	ReadByEmail(string) (*model.User, error)
}

type Tag interface {
	Insert(*model.Tag) error
	Read(int) (*model.Tag, error)
	Delete(int) error
	HasPosts(int) (bool, error)
	Select() ([]*model.Tag, error)
}

type Comment interface {
}

// Favourite represents the repository actions for managing user follows and post favorites
type Favourite interface {
	// User following operations
	Follow(userID, followUserID int) error
	Unfollow(userID, followUserID int) error
	IsFollowing(userID, followUserID int) (bool, error)
	SelectFollowed(userID int) ([]*model.User, error)
	SelectFollowedPosts(userIDs []int) ([]*model.Post, error)

	// Post favorite operations
	AddFavorite(userID, postID int) error
	RemoveFavorite(userID, postID int) error
	IsFavorite(userID, postID int) (bool, error)
	SelectFavorites(userID int) ([]*model.Post, error)
}
