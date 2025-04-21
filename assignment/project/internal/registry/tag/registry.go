package tag

import (
	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/tag"
	repo "golang-project/internal/repository/tag"
	svc "golang-project/internal/service/tag"
	"gorm.io/gorm"
)

// NewRegistry returns new resource handler for profile API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, svc.NewService(repo.NewRepository(db)))
}
