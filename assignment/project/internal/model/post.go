package model

// Post represents post table from the database
type Post struct {
	BaseModel
	Tiltle   string
	Body     string
	Slug     string
	IsPublic bool
	UserID   int
	Tags     []*Tag `gorm:"many2many:post_tags;"`
}
