package model

// Comment represents comment table from the database
type Comment struct {
	BaseModel
	Content         string
	PostID          int
	UserID          int
	ParentCommentID *int
	User            *User      `gorm:"foreignKey:UserID"`
	Post            *Post      `gorm:"foreignKey:PostID"`
	ChildComments   []*Comment `gorm:"foreignKey:ParentCommentID"`
}
