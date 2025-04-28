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

func (s service) ListComments(req *ct.ListCommentRequest) (*ct.ListCommentResponse, error) {
	// Get comments from repository
	comments, total, err := s.commentRepo.SelectComment(req)

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

// NewService returns a new implementation of service.Comment
func NewService(commentRepo repo.Comment) svc.Comment {
	return &service{
		commentRepo: commentRepo,
	}
}
