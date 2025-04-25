package tag

import (
	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
	"golang-project/static"

	"gorm.io/gorm"
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
		if err == gorm.ErrRecordNotFound {
			return static.ErrTagNotFound
		}
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

// ListPosts executes all posts retrieval logic by tag id
func (s *service) ListPosts(id int) (*ct.ListPostResponse, error) {
	posts, err := s.tagRepo.SelectPost(id)
	if err != nil {
		return nil, err
	}

	users := make([]*model.User, 0, len(posts))
	tagsLists := make([][]*model.Tag, 0, len(posts))
	for _, post := range posts {
		// Get all the tags associated with the post
		tagsList, err := s.tagRepo.SelectByPost(post.ID)
		if err != nil {
			return nil, err
		}
		tagsLists = append(tagsLists, tagsList)

		// Get the user who created the post
		user, err := s.tagRepo.SelectUserByPost(post.ID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return prepareListPostResponse(posts, users, tagsLists), nil
}
