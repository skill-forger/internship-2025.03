package versions

import (
	"golang-project/internal/model"

	"gorm.io/gorm"
)

func Migrate20250416000000(db *gorm.DB) error {

	type User struct {
		model.BaseModel
		FirstName    string  `gorm:"type:varchar(100); CHARACTER SET utf8 COLLATE utf8_unicode_ci; not null"`
		LastName     string  `gorm:"type:varchar(100); CHARACTER SET utf8 COLLATE utf8_unicode_ci; not null"`
		Email        string  `gorm:"type:varchar(255); CHARACTER SET utf8 COLLATE utf8_unicode_ci; not null; uniqueIndex"`
		Password     string  `gorm:"type:varchar(255); CHARACTER SET utf8 COLLATE utf8_unicode_ci; not null"`
		Pseudonym    string  `gorm:"type:varchar(100); CHARACTER SET utf8 COLLATE utf8_unicode_ci; not null; uniqueIndex"`
		ProfileImage string  `gorm:"type:text; CHARACTER SET utf8 COLLATE utf8_unicode_ci"`
		Biography    string  `gorm:"type:text; CHARACTER SET utf8 COLLATE utf8_unicode_ci"`
		IsVerified   bool    `gorm:"type:tinyint(1); default:false"`
		FollowUser   []*User `gorm:"many2many:follow_user;joinForeignKey:UserID;joinReferences:FollowUserID"`
		Followers    []*User `gorm:"many2many:follow_user;joinForeignKey:FollowUserID;joinReferences:UserID"`
	}

	type Post struct {
		model.BaseModel
		Title      string  `gorm:"type:varchar(255); not null"`
		Body       string  `gorm:"type:text; not null"`
		Slug       string  `gorm:"type:varchar(255); not null; unique"`
		IsPublic   bool    `gorm:"type:tinyint(1); default:false"`
		UserID     uint    `gorm:"type:bigint(11); unsigned; not null"`
		User       *User   `gorm:"foreignKey:UserID"`
		FavoriteBy []*User `gorm:"many2many:favorite_post"`
	}

	type Comment struct {
		model.BaseModel
		Content         string     `gorm:"type:text; not null"`
		PostID          uint       `gorm:"type:bigint(11); unsigned; not null"`
		UserID          uint       `gorm:"type:bigint(11); unsigned; not null"`
		ParentCommentID *uint      `gorm:"type:bigint(11); unsigned"`
		User            *User      `gorm:"foreignKey:UserID"`
		Post            *Post      `gorm:"foreignKey:PostID"`
		ParentComment   *Comment   `gorm:"foreignKey:ParentCommentID"`
		ChildComments   []*Comment `gorm:"foreignKey:ParentCommentID"`
	}

	type Tag struct {
		model.BaseModel
		Name  string  `gorm:"type:varchar(100); not null"`
		Posts []*Post `gorm:"many2many:post_tag"`
	}

	return db.AutoMigrate(
		&User{},
		&Post{},
		&Comment{},
		&Tag{},
	)
}

func Rollback20250416000000(db *gorm.DB) error {
	type User struct{}

	type Post struct{}

	type Comment struct{}

	type Tag struct{}

	return db.Migrator().DropTable(
		&User{},
		&Post{},
		&Comment{},
		&Tag{},
	)
}
