package service

import "test/models"

type UserServiceInterface interface {
	FindUserByID(req models.RequestID) (models.UserModels, error)
}
