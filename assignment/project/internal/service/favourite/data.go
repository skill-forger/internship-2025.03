package favourite

import (
	"time"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
)

// prepareProfileResponse converts a User model to a ProfileResponse
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


// prepareListProfileResponse converts slice of User models to ListProfileResponse
func prepareListProfileResponse(o []*model.User) *ct.ListProfileResponse {
	data := &ct.ListProfileResponse{
		Bloggers: make([]*ct.ProfileResponse, 0, len(o)),
	}

	for _, user := range o {
		data.Bloggers = append(data.Bloggers, prepareProfileResponse(user))
	}

	return data
}
