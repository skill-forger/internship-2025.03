package contract

import (
	"golang-project/static"
)

// PostResponse defines the full details of a blog post returned by the post detail API,
type PostResponse struct {
	ID          int              `json:"id,omitempty"`
	Title       string           `json:"title,omitempty"`
	Body        string           `json:"body,omitempty"`
	Slug        string           `json:"slug,omitempty"`
	IsPublished bool             `json:"is_published,omitempty"`
	User        *ProfileResponse `json:"user,omitempty"`
	Tags        []*TagResponse   `json:"tags,omitempty"`
	CreatedAt   string           `json:"created_at,omitempty"`
	UpdatedAt   string           `json:"updated_at,omitempty"`
}

// ListPostResponse defines the summary information of a blog post used in list endpoints,
type ListPostResponse struct {
	Posts []*PostResponse `json:"posts"`
}

// CreatePostRequest represents the required and optional data needed to create a new blog post.
type CreatePostRequest struct {
	IsPublished bool   `json:"is_published,omitempty" default:"false"`
	Title       string `json:"title" validate:"required"`
	Body        string `json:"body" validate:"required"`
	Tags        []int  `json:"tags,omitempty"`
}

// UpdatePostRequest represents the fields that can be updated in an existing blog post.
type UpdatePostRequest struct {
	Title       string `json:"title,omitempty"`
	Body        string `json:"body,omitempty"`
	Tags        []int  `json:"tags,omitempty"`
	IsPublished bool   `json:"is_published"`
}

// ListPostRequest defines the filter parameters for retrieving posts.
type ListPostRequest struct {
	Tag       string `query:"tag"`
	Pseudonym string `query:"pseudonym"`
	Title     string `query:"title"`
	Page      int    `query:"page"`      // Page number
	PageSize  int    `query:"page_size"` // Number of posts per page
}

// PostFavouriteStatusResponse represents the response when a user marks/unmarks a post as favourite,
// containing the post ID and the current favourite status
type PostFavouriteStatusResponse struct {
	PostID      int  `json:"post_id"`
	IsFavourite bool `json:"is_favourite"`
}

// PostFavouriteRequest represents the request payload for add to/remove from favourites actions
type PostFavouriteRequest struct {
	Action static.PostFavouriteAction `json:"action" validate:"required,oneof=favourite unfavourite"`
	PostID int                        `json:"post_id"`
}
