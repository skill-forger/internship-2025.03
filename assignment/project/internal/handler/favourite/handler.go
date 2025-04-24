package favourite

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"golang-project/internal/contract"
	hdl "golang-project/internal/handler"
	"golang-project/server"
)

// handler represents the implementation of handler.Favourite
type handler struct {
	route string
}

// NewHandler returns a new implementation of handler.Favourite
func NewHandler(route string) hdl.Favourite {
	return &handler{
		route: route,
	}
}

// RegisterRoutes registers the handler routes and returns the server.HandlerRegistry
func (h *handler) RegisterRoutes() server.HandlerRegistry {
	return server.HandlerRegistry{
		Route:           h.route,
		IsAuthenticated: true,
		Register: func(group *echo.Group) {
			group.PUT("/bloggers", h.UpdateBlogger)
			group.GET("/bloggers", h.ListBloggers)
			group.GET("/bloggers/posts", h.ListBloggerPosts)
			group.PUT("/posts", h.UpdatePost)
			group.GET("/posts", h.ListPosts)
		},
	}
}

// UpdateBlogger handles adding/removing blogger from following list
//
//	@Summary		Add/remove blogger from following list
//	@Description	Blogger can add/remove blogger from their following list
//	@Tags			favourite
//	@Accept			json
//	@Produce		json
//	@Security		BearerToken
//	@Param			request	body		contract.BloggerFollowRequest	true	"Follow (1) or unfollow (0) a blogger"
//	@Success		200		{object}	contract.BloggerFollowStatusResponse
//	@Failure		400		{object}	error
//	@Router			/favorites/bloggers [put]
//	@enum			Action	0:Unfollow,1:Follow
func (h *handler) UpdateBlogger(e echo.Context) error {
	var req contract.BloggerFollowRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	// Placeholder implementation
	isFollowing := req.Action == contract.Follow
	return e.JSON(http.StatusOK, &contract.BloggerFollowStatusResponse{
		UserID:      req.UserID,
		IsFollowing: isFollowing,
	})
}

// ListBloggers handles the request to get all bloggers from following list
//
//	@Summary		View all followed bloggers
//	@Description	Blogger can view all the bloggers from their following list
//	@Tags			favourite
//	@Accept			json
//	@Produce		json
//	@Security		BearerToken
//	@Success		200	{object}	contract.ListProfileResponse
//	@Failure		400	{object}	error
//	@Router			/favorites/bloggers [get]
func (h *handler) ListBloggers(e echo.Context) error {
	// Placeholder implementation
	return e.JSON(http.StatusOK, &contract.ListProfileResponse{
		Bloggers: []*contract.ProfileResponse{},
	})
}

// ListBloggerPosts handles the request to get all posts from followed bloggers
//
//	@Summary		View posts from followed bloggers
//	@Description	Blogger can view all the posts of the following bloggers
//	@Tags			favourite
//	@Accept			json
//	@Produce		json
//	@Security		BearerToken
//	@Success		200	{object}	contract.ListPostResponse
//	@Failure		400	{object}	error
//	@Router			/favorites/bloggers/posts [get]
func (h *handler) ListBloggerPosts(e echo.Context) error {
	// Placeholder implementation
	return e.JSON(http.StatusOK, &contract.ListPostResponse{
		Posts: []*contract.PostResponse{},
	})
}

// UpdatePost handles adding/removing a post from favourite list
//
//	@Summary		Add/remove post from favourite list
//	@Description	Blogger can add/remove a post from their favourite list
//	@Tags			favourite
//	@Accept			json
//	@Produce		json
//	@Security		BearerToken
//	@Param			request	body		contract.PostFavouriteRequest	true	"Add to favourite (1) or remove from favourites (0)"
//	@Success		200		{object}	contract.PostFavouriteStatusResponse
//	@Failure		400		{object}	error
//	@Router			/favorites/posts [put]
//	@enum			Action	0:Unfavourite,1:Favourite
func (h *handler) UpdatePost(e echo.Context) error {
	var req contract.PostFavouriteRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	// Placeholder implementation
	isFavourite := req.Action == contract.Favourite
	return e.JSON(http.StatusOK, &contract.PostFavouriteStatusResponse{
		PostID:      req.PostID,
		IsFavourite: isFavourite,
	})
}

// ListPosts handles the request to get all posts from favourite list
//
//	@Summary		View favourite posts
//	@Description	Blogger can view all posts from their favourite list
//	@Tags			favourite
//	@Accept			json
//	@Produce		json
//	@Security		BearerToken
//	@Success		200	{object}	contract.ListPostResponse
//	@Failure		400	{object}	error
//	@Router			/favorites/posts [get]
func (h *handler) ListPosts(e echo.Context) error {
	// Placeholder implementation
	return e.JSON(http.StatusOK, &contract.ListPostResponse{
		Posts: []*contract.PostResponse{},
	})
}
