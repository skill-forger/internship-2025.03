package favourite

import (
	"time"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
)

// prepareProfileResponse converts a User model to a ProfileResponse
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

// prepareListProfileResponse converts slice of User models to ListProfileResponse
func prepareListProfileResponse(o []*model.User) *ct.ListProfileResponse {
	data := &ct.ListProfileResponse{
		Bloggers: make([]*ct.ProfileResponse, 0, len(o)),
	}

	for _, user := range o {
		data.Bloggers = append(data.Bloggers, prepareProfileResponse(user))
	}

	return data
}

// preparePostResponse converts a Post model to a PostResponse
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
	if post.User != nil {
		data.User = prepareProfileResponse(post.User)
	}

	// Convert tags
	if post.Tags != nil {
		data.Tags = make([]*ct.TagResponse, len(post.Tags))
		for i, tag := range post.Tags {
			data.Tags[i] = prepareTagResponse(tag)
		}
	}

	return data
}

// prepareTagResponse transforms a Tag model to TagResponse
func prepareTagResponse(tag *model.Tag) *ct.TagResponse {
	if tag == nil {
		return nil
	}
	return &ct.TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

// prepareListPostResponse converts a slice of Post models to ListPostResponse
func prepareListPostResponse(posts []*model.Post) *ct.ListPostResponse {
	data := &ct.ListPostResponse{
		Posts: make([]*ct.PostResponse, 0, len(posts)),
	}

	for _, post := range posts {
		data.Posts = append(data.Posts, preparePostResponse(post))
	}

	return data
}