package post

import (
	"fmt"
	ct "golang-project/internal/contract"
	"golang-project/internal/model"
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

// Create a new post
func (s *service) Create(req *ct.CreatePostRequest) (*ct.PostResponse, error) {
	// Tạo slug
	slug, err := s.postRepo.GenerateSlug(req.Title)
	if err != nil {
		return nil, fmt.Errorf("không thể tạo slug: %w", err)
	}

	post := &model.Post{
		Title:       req.Title,
		Body:        req.Body,
		Slug:        slug,
		IsPublished: req.IsPublished,
		UserID:      req.UserID,
	}

	res, err := s.postRepo.Insert(post)
	if err != nil {
		return nil, err
	}

	// Process tags if provided
	if len(req.Tags) > 0 {
		fmt.Printf("Bắt đầu thêm %d tags cho post ID %d\n", len(req.Tags), res.ID)
		for _, tagID := range req.Tags {
			fmt.Printf("Đang thêm tag ID %d vào post ID %d\n", tagID, res.ID)
			err := s.postRepo.AddPostTag(res.ID, tagID)
			if err != nil {
				fmt.Printf("Lỗi khi thêm tag %d vào post %d: %v\n", tagID, res.ID, err)
			}
		}
	}

	response := preparePostResponse(res)

	// Add user information
	if res.UserID > 0 {
		user, err := s.userRepo.Read(res.UserID)
		if err == nil && user != nil {
			response.User = prepareProfileResponse(user)
		}
	}

	// Add tags if available
	tags, err := s.postRepo.GetTags(res.ID)
	if err == nil && len(tags) > 0 {
		response.Tags = make([]*ct.TagDetailResponse, 0, len(tags))
		for _, tag := range tags {
			response.Tags = append(response.Tags, prepareTagDetailResponse(tag))
		}
	}

	return response, nil
}
