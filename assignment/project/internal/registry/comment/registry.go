package comment

import (
	"gorm.io/gorm"

	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/comment"
	commentRepo "golang-project/internal/repository/comment"
	postRepo "golang-project/internal/repository/post"
	userRepo "golang-project/internal/repository/user"
	svc "golang-project/internal/service/comment"
)

// NewRegistry returns new resource handler for profile API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(
		route,
		svc.NewService(
			commentRepo.NewRepository(db),
			userRepo.NewRepository(db),
			postRepo.NewRepository(db),
		),
	)
}
