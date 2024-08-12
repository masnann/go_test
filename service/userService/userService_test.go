package userservice_test

import (
	"errors"
	"test/models"
	"test/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindByID(t *testing.T) {
	ts := test.SetupTestCase(t)

	req := models.RequestID{
		ID: 1,
	}

	t.Run("Failure Case - Error FindUserByID", func(t *testing.T) {
		expectedErr := errors.New("user not found")

		ts.UserRepo.On("FindUserByID", req.ID).Return(models.UserModels{}, expectedErr).Once()

		result, err := ts.UserService.FindUserByID(req)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, models.UserModels{}, result)
		ts.UserRepo.AssertExpectations(t)
	})

	t.Run("Success Case - User Found", func(t *testing.T) {
		expected := models.UserModels{
			ID:       1,
			Username: "John Doe",
		}

		ts.UserRepo.On("FindUserByID", req.ID).Return(expected, nil).Once()

		result, err := ts.UserService.FindUserByID(req)

		assert.Nil(t, err)
		assert.Equal(t, expected, result)

		ts.UserRepo.AssertExpectations(t)
	})
}
