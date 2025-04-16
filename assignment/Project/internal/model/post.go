package model

// User represents user table from the database
type Post struct {
	BaseModel
	Tiltle     string `gorm:"type:varchar(255); not null"`
	Body       string `gorm:"type:text; not null"`
	Slug       string `gorm:"type:varchar(255);not null; uniqueIndex"`
	IsPublic   bool   `gorm:"type:bool; default:false"`
	UserID     int    `gorm:"index"`
	User       User
	Comments   []Comment
	Tag        []*Tag  `gorm:"many2many:post_tag"`
	FavoriteBy []*User `gorm:"many2many:favorite_post"`
}
