package repository

import "test/models"

type UserRepositoryInterface interface {
	FindUserByID(id int64) (models.UserModels, error)
}
