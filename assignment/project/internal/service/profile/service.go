package profile

import (
	"strconv"

	ct "golang-project/internal/contract"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
	"golang-project/static"
	"golang-project/util/hashing"
)

// service represents the implementation of service.Profile
type service struct {
	userRepo repo.User
	postRepo repo.Post
	tagRepo  repo.Tag
}

// NewService returns a new implementation of service.Profile
func NewService(userRepo repo.User, postRepo repo.Post, tagRepo repo.Tag) svc.Profile {
	return &service{
		userRepo: userRepo,
		postRepo: postRepo,
		tagRepo:  tagRepo,
	}
}

// GetByID executes the profile detail retrieval logic
func (s *service) GetByID(id int) (*ct.ProfileResponse, error) {
	user, err := s.userRepo.Read(id)
	if err != nil {
		return nil, err
	}

	return prepareProfileResponse(user), nil
}

// Update executes the profile update logic
func (s *service) Update(id int, req *ct.UpdateProfileRequest) (*ct.ProfileResponse, error) {
	user, err := s.userRepo.Read(id)
	if err != nil {
		return nil, err
	}

	// Prepare updates user fields
	updates := prepareUpdateProfile(user, req)

	// Save updated user
	updatedUser, err := s.userRepo.Update(user, updates)
	if err != nil {
		return nil, err
	}

	return prepareProfileResponse(updatedUser), nil
}

// ChangePassword executes the password change logic
func (s *service) ChangePassword(id int, req *ct.ChangePasswordRequest) (*ct.ChangePasswordResponse, error) {
	user, err := s.userRepo.Read(id)
	if err != nil {
		return nil, err
	}

	// Verify current password
	hash := hashing.NewBcrypt()
	err = hash.Compare([]byte(user.Password), []byte(req.CurrentPassword))
	if err != nil {
		return nil, static.ErrInvalidPassword
	}

	// Verify new password matches confirm password
	if req.NewPassword != req.ConfirmNewPassword {
		return nil, static.ErrComfirmPassword
	}

	// Hash new password
	hashedPassword, err := hash.Generate([]byte(req.NewPassword))
	if err != nil {
		return nil, static.ErrPasswordHashingFailed
	}

	// Update password
	updates := map[string]any{
		"password": hashedPassword,
	}

	// Save updated password
	_, err = s.userRepo.Update(user, updates)
	if err != nil {
		return nil, err
	}

	return &ct.ChangePasswordResponse{
		Message: "Password changed successfully",
	}, nil
}

// GetPost executes the User get their own post detail retrieval logic
func (s *service) GetPost(postID, ctxUserID int) (*ct.PostResponse, error) {
	post, err := s.postRepo.ReadByCondition(map[string]any{
		"id": postID,
	}, "User", "Tags")
	if err != nil {
		return nil, err
	}

	// Check ctxUser own the post
	if ctxUserID != post.UserID {
		return nil, static.ErrPostOwner
	}

	return preparePostResponse(post), nil
}

// ListBloggerPosts executes the User get their own posts retrieval logic
func (s *service) ListBloggerPosts(id int, isPublishedParam string) (*ct.ListPostResponse, error) {
	var isPublishedFilter *bool
	if isPublishedParam != "" {
		if b, err := strconv.ParseBool(isPublishedParam); err == nil {
			isPublishedFilter = &b
		} else {
			return nil, static.ErrParamInvalid
		}
	}

	posts, err := s.userRepo.ReadOwnPosts(id, isPublishedFilter)
	if err != nil {
		return nil, static.ErrListBloggerPosts
	}

	responses := make([]*ct.PostResponse, 0, len(posts))
	for _, post := range posts {
		responses = append(responses, preparePostResponse(post))
	}

	return &ct.ListPostResponse{
		Posts: responses,
	}, nil
}
