package post

import (
	ct "golang-project/internal/contract"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
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
