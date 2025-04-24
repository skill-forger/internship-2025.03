package tag

import (
	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
	"golang-project/static"
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

// Delete a tag
func (s *service) Delete(id int) error {
	_, err := s.tagRepo.Read(id)
	if err != nil {
		return static.ErrReadTagID
	}

	hasPosts, err := s.tagRepo.HasPosts(id)
	if err != nil {
		return err
	}
	if hasPosts {
		return static.ErrHasPosts
	}

	return s.tagRepo.Delete(id)
}

// List executes all tags retrieval logic
func (s *service) List() (*ct.ListTagResponse, error) {
	tags, err := s.tagRepo.Select()
	if err != nil {
		return nil, err
	}

	return prepareListTagResponse(tags), nil
}
