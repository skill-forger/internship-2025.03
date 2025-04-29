package favourite

import (
	ct "golang-project/internal/contract"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
	"golang-project/static"
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
		return nil, static.ErrUserNotFound
	}

	users, err := s.favouriteRepo.SelectFollowing(userID)
	if err != nil {
		return nil, err
	}

	return prepareListProfileResponse(users), nil
}

// Follow handles follow/unfollow operations
func (s *service) Follow(userID, targetUserID int, isFollow bool) (*ct.BloggerFollowStatusResponse, error) {
	return nil, nil
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
