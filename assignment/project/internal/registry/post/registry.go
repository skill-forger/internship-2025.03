package post

import (
	"gorm.io/gorm"

	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/post"

	// repoComment "golang-project/internal/repository/Comment"
	repoPost "golang-project/internal/repository/post"
	repoPT "golang-project/internal/repository/posttag"
	repoTag "golang-project/internal/repository/tag"
	repoUser "golang-project/internal/repository/user"
	svc "golang-project/internal/service/post"
)

// NewRegistry returns a new resource handler for tag API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, svc.NewService(
		repoPost.NewRepository(db),
		repoUser.NewRepository(db),
		repoTag.NewRepository(db),
		repoPT.NewRepository(db),
		// repoComment.NewRepository(db),
	))
}
