package post

import (
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

// Create a new post
func (s *service) Create(req *ct.CreatePostRequest, userID int) (*ct.PostResponse, error) {
	// Create slug based on title
	slug := s.postRepo.GenerateSlug(req.Title)
	//prepare post data
	post := &model.Post{
		Title:       req.Title,
		Body:        req.Body,
		Slug:        slug,
		IsPublished: req.IsPublished,
		UserID:      userID,
	}
	//insert post with prepared data
	res, err := s.postRepo.Insert(post)
	if err != nil {
		return nil, static.ErrInsertPost
	}

	var tags []*model.Tag
	if len(req.Tags) > 0 {
		tags, err = s.tagRepo.Select(req.Tags)
		if err != nil || len(tags) != len(req.Tags) {
			return nil, static.ErrTagNotFoundOrDeleted
		}
		// Batch insert post_tag
		postTags := make([]*model.PostTag, 0, len(req.Tags))
		for _, tag := range tags {
			postTags = append(postTags, &model.PostTag{PostID: res.ID, TagID: tag.ID})
		}
		if len(postTags) > 0 {
			if err := s.postRepo.InsertManyPostTags(postTags); err != nil {
				return nil, err
			}
		}
	}

	response := preparePostResponse(res)

	if res.UserID > 0 {
		user, err := s.userRepo.Read(res.UserID)
		if err == nil && user != nil {
			response.User = prepareProfileResponse(user)
		}
	}

	if len(tags) > 0 {
		listTagResp := &ct.ListTagResponse{Tags: make([]*ct.TagDetailResponse, 0, len(tags))}
		for _, tag := range tags {
			listTagResp.Tags = append(listTagResp.Tags, prepareTagDetailResponse(tag))
		}
		response.Tags = listTagResp
	}

	return response, nil
}
