package post

import (
	ct "golang-project/internal/contract"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
)

// service represents the implementation of service.Post
type service struct {
	postRepo    repo.Post
	userRepo    repo.User
	tagRepo     repo.Tag
	commentRepo repo.Comment
}

// NewService returns a new implementation of service.Post
func NewService(postRepo repo.Post, userRepo repo.User, tagRepo repo.Tag, CommentRepo repo.Comment) svc.Post {
	return &service{
		postRepo:    postRepo,
		userRepo:    userRepo,
		tagRepo:     tagRepo,
		commentRepo: CommentRepo,
	}
}

// Update updates a post by its ID
func (s *service) Update(id int, updatePost *ct.UpdatePostRequest, tags []int) (*ct.PostResponse, error) {
	post, existErr := s.postRepo.Read(id)
	if existErr != nil {
		return nil, existErr
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
	postTags, updateTagErr := s.postRepo.UpdatePostTag(id, tags)
	if updateTagErr != nil {
		return nil, updateTagErr
	}

	listTagIDs := make([]int, 0, len(postTags))
	for _, tag := range postTags {
		listTagIDs = append(listTagIDs, tag.TagID)
	}

	listTag, tagsErr := s.tagRepo.Select(listTagIDs)
	if tagsErr != nil {
		return nil, tagsErr
	}
	user, userErr := s.tagRepo.SelectUser([]int{post.UserID})
	if userErr != nil {
		return nil, userErr
	}

	return preparePostResponse(post, user[0], listTag), nil
}
