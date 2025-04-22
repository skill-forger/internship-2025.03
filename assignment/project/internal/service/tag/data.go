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
		Tags: make([]ct.TagDetailResponse, len(o)),
	}

	for i, tag := range o {
		data.Tags[i] = ct.TagDetailResponse{
			ID:   tag.ID,
			Name: tag.Name,
		}

		if tag.CreatedAt != nil {
			data.Tags[i].CreatedAt = tag.CreatedAt.Format(time.RFC3339)
		}
		if tag.UpdatedAt != nil {
			data.Tags[i].UpdatedAt = tag.UpdatedAt.Format(time.RFC3339)
		}
	}

	return data
}
