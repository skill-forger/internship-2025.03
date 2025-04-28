package post

import (
	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	"time"
)

// preparePostResponse transfer data model.Post to contract.PostResponse
func preparePostResponse(post *model.Post) *ct.PostResponse {
	if post == nil {
		return nil
	}
	resp := &ct.PostResponse{
		ID:          post.ID,
		Title:       post.Title,
		Body:        post.Body,
		Slug:        post.Slug,
		IsPublished: post.IsPublished,
	}
	if post.CreatedAt != nil {
		resp.CreatedAt = post.CreatedAt.Format(time.RFC3339)
	}
	if post.UpdatedAt != nil {
		resp.UpdatedAt = post.UpdatedAt.Format(time.RFC3339)
	}
	return resp
}

// prepareTagDetailResponse transfer data model.Tag to contract.TagDetailResponse
func prepareTagDetailResponse(tag *model.Tag) *ct.TagDetailResponse {
	if tag == nil {
		return nil
	}
	return &ct.TagDetailResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

// prepareProfileResponse transfer data model.User to contract.ProfileResponse
func prepareProfileResponse(user *model.User) *ct.ProfileResponse {
	if user == nil {
		return nil
	}
	return &ct.ProfileResponse{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Pseudonym:    user.Pseudonym,
		ProfileImage: user.ProfileImage,
		Biography:    user.Biography,
	}
}
