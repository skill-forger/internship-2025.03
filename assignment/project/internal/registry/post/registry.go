package post

import (
	"gorm.io/gorm"

	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/post"
	postRepo "golang-project/internal/repository/post"
	tagRepo "golang-project/internal/repository/tag"
	userRepo "golang-project/internal/repository/user"
	svc "golang-project/internal/service/post"
)

// NewRegistry returns new resource handler for post API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	postRepo := postRepo.NewRepository(db)
	userRepo := userRepo.NewRepository(db)
	tagRepo := tagRepo.NewRepository(db)

	return hdl.NewHandler(route, svc.NewService(postRepo, userRepo, tagRepo))
}
