package contract

// CommentResponse defines the structure of a single comment
// with all child comments returned in the API response.
type CommentResponse struct {
	ID              int                     `json:"id,omitempty"`
	Content         string                  `json:"content,omitempty"`
	User            *ProfileResponse        `json:"user,omitempty"`
	Post            *PostResponse           `json:"post,omitempty"`
	ChildComments   []*ChildCommentResponse `json:"child_comments,omitempty" `
	ParentCommentID *int                    `json:"parent_comment_id,omitempty"`
	CreatedAt       string                  `json:"created_at,omitempty"`
	UpdatedAt       string                  `json:"updated_at,omitempty"`
}

// ChildCommentResponse defines the structure of a single child comment returned .
type ChildCommentResponse struct {
	ID              int              `json:"id,omitempty"`
	Content         string           `json:"content,omitempty"`
	ParentCommentID *int             `json:"parent_comment_id,omitempty"`
	User            *ProfileResponse `json:"user,omitempty"`
	CreatedAt       string           `json:"created_at,omitempty"`
	UpdatedAt       string           `json:"updated_at,omitempty"`
}

// Paging response returned in the list parent comments API response
type Paging struct {
	Page     int `json:"page" `
	PageSize int `json:"page_size" `
	Total    int `json:"total" `
}

// ListCommentResponse wraps a list of comments that are
// returned when fetching all comments for a post.
type ListCommentResponse struct {
	Comments []*CommentResponse `json:"comments"`
	Paging   Paging             `json:"paging"`
}

// ListCommentRequest defines the query parameters for retrieving comments.
type ListCommentRequest struct {
	PostID   int `json:"post_id" query:"post_id" validate:"required"`
	Page     int `json:"page" query:"page"`           // Page number
	PageSize int `json:"page_size" query:"page_size"` // Number of posts per page
}

// CreateCommentRequest defines the expected payload when
// a user wants to create a new comment.
type CreateCommentRequest struct {
	Content         string `json:"content" validate:"required"`
	PostID          int    `json:"post_id" validate:"required"`
	ParentCommentID *int   `json:"parent_comment_id,omitempty"`
}

// UpdateCommentRequest defines the expected payload when
// a user wants to update an exist comment.
type UpdateCommentRequest struct {
	ID      int    `param:"commentId" swaggerignore:"true"`
	Content string `json:"content" validate:"required"`
}
