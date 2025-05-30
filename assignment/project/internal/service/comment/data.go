package comment

import (
	"time"

	ct "golang-project/internal/contract"
	"golang-project/internal/model"
)

// prepareCommentModel creates a new comment model from request data
func prepareCommentModel(request *ct.CreateCommentRequest, userID int, post *model.Post) *model.Comment {
	return &model.Comment{
		Content:         request.Content,
		PostID:          request.PostID,
		UserID:          userID,
		ParentCommentID: request.ParentCommentID,
		Post:            post,
	}
}

// prepareUpdateComment updates fields of a Comment model and prepares the update map
func prepareUpdateComment(o *model.Comment, req *ct.UpdateCommentRequest) map[string]interface{} {
	if req.Content != "" {
		o.Content = req.Content
	}

	return map[string]interface{}{
		"content": o.Content,
	}
}

// prepareListTagResponse transforms the data and returns the List Tag Response
func prepareListTagResponse(o []*model.Tag) []*ct.TagResponse {
	data := make([]*ct.TagResponse, 0, len(o))

	for _, tag := range o {
		data = append(data, prepareTagResponse(tag))
	}

	return data
}

// preparePostResponse converts a model.Post to ct.PostResponse
func preparePostResponse(post *model.Post) *ct.PostResponse {
	postResp := &ct.PostResponse{
		ID:          post.ID,
		Title:       post.Title,
		Body:        post.Body,
		Slug:        post.Slug,
		IsPublished: post.IsPublished,
		CreatedAt:   post.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   post.UpdatedAt.Format(time.RFC3339),
	}

	if post.User != nil {
		postResp.User = prepareProfileResponse(post.User)
	}

	if post.Tags != nil {
		postResp.Tags = prepareListTagResponse(post.Tags)
	}

	return postResp
}

// prepareProfileResponse converts a model.User to ct.ProfileResponse
func prepareProfileResponse(user *model.User) *ct.ProfileResponse {
	return &ct.ProfileResponse{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Pseudonym:    user.Pseudonym,
		ProfileImage: user.ProfileImage,
		Biography:    user.Biography,
		CreatedAt:    user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    user.UpdatedAt.Format(time.RFC3339),
	}
}

// prepareTagResponse converts a model.Tag to ct.TagResponse
func prepareTagResponse(tag *model.Tag) *ct.TagResponse {
	return &ct.TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

// prepareChildCommentResponse converts a model.Comment to ct.ChildCommentResponse
func prepareChildCommentResponse(o *model.Comment, u *model.User) *ct.ChildCommentResponse {
	data := &ct.ChildCommentResponse{
		ID:              o.ID,
		Content:         o.Content,
		ParentCommentID: o.ParentCommentID,
		User:            prepareProfileResponse(u),
	}

	if o.CreatedAt != nil {
		data.CreatedAt = o.CreatedAt.Format(time.RFC3339)
	}

	if o.UpdatedAt != nil {
		data.UpdatedAt = o.UpdatedAt.Format(time.RFC3339)
	}

	return data
}

// prepareCommentResponse converts a model.Comment to ct.CommentResponse
func prepareCommentResponse(comment *model.Comment) *ct.CommentResponse {
	data := &ct.CommentResponse{
		ID:      comment.ID,
		Content: comment.Content,
	}

	if comment.ParentCommentID != nil {
		data.ParentCommentID = comment.ParentCommentID
	}

	if comment.CreatedAt != nil {
		data.CreatedAt = comment.CreatedAt.Format(time.RFC3339)
	}

	if comment.UpdatedAt != nil {
		data.UpdatedAt = comment.UpdatedAt.Format(time.RFC3339)
	}

	// Add user info if available
	if comment.User != nil {
		data.User = prepareProfileResponse(comment.User)
	}

	// Add post info if available
	if comment.Post != nil {
		data.Post = preparePostResponse(comment.Post)
	}

	// Add child comments if available
	data.ChildComments = make([]*ct.ChildCommentResponse, 0, len(comment.ChildComments))
	for _, child := range comment.ChildComments {
		childResp := prepareChildCommentResponse(child, child.User)
		if childResp != nil {
			data.ChildComments = append(data.ChildComments, childResp)
		}
	}

	return data
}

// prepareListCommentResponse converts a slice of model.Comment to ct.ListCommentResponse
func prepareListCommentResponse(comments []*model.Comment, paging ct.Paging) *ct.ListCommentResponse {
	response := &ct.ListCommentResponse{
		Comments: make([]*ct.CommentResponse, 0, len(comments)),
		Paging:   paging,
	}

	for _, comment := range comments {
		response.Comments = append(response.Comments, prepareCommentResponse(comment))
	}

	return response
}
