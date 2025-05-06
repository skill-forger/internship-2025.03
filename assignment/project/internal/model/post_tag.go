package model

// Tag represents post_tag table from the database
type PostTag struct {
	TagID  int
	PostID int
}

// TableName specifies the table name for PostTag
func (PostTag) TableName() string {
	return "post_tag"
}
