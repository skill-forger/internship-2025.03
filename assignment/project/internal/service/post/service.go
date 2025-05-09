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
func (s *service) Update(ctxUserID int, updatePost *ct.UpdatePostRequest) (*ct.PostResponse, error) {
	post, err := s.postRepo.ReadByCondition(map[string]any{
		"id": updatePost.ID,
	})
	if err != nil {
		return nil, err
	}

	// Check ctxUser permission to update
	if ctxUserID != post.UserID {
		return nil, static.ErrUserPermission
	}

	tags, err := s.tagRepo.Select(updatePost.Tags)
	if err != nil {
		return nil, err
	}

	updatePostErr := s.postRepo.UpdatePost(post, prepareUpdateMap(updatePost))
	if updatePostErr != nil {
		return nil, updatePostErr
	}

	updatePostTagErr := s.postRepo.UpdatePostTag(post, tags)
	if updatePostTagErr != nil {
		return nil, updatePostTagErr
	}

	// Reload updated post
	post, err = s.postRepo.ReadByCondition(map[string]any{
		"id": post.ID,
	}, "User", "Tags")
	if err != nil {
		return nil, err
	}

	return preparePostResponse(post), nil
}

// Delete deletes a post by its ID
func (s *service) Delete(postID, ctxUserID int) error {
	post, err := s.postRepo.ReadByCondition(map[string]any{
		"id": postID,
	})
	if err != nil {
		return err
	}

	// Check ctxUser permission to update
	if ctxUserID != post.UserID {
		return static.ErrUserPermission
	}

	return s.postRepo.Delete(postID)
}
