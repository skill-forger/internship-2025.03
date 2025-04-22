package contract

// TagDetailResponse represents the detailed information of a tag entity used in the response of GetTagByID.
type TagDetailResponse struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ListTagResponse wraps a list of tag details returned from the ListTags API.
type ListTagResponse struct {
	Tags []TagDetailResponse `json:"tags"`
}

// CreateTagRequest defines the structure for creating a new tag, including the validation rule .
type CreateTagRequest struct {
	Name string `json:"name" validate:"required"`
}
