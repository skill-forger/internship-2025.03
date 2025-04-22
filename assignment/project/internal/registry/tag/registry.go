package tag

import (
	"gorm.io/gorm"

	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/tag"
	repo "golang-project/internal/repository/tag"
	svc "golang-project/internal/service/tag"
)

// NewRegistry returns a new resource handler for tag API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, svc.NewService(repo.NewRepository(db)))
}
