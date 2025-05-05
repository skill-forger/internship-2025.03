package profile

import (
	ct "golang-project/internal/contract"
	repo "golang-project/internal/repository"
	svc "golang-project/internal/service"
)

// service represents the implementation of service.Profile
type service struct {
	userRepo repo.User
}

// NewService returns a new implementation of service.Profile
func NewService(userRepo repo.User) svc.Profile {
	return &service{
		userRepo: userRepo,
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

	// Update user fields
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Pseudonym != "" {
		user.Pseudonym = req.Pseudonym
	}
	if req.ProfileImage != "" {
		user.ProfileImage = req.ProfileImage
	}
	if req.Biography != "" {
		user.Biography = req.Biography
	}

	// Save updated user
	updatedUser, err := s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return prepareProfileResponse(updatedUser), nil
}
