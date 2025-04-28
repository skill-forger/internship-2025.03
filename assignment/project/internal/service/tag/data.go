package tag

import (
	"time"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
)

type postData struct {
	postdata *model.Post
	userdata *model.User
	tagsdata []*model.Tag
}

// prepareTagResponse transforms tag model to response DTO
func prepareTagResponse(o *model.Tag) *ct.TagDetailResponse {
	data := &ct.TagDetailResponse{
		ID:   o.ID,
		Name: o.Name,
	}

	if o.CreatedAt != nil {
		data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
	}

	if o.UpdatedAt != nil {
		data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
	}

	return data
}

// prepareListTagResponse transforms the data and returns the List Tag Response
func prepareListTagResponse(o []*model.Tag) *ct.ListTagResponse {
	data := &ct.ListTagResponse{
		Tags: make([]*ct.TagDetailResponse, 0, len(o)),
	}

	for _, tag := range o {
		data.Tags = append(data.Tags, prepareTagResponse(tag))
	}

	return data
}

// preparePostResponse transforms post model to response DTO
func preparePostResponse(o *model.Post, u *model.User, t []*model.Tag) *ct.PostResponse {
	data := &ct.PostResponse{
		ID:          o.ID,
		Title:       o.Title,
		Body:        o.Body,
		Slug:        o.Slug,
		IsPublished: o.IsPublished,
		User:        prepareProfileResponse(u),
		Tags:        prepareListTagResponse(t),
	}

	if o.CreatedAt != nil {
		data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
	}

	if o.UpdatedAt != nil {
		data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
	}

	return data
}

// prepareListPostResponse transforms the data and returns the List Post Response
func prepareListPostResponse(p []*model.Post, pt []*model.PostTag, t []*model.Tag, u []*model.User) *ct.ListPostResponse {
	data := &ct.ListPostResponse{
		Posts: make([]*ct.PostResponse, 0, len(p)),
	}

	postMap := make(map[int]*postData, len(p))
	tagMap := make(map[int]*model.Tag, len(t))

	// Build Tag Map
	for _, tag := range t {
		tagMap[tag.ID] = tag
	}

	// Build Post Map
	for _, post := range p {
		postMap[post.ID] = &postData{
			postdata: post,
		}

		for _, user := range u {
			if post.UserID == user.ID {
				postMap[post.ID].userdata = user
				break
			}
		}
	}

	for _, ptItem := range pt {
		if postData, ok := postMap[ptItem.PostID]; ok {
			if tag, exists := tagMap[ptItem.TagID]; exists {
				postData.tagsdata = append(postData.tagsdata, tag)
			}
		}
	}

	for _, postData := range postMap {
		data.Posts = append(data.Posts, preparePostResponse(postData.postdata, postData.userdata, postData.tagsdata))
	}

	return data
}

// prepareProfileResponse transforms user model to response DTO
func prepareProfileResponse(o *model.User) *ct.ProfileResponse {
	data := &ct.ProfileResponse{
		ID:           o.ID,
		FirstName:    o.FirstName,
		LastName:     o.LastName,
		Email:        o.Email,
		Pseudonym:    o.Pseudonym,
		Biography:    o.Biography,
		ProfileImage: o.ProfileImage,
	}

	if o.CreatedAt != nil {
		data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
	}

	if o.UpdatedAt != nil {
		data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
	}

	return data
}
