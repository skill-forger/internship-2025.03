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

// GetPostsByID executes the logic to retrieve all posts for a tag
func (s *service) GetPostsByID(id int) ([]*ct.PostResponse, error) {
	posts, err := s.tagRepo.GetPostByTagId(id)
	if err != nil {
		return nil, err
	}

	response := make([]*ct.PostResponse, 0, len(posts))
	for _, post := range posts {
		//response = append(response, profile.PrepareProfileResponse(post))
		response = append(response, preparePostResponse(post))
	}

	return response, nil
}

func (s *service) Create(request *ct.TagRequest) (*ct.TagResponse, error) {
	tag := &model.Tag{
		Name: request.Name,
	}

	result, err := s.tagRepo.Insert(tag)
	if err != nil {
		return nil, err
	}

	return prepareTagResponse(result), nil
}

func (s *service) Delete(id int) error {
	return s.tagRepo.Delete(id)
}

// NewService returns a new implementation of service.Tag
func NewService(tagRepo repo.Tag) svc.Tag {
	return &service{
		tagRepo: tagRepo,
	}
}

func (s *service) GetAll() ([]*ct.TagResponse, error) {

	tags, err := s.tagRepo.ReadAll()

	if err != nil {
		return nil, err
	}

	response := make([]*ct.TagResponse, 0, len(tags))
	for _, tag := range tags {
		response = append(response, prepareTagResponse(tag))
	}

	return response, nil
}

// GetByID executes the tag detail retrieval logic
func (s *service) GetByID(id int) (*ct.TagResponse, error) {
	tag, err := s.tagRepo.Read(id)
	if err != nil {
		return nil, err
	}

	return prepareTagResponse(tag), nil
}
