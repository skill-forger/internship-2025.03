package post

import (
	ct "golang-project/internal/contract"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
)

// service represents the implementation of service.Post
type service struct {
	postRepo repo.Post
	userRepo repo.User
	tagRepo  repo.Tag
}

// NewService returns a new implementation of service.Post
func NewService(postRepo repo.Post, userRepo repo.User, tagRepo repo.Tag) svc.Post {
	return &service{
		postRepo: postRepo,
		userRepo: userRepo,
		tagRepo:  tagRepo,
	}
}

// GetByID executes the Post detail retrieval logic
func (s *service) GetByID(id int) (*ct.PostResponse, error) {
	post, err := s.postRepo.Read(id)
	if err != nil {
		return nil, err
	}

	response := preparePostResponse(post)

	// Add user data
	response.User = prepareProfileResponse(post.User)

	// Add tags data
	response.Tags = &ct.ListTagResponse{
		Tags: make([]*ct.TagDetailResponse, len(post.Tags)),
	}
	for i, tag := range post.Tags {
		response.Tags.Tags[i] = prepareTagDetailResponse(tag)
	}

	return response, nil
}
