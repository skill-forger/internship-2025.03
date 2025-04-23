package model

// Post represents post table from the database
type Post struct {
	BaseModel
	Title    string
	Body     string
	Slug     string
	IsPublic bool
	UserID   int
}
