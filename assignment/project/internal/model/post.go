package model

// User represents user table from the database
type Post struct {
	BaseModel
	Tiltle   string
	Body     string
	Slug     string
	IsPublic bool
	UserID   int
}
