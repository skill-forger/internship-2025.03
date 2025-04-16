package model

type Tag struct {
	BaseModel
	Name  string  `gorm:"type:varchar(100); not null"`
	Posts []*Post `gorm:"many2many:post_tag"`
}
