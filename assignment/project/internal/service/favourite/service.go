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

// FollowUser handles follow/unfollow operations
func (s *service) FollowUser(userID, targetUserID int, isFollow bool) (*ct.BloggerFollowStatusResponse, error) {
	// Check if user exists
	_, err := s.userRepo.Read(userID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, static.ErrUserNotFound
		}
		return nil, static.ErrReadUserID
	}

	// Check if target user exists
	_, err = s.userRepo.Read(targetUserID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, static.ErrUserNotFound
		}
		return nil, static.ErrReadUserID
	}

	// Prevent self-following
	if userID == targetUserID {
		return nil, static.ErrSelfFollow
	}

	// Check current follow status
	isFollowing, err := s.favouriteRepo.IsFollowing(userID, targetUserID)
	if err != nil {
		return nil, static.ErrReadFollowStatus
	}

	// Handle follow/unfollow based on current status
	if isFollow {
		if isFollowing {
			return nil, static.ErrAlreadyFollowing
		}
		err = s.favouriteRepo.Follow(userID, targetUserID)
	} else {
		if !isFollowing {
			return nil, static.ErrNotFollowing
		}
		err = s.favouriteRepo.Unfollow(userID, targetUserID)
	}

	if err != nil {
		return nil, err
	}

	// Get updated follow status
	isFollowing, err = s.favouriteRepo.IsFollowing(userID, targetUserID)
	if err != nil {
		return nil, static.ErrReadFollowStatus
	}

	return &ct.BloggerFollowStatusResponse{
		UserID:      targetUserID,
		IsFollowing: isFollowing,
	}, nil
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

// ListFollowerPosts returns all posts from bloggers that a user is following
func (s *service) ListFollowerPosts(userID int) (*ct.ListPostResponse, error) {
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

	userIDs := make([]int, 0, len(users))
	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}

	// If user doesn't follow anyone, return empty list
	if len(userIDs) == 0 {
		return &ct.ListPostResponse{
			Posts: []*ct.PostResponse{},
		}, nil
	}

	posts, err := s.favouriteRepo.SelectFollowedPosts(userIDs)
	if err != nil {
		return nil, err
	}

	return prepareListPostResponse(posts), nil
}

// FavouritePost handles favorite/unfavorite operations
func (s *service) FavouritePost(userID, postID int, isFavourite bool) (*ct.PostFavouriteStatusResponse, error) {
	// Check if user exists
	_, err := s.userRepo.Read(userID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, static.ErrUserNotFound
		}
		return nil, static.ErrReadUserID
	}

	// Check post existence is handled at repository level

	// Check current favourite status
	isFav, err := s.favouriteRepo.IsFavorite(userID, postID)
	if err != nil {
		return nil, err
	}

	// Handle favourite/unfavourite based on current status
	if isFavourite {
		if isFav {
			// Already favourited
			return &ct.PostFavouriteStatusResponse{
				PostID:      postID,
				IsFavourite: true,
			}, nil
		}
		err = s.favouriteRepo.AddFavorite(userID, postID)
	} else {
		if !isFav {
			// Not favourited
			return &ct.PostFavouriteStatusResponse{
				PostID:      postID,
				IsFavourite: false,
			}, nil
		}
		err = s.favouriteRepo.RemoveFavorite(userID, postID)
	}

	if err != nil {
		return nil, err
	}

	// Get updated favourite status
	isFav, err = s.favouriteRepo.IsFavorite(userID, postID)
	if err != nil {
		return nil, err
	}

	return &ct.PostFavouriteStatusResponse{
		PostID:      postID,
		IsFavourite: isFav,
	}, nil
}

// ListFavourites returns all posts that a user has favorited
func (s *service) ListFavourites(userID int) (*ct.ListPostResponse, error) {
	// Check if user exists
	_, err := s.userRepo.Read(userID)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, static.ErrUserNotFound
		}
		return nil, static.ErrReadUserID
	}

	posts, err := s.favouriteRepo.SelectFavorites(userID)
	if err != nil {
		return nil, err
	}

	return prepareListPostResponse(posts), nil
}
