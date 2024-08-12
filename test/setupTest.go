package test

import (
	"test/repository/mocks"
	"test/service"
	userservice "test/service/userService"
	"testing"
)

type TestSuite struct {
	UserRepo *mocks.UserRepositoryInterface

	Service     service.Service
	UserService userservice.UserService
}

func SetupTestCase(t *testing.T) *TestSuite {

	userRepo := mocks.NewUserRepositoryInterface(t)

	svc := service.NewService(userRepo)

	userService := userservice.NewUserService(svc)

	return &TestSuite{

		UserRepo: userRepo,

		Service:     svc,
		UserService: userService,
	}
}
