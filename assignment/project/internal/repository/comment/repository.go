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

// NewRepository returns a new implementation of repository.Comment
func NewRepository(db *gorm.DB) repo.Comment {
	return &repository{db: db}
}

func (r *repository) Select(request *ct.ListCommentRequest) ([]*model.Comment, int64, error) {
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

// Read finds and returns the comment model by id
func (r *repository) Read(id int) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// Insert creates a new comment in the database
func (r *repository) Insert(comment *model.Comment) (*model.Comment, error) {
	if err := r.db.Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}
