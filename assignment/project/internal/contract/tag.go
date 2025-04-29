package contract

// TagResponse specifies the data and types for tag API response
type TagResponse struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ListTagResponse specifies the data and types for list tag API response
type ListTagResponse struct {
	Tags []*TagResponse `json:"tags,omitempty"`
}

// CreateTagRequest specifies the data and types for tag API request
type CreateTagRequest struct {
	Name string `json:"name" validate:"required"`
}
