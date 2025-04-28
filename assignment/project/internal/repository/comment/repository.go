package comment

import (
	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	"gorm.io/gorm"
)

// repository represents the implementation of repository.Comment
type repository struct {
	db *gorm.DB
}

func (r repository) SelectComment(request *ct.ListCommentRequest) ([]*model.Comment, int64, error) {

	var comments []*model.Comment
	var total int64

	// Count total comments
	if err := r.db.Model(&model.Comment{}).
		Where("post_id = ? AND parent_comment_id IS NULL", request.PostID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated parent comments with all relationships
	offset := (request.Page - 1) * request.PageSize

	query := r.db.
		Preload("User").
		Preload("Post").
		Preload("Post.User").
		Preload("Post.Tags").
		Preload("ChildComments").
		Where("post_id = ? AND parent_comment_id IS NULL", request.PostID).
		Order("created_at DESC").
		Offset(offset).
		Limit(request.PageSize)

	if err := query.Find(&comments).Error; err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

// NewRepository returns a new implementation of repository.Comment
func NewRepository(db *gorm.DB) repo.Comment {
	return &repository{db: db}
}
