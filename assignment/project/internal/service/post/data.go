package post

import (
	"time"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
)

// preparePostResponse transforms the model.Post and returns the Post Response
func preparePostResponse(post *model.Post) *ct.PostResponse {
	if post == nil {
		return nil
	}

	data := &ct.PostResponse{
		ID:          post.ID,
		Title:       post.Title,
		Body:        post.Body,
		Slug:        post.Slug,
		IsPublished: post.IsPublished,
	}

	if post.CreatedAt != nil {
		data.CreatedAt = post.CreatedAt.Format(time.RFC3339)
	}
	if post.UpdatedAt != nil {
		data.UpdatedAt = post.UpdatedAt.Format(time.RFC3339)
	}

	// Convert user
	data.User = prepareProfileResponse(post.User)

	// Convert tags
	if post.Tags != nil {
		data.Tags = make([]*ct.TagResponse, len(post.Tags))
		for i, tag := range post.Tags {
			data.Tags[i] = prepareTagDetailResponse(tag)
		}
	}

	return data
}

// prepareTagDetailResponse transform model.Tag to contract.TagResponse
func prepareTagDetailResponse(tag *model.Tag) *ct.TagResponse {
	if tag == nil {
		return nil
	}
	return &ct.TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

// prepareProfileResponse transform model.User to contract.ProfileResponse
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
