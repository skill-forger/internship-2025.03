package comment

import (
	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/comment"
	repo "golang-project/internal/repository/comment"
	svc "golang-project/internal/service/comment"
	"gorm.io/gorm"
)

// NewRegistry returns new resource handler for profile API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, svc.NewService(repo.NewRepository(db)))
}
