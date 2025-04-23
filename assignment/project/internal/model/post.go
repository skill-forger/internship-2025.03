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
