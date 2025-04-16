package model

// User represents user table from the database
type User struct {
	BaseModel
	FirstName    string `gorm:"type:varchar(100); CHARACTER SET utf8 COLLATE utf8_unicode_ci; not null"`
	LastName     string `gorm:"type:varchar(100); CHARACTER SET utf8 COLLATE utf8_unicode_ci; not null"`
	Email        string `gorm:"type:varchar(255); CHARACTER SET utf8 COLLATE utf8_unicode_ci; not null; uniqueIndex"`
	Password     string `gorm:"type:varchar(255); CHARACTER SET utf8 COLLATE utf8_unicode_ci; not null"`
	Pseudonym    string `gorm:"type:varchar(100); CHARACTER SET utf8 COLLATE utf8_unicode_ci; not null; uniqueIndex"`
	ProfileImage string `gorm:"type:text; CHARACTER SET utf8 COLLATE utf8_unicode_ci"`
	Biography    string `gorm:"type:text; CHARACTER SET utf8 COLLATE utf8_unicode_ci"`
	Posts        []Post
	Comments     []Comment
	FavoritePost []*Post `gorm:"many2many:favorite_post"`
	FollowUser   []*User `gorm:"many2many:follow_user;joinForeignKey:UserID;joinReferences:FollowUserID"`
	Followers    []*User `gorm:"many2many:follow_user;joinForeignKey:FollowUserID;joinReferences:UserID"`
}
