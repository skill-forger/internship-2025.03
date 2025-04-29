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
func (s *service) Create(name string) (*ct.TagResponse, error) {
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
	tags, err := s.tagRepo.Select(nil) // If the postIDs parameter is nil, it will return all tags
	if err != nil {
		return nil, err
	}

	return &ct.ListTagResponse{Tags: prepareListTagResponse(tags)}, nil
}

// ListPosts executes all posts retrieval logic by tag id
func (s *service) ListPosts(id int) (*ct.ListPostResponse, error) {
	//Select posts by tag id
	posts, errPosts := s.tagRepo.SelectPost(id)
	if errPosts != nil {
		return nil, errPosts
	}

	// make a list of posts ID & user IDs
	postIDs := make([]int, 0, len(posts))
	userIDs := make([]int, 0, len(posts))
	for _, post := range posts {
		postIDs = append(postIDs, post.ID)
		userIDs = append(userIDs, post.UserID)
	}

	// Select all post_tags by list of posts ID
	post_tags, errPostTags := s.tagRepo.SelectPostTag(postIDs)
	if errPostTags != nil {
		return nil, errPostTags
	}

	// make a list of tags ID
	tagIDs := make([]int, 0, len(post_tags))
	for _, post_tag := range post_tags {
		tagIDs = append(tagIDs, post_tag.TagID)
	}

	// Select all tags by list of tags ID
	tags, errTags := s.tagRepo.Select(tagIDs)
	if errTags != nil {
		return nil, errTags
	}

	// Select all users by list of posts ID
	users, errUsers := s.tagRepo.SelectUser(userIDs)
	if errUsers != nil {
		return nil, errUsers
	}

	return prepareListPostResponse(posts, post_tags, tags, users), nil
}
