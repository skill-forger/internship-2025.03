package model

// Comment represents comment table from the database
type Comment struct {
	BaseModel
	Content         string
	PostID          int
	UserID          int
	ParentCommentID *int
}
