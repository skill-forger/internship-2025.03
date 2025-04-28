package post

import (
	"fmt"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

// repository represents the implementation of repository.Post
type repository struct {
	db *gorm.DB
}

// NewRepository returns a new implementation if repository.Post
func NewRepository(db *gorm.DB) repo.Post {
	return &repository{db: db}
}

func (r *repository) Insert(post *model.Post) (*model.Post, error) {
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

// GenerateSlug generates an unique slug from the post title
func (r *repository) GenerateSlug(title string) string {
	base := slug.Make(title)

	var total int64
	r.db.Model(&model.Post{}).
		Where("slug = ? OR slug LIKE ?", base, base+"-%").
		Count(&total)

	if total == 0 {
		return base
	}
	return fmt.Sprintf("%s-%d", base, total+1)
}

// GetTags retrieves all tags associated with a post
func (r *repository) GetTags(postID int) ([]*model.Tag, error) {
	var tags []*model.Tag

	err := r.db.Table("tags").
		Joins("JOIN post_tag ON tags.id = post_tag.tag_id").
		Where("post_tag.post_id = ?", postID).
		Find(&tags).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}

// AddPostTag associates a tag with a post
func (r *repository) AddPostTag(postID int, tagID int) error {
	postTag := &model.PostTag{
		PostID: postID,
		TagID:  tagID,
	}

	return r.db.Create(postTag).Error
}

// InsertManyPostTags batch insert post_tag
func (r *repository) InsertManyPostTags(postTags []*model.PostTag) error {
	return r.db.Create(&postTags).Error
}

//import (
//"fmt"
//"github.com/gosimple/slug"
//"github.com/oklog/ulid/v2"
//)
//
//// Trả về slug duy nhất kiểu `hello-world` hoặc `hello-world-01HJK7`
//func (r *repository) GenerateSlugULID(title string) string {
//	base := slug.Make(title)                 // 1) slugify
//
//	// 2) Nếu slug gốc chưa có, dùng ngay
//	if !r.slugExists(base) {                 // SELECT count(*) ...
//		return base
//	}
//
//	// 3) Đụng hàng ⇒ gắn 6 ký tự ULID, gần như 0 % va chạm
//	uid := ulid.Make().String()[:6]          // vd 01HJK7
//	return fmt.Sprintf("%s-%s", base, uid)   // hello-world-01HJK7
//}
