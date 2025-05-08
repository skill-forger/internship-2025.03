package favourite

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"golang-project/internal/contract"
	hdl "golang-project/internal/handler"
	svc "golang-project/internal/service"
	"golang-project/server"
	"golang-project/static"
)

// handler represents the implementation of handler.Favourite
type handler struct {
	route        string
	favouriteSvc svc.Favourite
}

// NewHandler returns a new implementation of handler.Favourite
func NewHandler(route string, favouriteSvc svc.Favourite) hdl.Favourite {
	return &handler{
		route:        route,
		favouriteSvc: favouriteSvc,
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
//	@Param			request	body		contract.BloggerFollowRequest	true	"Follow/unfollow action with user ID"
//	@Success		200		{object}	contract.BloggerFollowStatusResponse
//	@Failure		400		{object}	error
//	@Router			/favorites/bloggers [put]
func (h *handler) UpdateBlogger(e echo.Context) error {
	var req contract.BloggerFollowRequest
	if err := e.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	ctxUser, err := hdl.GetContextUser(e)
	if err != nil {
		return err
	}

	resp, err := h.favouriteSvc.UpdateFollowStatus(ctxUser.ID, &req)
	if err != nil {
		switch err {
		case static.ErrUserNotFound:
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		case static.ErrSelfFollow:
			return echo.NewHTTPError(http.StatusBadRequest, "Cannot follow yourself")
		case static.ErrUnsupportedFollowAction:
			return echo.NewHTTPError(http.StatusBadRequest, "Unsupported follow action")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update follow status")
		}
	}

	return e.JSON(http.StatusOK, resp)
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
	ctxUser, err := hdl.GetContextUser(e)
	if err != nil {
		return err
	}

	resp, err := h.favouriteSvc.ListFollowingUsers(ctxUser.ID)
	if err != nil {
		if err == static.ErrUserNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "User not found")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve followed bloggers")
	}

	return e.JSON(http.StatusOK, resp)
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
	ctxUser, err := hdl.GetContextUser(e)
	if err != nil {
		return err
	}

	resp, err := h.favouriteSvc.ListUserPosts(ctxUser.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve posts from followed bloggers")
	}

	return e.JSON(http.StatusOK, resp)
}

// UpdatePost handles adding/removing a post from favourite list
//
//	@Summary		Add/remove post from favourite list
//	@Description	Blogger can add/remove a post from their favourite list
//	@Tags			favourite
//	@Accept			json
//	@Produce		json
//	@Security		BearerToken
//	@Param			request	body		contract.PostFavouriteRequest	true	"Add/remove post from favourites action with post ID"
//	@Success		200		{object}	contract.PostFavouriteStatusResponse
//	@Failure		400		{object}	error
//	@Router			/favorites/posts [put]
func (h *handler) UpdatePost(e echo.Context) error {
	var req contract.PostFavouriteRequest
	if err := e.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	ctxUser, err := hdl.GetContextUser(e)
	if err != nil {
		return err
	}

	resp, err := h.favouriteSvc.UpdateFavouriteStatus(ctxUser.ID, &req)
	if err != nil {
		switch err {
		case static.ErrPostNotFound:
			return echo.NewHTTPError(http.StatusNotFound, "Post not found")
		case static.ErrUnsupportedFavouriteAction:
			return echo.NewHTTPError(http.StatusBadRequest, "Unsupported favourite action")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update favourite status")
		}
	}

	return e.JSON(http.StatusOK, resp)
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
	ctxUser, err := hdl.GetContextUser(e)
	if err != nil {
		return err
	}

	resp, err := h.favouriteSvc.ListFavouritePosts(ctxUser.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve favourite posts")
	}

	return e.JSON(http.StatusOK, resp)
}
