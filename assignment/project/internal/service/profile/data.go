package profile

import (
	"time"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
)

// prepareSignInResponse transforms the data and returns the Profile Response
func prepareProfileResponse(o *model.User) *ct.ProfileResponse {
	data := &ct.ProfileResponse{
		ID:           o.ID,
		FirstName:    o.FirstName,
		LastName:     o.LastName,
		Email:        o.Email,
		Pseudonym:    o.Pseudonym,
		ProfileImage: o.ProfileImage,
		Biography:    o.Biography,
	}

	if o.CreatedAt != nil {
		data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
	}

	if o.UpdatedAt != nil {
		data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
	}

	return data
}

// prepareUpdateProfile updates fields of a User model and prepares the update map
func prepareUpdateProfile(o *model.User, req *ct.UpdateProfileRequest) map[string]any {
	if req.FirstName != "" {
		o.FirstName = req.FirstName
	}
	if req.LastName != "" {
		o.LastName = req.LastName
	}
	if req.Pseudonym != "" {
		o.Pseudonym = req.Pseudonym
	}

	// ProfileImage & Biography can be empty strings
	o.ProfileImage = req.ProfileImage
	o.Biography = req.Biography

	return map[string]any{
		"first_name":    o.FirstName,
		"last_name":     o.LastName,
		"pseudonym":     o.Pseudonym,
		"profile_image": o.ProfileImage,
		"biography":     o.Biography,
	}
}

// prepareTagDetailResponse transform model.Tag to contract.TagResponse
func prepareTagDetailResponse(tag *model.Tag) *ct.TagResponse {
	if tag == nil {
		return nil
	}

	data := &ct.TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}

	if tag.CreatedAt != nil {
		data.CreatedAt = tag.CreatedAt.Format(time.RFC3339)
	}

	if tag.UpdatedAt != nil {
		data.UpdatedAt = tag.UpdatedAt.Format(time.RFC3339)
	}

	return data
}

func preparePostResponse(post *model.Post) *ct.PostResponse {
	if post == nil {
		return &ct.PostResponse{}
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
	data.Tags = make([]*ct.TagResponse, len(post.Tags))
	for i, tag := range post.Tags {
		data.Tags[i] = prepareTagDetailResponse(tag)
	}

	return data
}
