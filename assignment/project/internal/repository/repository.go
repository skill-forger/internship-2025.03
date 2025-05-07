package repository

import (
	"golang-project/internal/contract"
	"golang-project/internal/model"
)

// User represents the repository actions to the user table
type User interface {
	Read(int) (*model.User, error)
	Insert(*model.User) (*model.User, error)
	Update(*model.User, map[string]interface{}) (*model.User, error)
	ReadByEmail(string) (*model.User, error)
}

type Tag interface {
	Insert(*model.Tag) error
	Read(int) (*model.Tag, error)
	Delete(int) error
	HasPosts(int) (bool, error)
	Select([]int) ([]*model.Tag, error)
	SelectPost(int) ([]*model.Post, error)
	SelectPostTag([]int) ([]*model.PostTag, error)
	SelectUser([]int) ([]*model.User, error)
}

type Comment interface {
	Select(*contract.ListCommentRequest) ([]*model.Comment, int64, error)
	Insert(*model.Comment) (*model.Comment, error)
	Read(int) (*model.Comment, error)
	UpdateCommentByID(int, map[string]interface{}) error
}

type Post interface {
	Read(int) (*model.Post, error)
	Insert(*model.Post) (*model.Post, error)
	AddPostTags(int, []int) error
	FindSlugsLike(string) ([]string, error)
	GetTags(int) ([]*model.Tag, error)
	ReadByCondition(map[string]interface{}, ...string) (*model.Post, error)
	Select(*contract.ListPostRequest) ([]*model.Post, error)
	UpdatePost(*model.Post, map[string]interface{}) error
	UpdatePostTag(*model.Post, []*model.Tag) error
	Delete(int) error
}

// Favourite represents the repository actions for managing user follows and post favorites
type Favourite interface {
	// User following operations
	IsFollowing(userID, followUserID int) (bool, error)
	SelectFollowing(userID int) ([]*model.User, error)
	Follow(*model.FollowUser) error
	Unfollow(userID, followUserID int) error

	// Post favourite operations
	SelectFavouritePosts(userID int) ([]*model.Post, error)
}
