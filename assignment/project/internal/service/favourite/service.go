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

// ListFollowers returns a list of all bloggers that a user is following
func (s *service) ListFollowers(userID int) (*ct.ListProfileResponse, error) {
	// Check if user exists
	_, err := s.userRepo.Read(userID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, static.ErrUserNotFound
		}
		return nil, static.ErrReadUserID
	}

	users, err := s.favouriteRepo.SelectFollowed(userID)
	if err != nil {
		return nil, err
	}

	return prepareListProfileResponse(users), nil
}

// FollowUser handles follow/unfollow operations
func (s *service) FollowUser(userID, targetUserID int, isFollow bool) (*ct.BloggerFollowStatusResponse, error) {
	return nil, nil
}

// ListFollowerPosts returns all posts from bloggers that a user is following
func (s *service) ListFollowerPosts(userID int) (*ct.ListPostResponse, error) {
	return nil, nil
}

// FavouritePost handles favorite/unfavorite operations
func (s *service) FavouritePost(userID, postID int, isFavourite bool) (*ct.PostFavouriteStatusResponse, error) {
	return nil, nil
}

// ListFavourites returns all posts that a user has favorited
func (s *service) ListFavourites(userID int) (*ct.ListPostResponse, error) {
	return nil, nil
}
