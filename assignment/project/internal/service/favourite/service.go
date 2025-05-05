package favourite

import (
	"errors"
	ct "golang-project/internal/contract"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// service represents the implementation of service.Favourite
type service struct {
	favouriteRepo repo.Favourite
	userRepo      repo.User
}

// NewService returns a new implementation of service.Favourite
func NewService(favouriteRepo repo.Favourite, userRepo repo.User) svc.Favourite {
	return &service{
		favouriteRepo: favouriteRepo,
		userRepo:      userRepo,
	}
}

// ListFollowingUsers returns a list of all bloggers that a user is following
func (s *service) ListFollowingUsers(userID int) (*ct.ListProfileResponse, error) {
	// Check if user exists
	_, err := s.userRepo.Read(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}

	users, err := s.favouriteRepo.SelectFollowing(userID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve followed users")
	}

	return prepareListProfileResponse(users), nil
}

// UpdateFollowStatus handles follow/unfollow operations
func (s *service) UpdateFollowStatus(userID, targetUserID int, isFollow bool) (*ct.BloggerFollowStatusResponse, error) {
	// Check if target user exists
	_, err := s.userRepo.Read(targetUserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Database error")
	}
	// Prevent self-following
	if userID == targetUserID {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Cannot follow yourself")
	}
	// Check current follow status
	isFollowing, err := s.favouriteRepo.IsFollowing(userID, targetUserID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to check follow status")
	}
	// Handle follow/unfollow based on current status
	if isFollow {
		if isFollowing {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "Already following this user")
		}
		err = s.favouriteRepo.Follow(userID, targetUserID)
	} else {
		if !isFollowing {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "Not following this user")
		}
		err = s.favouriteRepo.Unfollow(userID, targetUserID)
	}
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to update follow status")
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
	return nil, nil
}
