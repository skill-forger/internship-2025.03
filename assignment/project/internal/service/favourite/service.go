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

	// Handle different actions
	switch req.Action {
	case static.Follow:
		if isFollowing {
			return &ct.BloggerFollowStatusResponse{
				UserID:      targetUserID,
				IsFollowing: true,
			}, nil
		}

		// Add to following list if not already there
		follow := &model.FollowUser{
			UserID:       userID,
			FollowUserID: targetUserID,
		}
		if err = s.favouriteRepo.Follow(follow); err != nil {
			return nil, static.ErrFollowStatusUpdate
		}

		return &ct.BloggerFollowStatusResponse{
			UserID:      targetUserID,
			IsFollowing: true,
		}, nil

	case static.Unfollow:
		if !isFollowing {
			return &ct.BloggerFollowStatusResponse{
				UserID:      targetUserID,
				IsFollowing: false,
			}, nil
		}

		// Remove from following list if it's there
		if err = s.favouriteRepo.Unfollow(userID, targetUserID); err != nil {
			return nil, static.ErrFollowStatusUpdate
		}

		return &ct.BloggerFollowStatusResponse{
			UserID:      targetUserID,
			IsFollowing: false,
		}, nil

	default:
		return nil, static.ErrUnsupportedFollowAction
	}
}

// ListUserPosts returns all posts from bloggers that a user is following
func (s *service) ListUserPosts(userID int) (*ct.ListPostResponse, error) {
	posts, err := s.favouriteRepo.SelectFollowingUsersPosts(userID)
	if err != nil {
		return nil, static.ErrGetFollowedBloggerPosts
	}

	return prepareListPostResponse(posts), nil
}

// UpdateFavouriteStatus handles favorite/unfavorite operations
func (s *service) UpdateFavouriteStatus(userID int, req *ct.PostFavouriteRequest) (*ct.PostFavouriteStatusResponse, error) {
	postID := req.PostID

	// Check if post exists
	_, err := s.postRepo.Read(postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, static.ErrPostNotFound
		}
		return nil, static.ErrDatabaseOperation
	}

	// Check current favourite status
	isAlreadyFavourited, err := s.favouriteRepo.IsFavourite(userID, postID)
	if err != nil {
		return nil, static.ErrDatabaseOperation
	}

	// Handle different actions
	switch req.Action {
	case static.Favourite:
		if isAlreadyFavourited {
			return &ct.PostFavouriteStatusResponse{
				PostID:      postID,
				IsFavourite: true,
			}, nil
		}

		// Add to favourites if not already there
		favourite := &model.FavoritePost{
			UserID: userID,
			PostID: postID,
		}
		if err = s.favouriteRepo.Favourite(favourite); err != nil {
			return nil, static.ErrFavouriteStatusUpdate
		}

		return &ct.PostFavouriteStatusResponse{
			PostID:      postID,
			IsFavourite: true,
		}, nil

	case static.Unfavourite:
		if !isAlreadyFavourited {
			return &ct.PostFavouriteStatusResponse{
				PostID:      postID,
				IsFavourite: false,
			}, nil
		}

		// Remove from favourites if it's there
		if err = s.favouriteRepo.Unfavourite(userID, postID); err != nil {
			return nil, static.ErrFavouriteStatusUpdate
		}

		return &ct.PostFavouriteStatusResponse{
			PostID:      postID,
			IsFavourite: false,
		}, nil

	default:
		return nil, static.ErrUnsupportedFavouriteAction
	}
}

// ListFavouritePosts returns all posts that a user has favorited
func (s *service) ListFavouritePosts(userID int) (*ct.ListPostResponse, error) {
	posts, err := s.favouriteRepo.SelectFavouritePosts(userID)
	if err != nil {
		return nil, static.ErrGetFavouritePosts
	}

	return prepareListPostResponse(posts), nil
}
