package comment

import (
	"gorm.io/gorm"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	"golang-project/util/pagination"
)

// repository represents the implementation of repository.Comment
type repository struct {
	db *gorm.DB
}

func (r repository) Select(request *ct.ListCommentRequest) ([]*model.Comment, int64, error) {
	var total int64

	// Count total parent comments
	if err := r.db.Model(&model.Comment{}).
		Where("post_id = ? AND parent_comment_id is null", request.PostID).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get all comments for this post
	var allComments []*model.Comment

	if err := r.db.
		Preload("User").
		Preload("Post").
		Preload("Post.User").
		Preload("Post.Tags").
		Preload("ChildComments").
		Preload("ChildComments.User").
		Where("post_id = ? AND parent_comment_id is null", request.PostID).
		Order("created_at DESC").
		Offset(pagination.CalculateOffset(request.Page, request.PageSize)).
		Limit(request.PageSize).
		Find(&allComments).
		Error; err != nil {
		return nil, 0, err
	}

	return allComments, total, nil
}

// NewRepository returns a new implementation of repository.Comment
func NewRepository(db *gorm.DB) repo.Comment {
	return &repository{db: db}
}
