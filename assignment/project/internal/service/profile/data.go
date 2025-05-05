package profile

import (
	"time"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
)

// prepareSignInResponse transforms the data and returns the Profile Response
func prepareProfileResponse(o *model.User) *ct.ProfileResponse {
	data := &ct.ProfileResponse{
		ID:           o.ID,
		FirstName:    o.FirstName,
		LastName:     o.LastName,
		Email:        o.Email,
		Pseudonym:    o.Pseudonym,
		ProfileImage: o.ProfileImage,
		Biography:    o.Biography,
	}

	if o.CreatedAt != nil {
		data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
	}

	if o.UpdatedAt != nil {
		data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
	}

	return data
}

// updateProfileFields updates fields of a User model using data from UpdateProfileRequest
func updateProfileFields(user *model.User, req *ct.UpdateProfileRequest) {
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
}
