package post

import (
	"fmt"
	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
	"golang-project/static"

	"github.com/gosimple/slug"
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
	response.Tags = make([]*ct.TagResponse, len(post.Tags))
	for i, tag := range post.Tags {
		response.Tags[i] = prepareTagDetailResponse(tag)
	}

	return response, nil
}

// generateSlug generates a URL-friendly and unique slug from the post title
func (s *service) generateSlug(title string) string {
	baseSlug := slug.Make(title)

	// call base slug
	slugs, err := s.postRepo.FindSlugsLike(baseSlug)
	if err != nil {
		return baseSlug
	}

	// return if no slug-like found
	if len(slugs) == 0 {
		return baseSlug
	}

	// count the suffix
	maxSuffix := 0
	for _, s := range slugs {
		if s == baseSlug {
			continue
		}
		var n int
		if _, err := fmt.Sscanf(s, baseSlug+"-%d", &n); err == nil && n > maxSuffix {
			maxSuffix = n
		}
	}
	return fmt.Sprintf("%s-%d", baseSlug, maxSuffix+1)
}

// List executes the Post list retrieval logic
func (s *service) List(filter *ct.ListPostRequest) (*ct.ListPostResponse, error) {
	posts, err := s.postRepo.Select(filter)
	if err != nil {
		return nil, err
	}

	responses := make([]*ct.PostResponse, len(posts))
	for i, post := range posts {
		response := preparePostResponse(post)

		// Add user data
		response.User = prepareProfileResponse(post.User)

		// Add tags data
		response.Tags = make([]*ct.TagResponse, len(post.Tags))
		for j, tag := range post.Tags {
			response.Tags[j] = prepareTagDetailResponse(tag)
		}

		responses[i] = response
	}

	return &ct.ListPostResponse{
		Posts: responses,
	}, nil
}
