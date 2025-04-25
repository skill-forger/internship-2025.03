package comment

import (
	"gorm.io/gorm"

	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/comment"
	repo "golang-project/internal/repository/comment"
	svc "golang-project/internal/service/comment"
)

// NewRegistry returns new resource handler for profile API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, svc.NewService(repo.NewRepository(db)))
}
