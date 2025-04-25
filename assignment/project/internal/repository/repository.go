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
	SelectPost(int) ([]*model.Post, error)
	SelectByPost(int) ([]*model.Tag, error)
	SelectUserByPost(int) (*model.User, error)
}

type Comment interface {
}
