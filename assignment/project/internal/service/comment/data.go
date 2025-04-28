package comment

import (
	ct "golang-project/internal/contract"
	"golang-project/internal/model"
	"time"
)

// prepareCommentResponse converts a model.Comment to ct.CommentResponse
func prepareCommentResponse(comment *model.Comment) *ct.CommentResponse {
	if comment == nil {
		return nil
	}

	commentResp := &ct.CommentResponse{
		ID:              comment.ID,
		Content:         comment.Content,
		ParentCommentID: comment.ParentCommentID,
		CreatedAt:       comment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       comment.UpdatedAt.Format(time.RFC3339),
		ChildComments:   make([]*ct.ChildCommentResponse, 0),
	}

	// Add user info if available
	if comment.User != nil {
		commentResp.User = prepareProfileResponse(comment.User)
	}

	// Add post info if available
	if comment.Post != nil {
		commentResp.Post = preparePostResponse(comment.Post)
	}

	// Add child comments if available
	if comment.ChildComments != nil && len(comment.ChildComments) > 0 {
		commentResp.ChildComments = make([]*ct.ChildCommentResponse, 0, len(comment.ChildComments))
		for _, child := range comment.ChildComments {
			
			childResp := prepareChildCommentResponse(child)
			if childResp != nil {
				commentResp.ChildComments = append(commentResp.ChildComments, childResp)
			}
		}
	}

	return commentResp
}

// prepareChildCommentResponse converts a model.Comment to ct.ChildCommentResponse
func prepareChildCommentResponse(comment *model.Comment) *ct.ChildCommentResponse {
	if comment == nil {
		return nil
	}

	childResp := &ct.ChildCommentResponse{
		ID:              comment.ID,
		Content:         comment.Content,
		ParentCommentID: comment.ParentCommentID,
		CreatedAt:       comment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       comment.UpdatedAt.Format(time.RFC3339),
	}

	if comment.User != nil {
		childResp.User = prepareProfileResponse(comment.User)
	}

	return childResp
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

	if post.Tags != nil && len(post.Tags) > 0 {
		postResp.Tags = make([]*ct.TagDetailResponse, 0, len(post.Tags))
		for _, tag := range post.Tags {
			tagResp := prepareTagResponse(tag)
			if tagResp != nil {
				postResp.Tags = append(postResp.Tags, tagResp)
			}
		}
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

// prepareTagResponse converts a model.Tag to ct.TagDetailResponse
func prepareTagResponse(tag *model.Tag) *ct.TagDetailResponse {
	if tag == nil {
		return nil
	}

	return &ct.TagDetailResponse{
		ID:        tag.ID,
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt.Format(time.RFC3339),
		UpdatedAt: tag.UpdatedAt.Format(time.RFC3339),
	}
}

// prepareListCommentResponse converts a slice of model.Comment to ct.ListCommentResponse
func prepareListCommentResponse(comments []*model.Comment, paging ct.Paging) *ct.ListCommentResponse {
	if comments == nil {
		comments = make([]*model.Comment, 0)
	}

	response := &ct.ListCommentResponse{
		Comments: make([]*ct.CommentResponse, 0, len(comments)),
		Paging:   paging,
	}

	for _, comment := range comments {
		commentResp := prepareCommentResponse(comment)
		if commentResp != nil {
			response.Comments = append(response.Comments, commentResp)
		}
	}

	return response
}
