package tag

import (
	"time"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
)

// prepareTagResponse transforms tag model to response DTO
func prepareTagResponse(o *model.Tag) *ct.TagDetailResponse {
	data := &ct.TagDetailResponse{
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
func prepareListTagResponse(o []*model.Tag) *ct.ListTagResponse {
	data := &ct.ListTagResponse{
		Tags: make([]*ct.TagDetailResponse, 0, len(o)),
	}

	for _, tag := range o {
		data.Tags = append(data.Tags, prepareTagResponse(tag))
	}

	return data
}

// preparePostResponse transforms post model to response DTO
func preparePostResponse(o *model.Post) *ct.PostDetailResponse {
	data := &ct.PostDetailResponse{
		ID:          o.ID,
		Title:       o.Title,
		Body:        o.Body,
		Slug:        o.Slug,
		IsPublished: o.IsPublished,
		UserID:      o.UserID,
	}

	if o.CreatedAt != nil {
		data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
	}

	if o.UpdatedAt != nil {
		data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
	}

	return data
}

// prepareListPostResponse transforms the data and returns the List Post Response
func prepareListPostResponse(o []*model.Post) *ct.ListPostResponse {
	data := &ct.ListPostResponse{
		Posts: make([]*ct.PostDetailResponse, 0, len(o)),
	}

	for _, post := range o {
		data.Posts = append(data.Posts, preparePostResponse(post))
	}

	return data
}
