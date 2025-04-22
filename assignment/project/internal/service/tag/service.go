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

// Delete a tag by ID
func (s *service) Delete(id int) error {
	//Check if the tag exists
	_, err := s.tagRepo.Read(id)
	if err != nil {
		return svc.ErrNotFound
	}

	//Check if tag is associated with any post
	hasPost, err := s.tagRepo.HasPost(id)
	if err != nil {
		return err
	}
	if hasPost {
		return svc.ErrTagAssociatedWithPosts
	}

	return s.tagRepo.Delete(id)
}
