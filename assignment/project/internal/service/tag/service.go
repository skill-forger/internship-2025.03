package tag

import (
	ct "golang-project/internal/contract"
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

// GetAllTags executes all tags retrieval logic
func (s *service) GetAllTags() (*ct.ListTagResponse, error) {
	tags, err := s.tagRepo.ReadAll()
	if err != nil {
		return nil, err
	}

	return prepareListTagResponse(tags), nil
}
