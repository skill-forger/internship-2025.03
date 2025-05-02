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

// Create executes the logic to create a new blog post
func (s *service) Create(req *ct.CreatePostRequest, userID int) (*ct.PostResponse, error) {
	// Create slug based on title
	slug := s.generateSlug(req.Title)

	// Prepare post data
	post := &model.Post{
		Title:       req.Title,
		Body:        req.Body,
		Slug:        slug,
		IsPublished: req.IsPublished,
		UserID:      userID,
	}

	// Get tags if provided
	var tagIDs []int
	if len(req.Tags) > 0 {
		tags, err := s.tagRepo.Select(req.Tags)
		if err != nil || len(tags) != len(req.Tags) {
			return nil, static.ErrTagNotFoundOrDeleted
		}
		tagIDs = req.Tags // Lưu lại để thêm liên kết sau khi tạo post
	}

	// Insert post with prepared data
	res, err := s.postRepo.Insert(post)
	if err != nil {
		return nil, static.ErrInsertPost
	}

	// Add Post-Tag associations if needed
	if len(tagIDs) > 0 {
		if err := s.postRepo.AddPostTags(res.ID, tagIDs); err != nil {
			return nil, static.ErrInsertPostTags
		}
	}

	// Fetch post with all related data in a single query
	fullPost, err := s.postRepo.ReadByCondition(
		map[string]interface{}{"id": res.ID},
		"User", "Tags",
	)
	if err != nil {
		return nil, static.ErrFetchPostDetail
	}

	// Prepare response (user, tags)
	response := preparePostResponse(fullPost)
	return response, nil
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
