package contract

// PostDetailResponse specifies the data and types for post API response
type PostDetailResponse struct {
	ID        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
	Tags      []int  `json:"tags,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ListPostResponse specifies the data and types for list post API response
type ListPostResponse struct {
	Posts []PostDetailResponse `json:"posts"`
}
