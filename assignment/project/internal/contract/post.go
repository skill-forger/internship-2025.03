package contract

// PostResponse defines the full details of a blog post returned by the post detail API
type PostResponse struct {
	ID        int      `json:"id,omitempty"`
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Slug      string   `json:"slug,omitempty"`
	IsPublic  bool     `json:"is_public,omitempty"`
	UserID    int      `json:"user_id,omitempty"`
	Pseudonym string   `json:"pseudonym,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	CreatedAt string   `json:"created_at,omitempty"`
	UpdatedAt string   `json:"updated_at,omitempty"`
}

// PostListResponse defines the summary information of a blog post used in list endpoints
type PostListResponse struct {
	ID        int      `json:"id,omitempty"`
	Title     string   `json:"title,omitempty"`
	Slug      string   `json:"slug,omitempty"`
	IsPublic  bool     `json:"is_public,omitempty"`
	CreatedAt string   `json:"created_at,omitempty"`
	Pseudonym string   `json:"pseudonym,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

// CreatePostRequest represents the required and optional data needed to create a new blog post
type CreatePostRequest struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
	Tags  []int  `json:"tags,omitempty"`
}

// UpdatePostRequest represents the fields that can be updated in an existing blog post
type UpdatePostRequest struct {
	Title    string `json:"title,omitempty"`
	Body     string `json:"body,omitempty"`
	Tags     []int  `json:"tags,omitempty"`
	IsPublic bool   `json:"is_public" default:"false"`
}

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
