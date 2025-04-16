package model

// User represents user table from the database
type Comment struct {
	BaseModel
	Content         string `gorm:"type:text; not null"`
	PostID          int    `gorm:"index"`
	UserID          int    `gorm:"index"`
	ParentCommentID *int   `gorm:"index"`
	User            User
	Post            Post
	ParentComment   *Comment   `gorm:"foreignKey:ParentCommentID"`
	ChildComments   []*Comment `gorm:"foreignKey:ParentCommentID"`
}
