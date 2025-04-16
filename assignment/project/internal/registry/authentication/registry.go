package authentication

import (
	"gorm.io/gorm"

	"golang-project/internal/handler"
	hdl "golang-project/internal/handler/authentication"
	repo "golang-project/internal/repository/user"
	svc "golang-project/internal/service/authentication"
	"golang-project/util/hashing"
)

// NewRegistry returns new resource handler for authentication API
func NewRegistry(route string, db *gorm.DB) handler.ResourceHandler {
	return hdl.NewHandler(route, svc.NewService(repo.NewRepository(db), hashing.NewBcrypt()))
}
