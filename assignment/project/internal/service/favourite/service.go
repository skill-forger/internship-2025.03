package favourite

import (
	"errors"
	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
	"golang-project/static"

	"gorm.io/gorm"
)

// service represents the implementation of service.Favourite
type service struct {
	favouriteRepo repo.Favourite
	userRepo      repo.User
	postRepo      repo.Post
}

// NewService returns a new implementation of service.Favourite
func NewService(favouriteRepo repo.Favourite, userRepo repo.User, postRepo repo.Post) svc.Favourite {
	return &service{
		favouriteRepo: favouriteRepo,
		userRepo:      userRepo,
		postRepo:      postRepo,
	}
}

// ListFollowingUsers returns a list of all bloggers that a user is following
func (s *service) ListFollowingUsers(userID int) (*ct.ListProfileResponse, error) {
	// Check if user exists
	_, err := s.userRepo.Read(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, static.ErrUserNotFound
		}
		return nil, static.ErrDatabaseOperation
	}

	users, err := s.favouriteRepo.SelectFollowing(userID)
	if err != nil {
		return nil, static.ErrDatabaseOperation
	}

	return prepareListProfileResponse(users), nil
}

// UpdateFollowStatus handles follow/unfollow operations
func (s *service) UpdateFollowStatus(userID int, req *ct.BloggerFollowRequest) (*ct.BloggerFollowStatusResponse, error) {
	targetUserID := req.UserID
	isFollow := req.Action == static.Follow

	// Check if target user exists
	_, err := s.userRepo.Read(targetUserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, static.ErrUserNotFound
		}
		return nil, static.ErrDatabaseOperation
	}

	// Prevent self-following
	if userID == targetUserID {
		return nil, static.ErrSelfFollow
	}

	// Check current follow status
	isFollowing, err := s.favouriteRepo.IsFollowing(userID, targetUserID)
	if err != nil {
		return nil, static.ErrDatabaseOperation
	}

	// Handle follow/unfollow based on current status
	if isFollow {
		if isFollowing {
			return nil, static.ErrAlreadyFollowing
		}

		follow := &model.FollowUser{
			UserID:       userID,
			FollowUserID: targetUserID,
		}
		err = s.favouriteRepo.Follow(follow)
	} else {
		if !isFollowing {
			return nil, static.ErrNotFollowing
		}
		err = s.favouriteRepo.Unfollow(userID, targetUserID)
	}

	if err != nil {
		return nil, static.ErrFollowStatusUpdate
	}

	return &ct.BloggerFollowStatusResponse{
		UserID:      targetUserID,
		IsFollowing: isFollow,
	}, nil
}

// ListUserPosts returns all posts from bloggers that a user is following
func (s *service) ListUserPosts(userID int) (*ct.ListPostResponse, error) {
	return nil, nil
}

// Favourite handles favorite/unfavorite operations
func (s *service) Favourite(userID, postID int, isFavourite bool) (*ct.PostFavouriteStatusResponse, error) {
	return nil, nil
}

// ListFavouritePosts returns all posts that a user has favorited
func (s *service) ListFavouritePosts(userID int) (*ct.ListPostResponse, error) {
	posts, err := s.favouriteRepo.SelectFavouritePosts(userID)
	if err != nil {
		return nil, static.ErrGetFavouritePosts
	}

	return prepareListPostResponse(posts), nil
}
