package app

import (
	"test/handler"
	"test/repository"
	userrepository "test/repository/userRepository"
	"test/service"
	userservice "test/service/userService"
)

func SetupApp(repo repository.Repository) handler.Handler {
	userRepo := userrepository.NewUserRepository(repo)
	service := service.NewService(userRepo)

	userService := userservice.NewUserService(service)

	handler := handler.NewHandler(userService)

	return handler
}
