package tag

import (
	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
)

// service represents the implementation of service.Tag
type service struct {
	tagRepo repo.Tag
}

// NewService returns a new implementation of service.Tag
func NewService(tagRepo repo.Tag) svc.Tag {
	return &service{
		tagRepo: tagRepo,
	}
}

// Create a new tag
func (s *service) Create(name string) (*ct.TagDetailResponse, error) {
	tag := &model.Tag{
		Name: name,
	}

	if err := s.tagRepo.Insert(tag); err != nil {
		return nil, err
	}

	return prepareTagResponse(tag), nil
}

// List executes all tags retrieval logic
func (s *service) List() (*ct.ListTagResponse, error) {
	tags, err := s.tagRepo.Select()
	if err != nil {
		return nil, err
	}

	return prepareListTagResponse(tags), nil
}

// ListPosts executes all posts retrieval logic by tag id
func (s *service) ListPosts(id int) (*ct.ListPostResponse, error) {
	posts, err := s.tagRepo.SelectPost(id)
	if err != nil {
		return nil, err
	}

	return prepareListPostResponse(posts), nil
}
