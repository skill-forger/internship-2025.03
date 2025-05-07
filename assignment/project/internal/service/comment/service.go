package comment

import (
	"errors"

	"gorm.io/gorm"

	ct "golang-project/internal/contract"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
	"golang-project/static"
)

// service represents the implementation of service.Comment
type service struct {
	commentRepo repo.Comment
	userRepo    repo.User
	postRepo    repo.Post
}

// NewService returns a new implementation of service.Comment
func NewService(commentRepo repo.Comment, userRepo repo.User, postRepo repo.Post) svc.Comment {
	return &service{
		commentRepo: commentRepo,
		userRepo:    userRepo,
		postRepo:    postRepo,
	}
}

// List executes the Comment list retrieval logic
func (s *service) List(req *ct.ListCommentRequest) (*ct.ListCommentResponse, error) {
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

func (s *service) Create(request *ct.CreateCommentRequest, userID int) (*ct.CommentResponse, error) {
	// Get post info to validate post exists
	post, err := s.postRepo.Read(request.PostID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, static.ErrPostNotFound
		}
		return nil, err
	}

	// If this is a reply to another comment, validate the parent comment
	if request.ParentCommentID != nil {
		parentComment, err := s.commentRepo.Read(*request.ParentCommentID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, static.ErrCommentNotFound
			}
			return nil, err
		}

		// Check if parent comment belongs to the same post
		if parentComment.PostID != request.PostID {
			return nil, errors.New("parent comment does not belong to the same post")
		}

		// Check if parent comment is a root comment
		if parentComment.ParentCommentID != nil {
			return nil, errors.New("can only reply to root comments")
		}
	}

	// Create new comment model
	comment := prepareCommentModel(request, userID, post)

	// Save to database
	createdComment, err := s.commentRepo.Insert(comment)
	if err != nil {
		return nil, err
	}

	// Get the created comment with preloaded user and post info
	commentPreloadResponse, err := s.commentRepo.Read(createdComment.ID)
	if err != nil {
		return nil, err
	}

	return prepareCommentResponse(commentPreloadResponse), nil
}

func (s *service) Update(req *ct.UpdateCommentRequest, userID int) (*ct.CommentResponse, error) {
	// Get comment by ID with preloaded data
	comment, err := s.commentRepo.Read(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, static.ErrCommentNotFound
		}
		return nil, err
	}

	// Check if user is the owner of the comment
	if comment.UserID != userID {
		return nil, static.ErrUserPermission
	}

	// Update comment
	updateCommentErr := s.commentRepo.UpdateCommentByID(req.ID, prepareUpdateComment(comment, req))
	if updateCommentErr != nil {
		return nil, err
	}

	return prepareCommentResponse(comment), nil
}
