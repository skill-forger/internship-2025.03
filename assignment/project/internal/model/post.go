package model

// Post represents post table from the database
type Post struct {
	BaseModel
	Title       string
	Body        string
	Slug        string
	IsPublished bool
	UserID      int
}

// PostTag represents the many-to-many relationship between posts and tags
type PostTag struct {
	PostID int `gorm:"primaryKey"`
	TagID  int `gorm:"primaryKey"`
}

func (PostTag) TableName() string {
	return "post_tag"
}
