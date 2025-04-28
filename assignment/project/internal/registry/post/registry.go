package post

import (
	"gorm.io/gorm"

	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/post"
	repoPost "golang-project/internal/repository/post"
	repoTag "golang-project/internal/repository/tag"
	repoUser "golang-project/internal/repository/user"
	svc "golang-project/internal/service/post"
)

// NewRegistry returns new resource handler for post API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	postRepo := repoPost.NewRepository(db)
	userRepo := repoUser.NewRepository(db)
	tagRepo := repoTag.NewRepository(db)

	return hdl.NewHandler(route, svc.NewService(postRepo, userRepo, tagRepo))
}
