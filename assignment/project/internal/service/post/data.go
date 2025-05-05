package post

import (
	"time"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
)

// prepareTagResponse transforms tag model to response DTO
func prepareTagResponse(o *model.Tag) *ct.TagResponse {
	data := &ct.TagResponse{
		ID:   o.ID,
		Name: o.Name,
	}

	if o.CreatedAt != nil {
		data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
	}

	if o.UpdatedAt != nil {
		data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
	}

	return data
}

// prepareListTagResponse transforms the data and returns the List Tag Response
func prepareListTagResponse(o []*model.Tag) []*ct.TagResponse {
	data := make([]*ct.TagResponse, 0, len(o))

	for _, tag := range o {
		data = append(data, prepareTagResponse(tag))
	}

	return data
}

// preparePostResponse transforms post model to response DTO
func preparePostResponse(o *model.Post, u *model.User, t []*model.Tag) *ct.PostResponse {
	data := &ct.PostResponse{
		ID:          o.ID,
		Title:       o.Title,
		Body:        o.Body,
		Slug:        o.Slug,
		IsPublished: o.IsPublished,
		User:        prepareProfileResponse(u),
		Tags:        prepareListTagResponse(t),
	}

	if o.CreatedAt != nil {
		data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
	}

	if o.UpdatedAt != nil {
		data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
	}

	return data
}

// prepareProfileResponse transforms user model to response DTO
func prepareProfileResponse(o *model.User) *ct.ProfileResponse {
	data := &ct.ProfileResponse{
		ID:           o.ID,
		FirstName:    o.FirstName,
		LastName:     o.LastName,
		Email:        o.Email,
		Pseudonym:    o.Pseudonym,
		Biography:    o.Biography,
		ProfileImage: o.ProfileImage,
	}

	if o.CreatedAt != nil {
		data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
	}

	if o.UpdatedAt != nil {
		data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
	}

	return data
}
