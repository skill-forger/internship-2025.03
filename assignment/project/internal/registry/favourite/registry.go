package favourite

import (
	"gorm.io/gorm"

	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/favourite"
	repo "golang-project/internal/repository/favourite"
	postRepo "golang-project/internal/repository/post"
	userRepo "golang-project/internal/repository/user"
	svc "golang-project/internal/service/favourite"
)

// NewRegistry returns new resource handler for favourite API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	favouriteRepo := repo.NewRepository(db)
	userRepoInst := userRepo.NewRepository(db)
	postRepoInst := postRepo.NewRepository(db)
	favouriteSvc := svc.NewService(favouriteRepo, userRepoInst, postRepoInst)
	return hdl.NewHandler(route, favouriteSvc)
}
