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
	ID          int                 `json:"id,omitempty"`
	Title       string              `json:"title,omitempty"`
	Body        string              `json:"body,omitempty"`
	Slug        string              `json:"slug,omitempty"`
	IsPublished bool                `json:"is_published,omitempty"`
	User        ProfileResponse     `json:"user,omitempty"`
	Tags        []TagDetailResponse `json:"tags,omitempty"`
	CreatedAt   string              `json:"created_at,omitempty"`
	UpdatedAt   string              `json:"updated_at,omitempty"`
}

// ListPostResponse defines the summary information of a blog post used in list endpoints,
type ListPostResponse struct {
	Posts []*PostResponse `json:"posts"`
}

// CreatePostRequest represents the required and optional data needed to create a new blog post.
type CreatePostRequest struct {
	UserID      int                 `json:"user_id"`
	IsPublished bool                `json:"is_published,omitempty" default:"false"`
	Title       string              `json:"title" validate:"required"`
	Body        string              `json:"body" validate:"required"`
	Tags        []TagDetailResponse `json:"tags,omitempty"`
}

// ListPostRequest represents the fields that can be updated in an existing blog post.
type ListPostRequest struct {
	Title       string              `json:"title,omitempty"`
	Body        string              `json:"body,omitempty"`
	Tags        []TagDetailResponse `json:"tags,omitempty"`
	IsPublished bool                `json:"is_published"`
}

// PostQuery defines the filter parameters for retrieving posts.
type PostQuery struct {
	Tag       string `query:"tag"`
	Pseudonym string `query:"pseudonym"`
	Title     string `query:"title"`
	Limit     int    `query:"limit"`  // Number of posts to return
	Offset    int    `query:"offset"` // Starting point for pagination
}
