package model

// Tag represents tag table from the database
type Tag struct {
	BaseModel
	Name string
	Posts []*Post `gorm:"many2many:post_tag"`
}
