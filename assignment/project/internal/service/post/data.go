package post

import (
	ct "golang-project/internal/contract"
	"golang-project/internal/model"
)

// preparePostResponse chuyển đổi model.Post sang contract.PostResponse
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
		resp.CreatedAt = post.CreatedAt.Format("2006-01-02 15:04:05")
	}
	if post.UpdatedAt != nil {
		resp.UpdatedAt = post.UpdatedAt.Format("2006-01-02 15:04:05")
	}
	return resp
}

// prepareTagDetailResponse chuyển đổi model.Tag sang contract.TagDetailResponse
func prepareTagDetailResponse(tag *model.Tag) *ct.TagDetailResponse {
	if tag == nil {
		return nil
	}
	return &ct.TagDetailResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

// prepareProfileResponse chuyển đổi model.User sang contract.ProfileResponse
func prepareProfileResponse(user *model.User) ct.ProfileResponse {
	if user == nil {
		return ct.ProfileResponse{}
	}
	return ct.ProfileResponse{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Pseudonym:    user.Pseudonym,
		ProfileImage: user.ProfileImage,
		Biography:    user.Biography,
	}
}
