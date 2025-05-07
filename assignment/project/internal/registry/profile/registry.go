package profile

import (
	"gorm.io/gorm"

	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/profile"
	postRepo "golang-project/internal/repository/post"
	tagRepo "golang-project/internal/repository/tag"
	userRepo "golang-project/internal/repository/user"
	svc "golang-project/internal/service/profile"
)

// NewRegistry returns new resource handler for profile API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, svc.NewService(
		userRepo.NewRepository(db),
		postRepo.NewRepository(db),
		tagRepo.NewRepository(db),
	))
}
