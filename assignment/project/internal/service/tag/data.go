package tag

import (
	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	"time"
)

// prepareTagResponse transforms the data and returns the Tag Response
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

// preparePostResponse transforms the data and returns the Post Response
func preparePostResponse(o *model.Post) *ct.PostResponse {
	data := &ct.PostResponse{
		ID:     o.ID,
		Title:  o.Tiltle,
		Body:   o.Body,
		UserID: o.UserID,
	}

	// Add tags if available
	if o.Tags != nil {
		tagId := make([]int, 0, len(o.Tags))
		for _, tag := range o.Tags {
			tagId = append(tagId, tag.ID)
		}
		data.Tags = tagId
	}

	if o.CreatedAt != nil {
		data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
	}

	if o.UpdatedAt != nil {
		data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
	}

	return data
}
