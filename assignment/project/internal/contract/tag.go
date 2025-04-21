package contract

// TagDetailResponse specifies the data and types for tag API response
type TagDetailResponse struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ListTagResponse specifies the data and types for list tag API response
type ListTagResponse struct {
	Tags []TagDetailResponse `json:"tags"`
}

// CreateTagRequest specifies the data and types for tag API request
type CreateTagRequest struct {
	Name string `json:"name" validate:"required"`
}
