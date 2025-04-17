package model

// User represents user table from the database
type Comment struct {
	BaseModel
	Content         string
	PostID          int
	UserID          int
	ParentCommentID *int
}
