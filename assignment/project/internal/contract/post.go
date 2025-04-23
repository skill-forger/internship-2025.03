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

// PostResponse defines the full details of a blog post returned by the post detail API,
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

// ListPostResponse defines the summary information of a blog post used in list endpoints,
type ListPostResponse struct {
	Posts []*PostResponse `json:"posts"`
}

// CreatePostRequest represents the required and optional data needed to create a new blog post.
type CreatePostRequest struct {
	Title string   `json:"title" validate:"required"`
	Body  string   `json:"body" validate:"required"`
	Tags  []string `json:"tags,omitempty"`
}

// UpdatePostRequest represents the fields that can be updated in an existing blog post.
type UpdatePostRequest struct {
	Title    string   `json:"title,omitempty"`
	Body     string   `json:"body,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	IsPublic bool     `json:"is_public" default:"false"`
}

// PostQuery defines the filter parameters for retrieving posts.
type PostQuery struct {
	Tag       string `query:"tag"`
	Pseudonym string `query:"pseudonym"`
	Title     string `query:"title"`
}
