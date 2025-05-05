package post

import (
	"gorm.io/gorm"

	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/post"

	// repoComment "golang-project/internal/repository/Comment"
	postRepo "golang-project/internal/repository/post"
	ptRepo "golang-project/internal/repository/posttag"
	tagRepo "golang-project/internal/repository/tag"
	userRepo "golang-project/internal/repository/user"
	svc "golang-project/internal/service/post"
)

// NewRegistry returns a new resource handler for tag API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, svc.NewService(
		postRepo.NewRepository(db),
		userRepo.NewRepository(db),
		tagRepo.NewRepository(db),
		ptRepo.NewRepository(db),
		// repoComment.NewRepository(db),
	))
}
