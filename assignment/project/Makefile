# Set env SERVER_ENV=dev before run -g ./server/engine.go
serve:
	go run main.go serve
swag:
	swag fmt
	swag init --parseDependency --parseDependencyLevel 3 -g main.go -g handler.go -d ./internal/handler -o ./docs/swagger
migrate:
	go run main.go migration migrate --schema --data
rollback:
	go run main.go migration rollback --schema --data