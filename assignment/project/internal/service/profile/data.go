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
func updateProfileFields(o *model.User, req *ct.UpdateProfileRequest) {
	if req.FirstName != "" {
		o.FirstName = req.FirstName
	}
	if req.LastName != "" {
		o.LastName = req.LastName
	}
	if req.Pseudonym != "" {
		o.Pseudonym = req.Pseudonym
	}

	// ProfileImage & Biography can be empty strings
	o.ProfileImage = req.ProfileImage
	o.Biography = req.Biography
}

// prepareUpdatesMap prepares the map of fields to update
func prepareUpdatesMap(o *model.User) map[string]interface{} {
	return map[string]interface{}{
		"first_name":    o.FirstName,
		"last_name":     o.LastName,
		"pseudonym":     o.Pseudonym,
		"profile_image": o.ProfileImage,
		"biography":     o.Biography,
	}
}
