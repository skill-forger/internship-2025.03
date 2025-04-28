package model

// Post represents post table from the database
type Post struct {
	BaseModel
	Title       string
	Body        string
	Slug        string
	IsPublished bool
	UserID      int
	User        *User
	Tags        []*Tag `gorm:"many2many:post_tag"`
}
