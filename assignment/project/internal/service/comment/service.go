package comment

import (
	ct "golang-project/internal/contract"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
)

// service represents the implementation of service.Comment
type service struct {
	commentRepo repo.Comment
}

// NewService returns a new implementation of service.Comment
func NewService(commentRepo repo.Comment) svc.Comment {
	return &service{
		commentRepo: commentRepo,
	}
}

// List executes the Comment list retrieval logic
func (s service) List(req *ct.ListCommentRequest) (*ct.ListCommentResponse, error) {
	// Get comments from repository
	comments, total, err := s.commentRepo.Select(req)

	if err != nil {
		return nil, err
	}

	pagingResponse := ct.Paging{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    int(total),
	}

	return prepareListCommentResponse(comments, pagingResponse), nil
}
