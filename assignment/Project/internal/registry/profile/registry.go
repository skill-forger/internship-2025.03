package profile

import (
	"gorm.io/gorm"

	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/profile"
	repo "golang-project/internal/repository/user"
	svc "golang-project/internal/service/profile"
)

// NewRegistry returns new resource handler for profile API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, svc.NewService(repo.NewRepository(db)))
}
