package model

// Tag represents post_tag table from the database
type PostTag struct {
	PostID int
	TagID  int
}

func (PostTag) TableName() string {
	return "post_tag"
}
