package comment

import (
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
)

// service represents the implementation of service.Comment
type service struct {
	commentRepo repo.Comment
}

// NewService returns a new implementation of service.Comment
func NewService(commentRepo repo.Comment) svc.Comment {
	return &service{
		commentRepo: commentRepo,
	}
}
