package post

import (
	"fmt"

	"github.com/gosimple/slug"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
	"golang-project/static"
)

// service represents the implementation of service.Post
type service struct {
	postRepo repo.Post
	userRepo repo.User
	tagRepo  repo.Tag
	ptRepo   repo.PostTag
	// commentRepo repo.Comment
}

// NewService returns a new implementation of service.Post
func NewService(postRepo repo.Post, userRepo repo.User, tagRepo repo.Tag, ptRepo repo.PostTag) svc.Post {
	return &service{
		postRepo: postRepo,
		userRepo: userRepo,
		tagRepo:  tagRepo,
		ptRepo:   ptRepo,
		// commentRepo: CommentRepo,
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

// Update updates a post by its ID
func (s *service) Update(id int, updatePost *ct.UpdatePostRequest, tags []int) (*ct.PostResponse, error) {
	post, err := s.postRepo.Read(id)
	if err != nil {
		return nil, err
	}

	//check if the post is changed
	isChanged := false

	if updatePost.Title != post.Title {
		post.Title = updatePost.Title
		isChanged = true
	}

	if updatePost.Body != post.Body {
		post.Body = updatePost.Body
		isChanged = true
	}

	if updatePost.IsPublished != post.IsPublished {
		post.IsPublished = updatePost.IsPublished
		isChanged = true
	}

	if isChanged {
		updatePostErr := s.postRepo.Update(post)
		if updatePostErr != nil {
			return nil, updatePostErr
		}
	}

	// Uppdate the post_tag table by inserting and deleting the post_tag models
	updateTagErr := s.ptRepo.Update(id, tags)
	if updateTagErr != nil {
		return nil, updateTagErr
	}

	response := preparePostResponse(post)

	// Reload updated post
	post, err = s.postRepo.Read(id)
	if err != nil {
		return nil, err
	}

	// Add user data
	response.User = prepareProfileResponse(post.User)

	// Add tags data
	response.Tags = make([]*ct.TagResponse, len(post.Tags))
	for i, tag := range post.Tags {
		response.Tags[i] = prepareTagDetailResponse(tag)
	}

	return response, nil
}
