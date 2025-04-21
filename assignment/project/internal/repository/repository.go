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
	Read(int) (*model.Tag, error)
	ReadAll() ([]*model.Tag, error)
	Insert(*model.Tag) (*model.Tag, error)
	Update(*model.Tag) (*model.Tag, error)
	ReadByName(string) (*model.Tag, error)
	GetPostByTagId(int) ([]*model.Post, error)
	Delete(int) error
}
