package userhandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"test/app"
	"test/constants"
	"test/models"
	"test/repository"
	"test/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByID(t *testing.T) {
	// Setup Mock DB
	db, mock := test.SetupMockDB(t)
	defer db.Close()

	// Setup Repository and Handler
	repo := repository.NewRepository(db)
	handler := app.SetupApp(repo)
	userHandler := NewUserHandler(handler)

	req := models.RequestID{
		ID: 1,
	}

	query := "SELECT id, username, email, password, status, created_at, updated_at FROM users WHERE id = \\$1 AND status = 'active'"

	mockFindUser := sqlmock.NewRows([]string{"id", "username", "email", "password", "status", "created_at", "updated_at"}).
		AddRow(1, "testuser", "testuser@example.com", "hashedpassword", "active", "2024-08-09 15:01:54", "2024-08-09 15:01:54")

	// Helper function to setup and execute requests
	executeRequest := func(
		t *testing.T,
		reqBody models.RequestID,
		expected models.TestingHandlerExpected,
	) {
		_, rec, ctx := test.NewRequestRecorder(
			models.TestingHandlerRequest{
				Method: http.MethodPost,
				Path:   "/api/v1/private/user/findbyid",
				Body:   reqBody,
			},
		)

		// Jalankan handler
		err := userHandler.FindUserByID(ctx)
		assert.Nil(t, err)
		assert.Equal(t, expected.StatusCode, rec.Code)

		// Periksa response
		var result models.Response
		err = json.Unmarshal(rec.Body.Bytes(), &result)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected.Success, result.Success)
		assert.Equal(t, expected.Code, result.StatusCode)
		assert.Contains(t, result.Message, expected.Message)
	}
	t.Run("Error Validation", func(t *testing.T) {
		executeRequest(t,
			models.RequestID{},
			models.TestingHandlerExpected{
				StatusCode: http.StatusBadRequest,
				Code:       constants.VALIDATE_ERROR_CODE,
				Success:    false,
				Message:    "Field 'ID' is required",
			})
	})

	t.Run("Error Case", func(t *testing.T) {
		mock.ExpectQuery(query).
			WithArgs(4).
			WillReturnError(errors.New("error scanning row"))

		executeRequest(t,
			models.RequestID{
				ID: 4,
			},
			models.TestingHandlerExpected{
				StatusCode: http.StatusInternalServerError,
				Code:       constants.SYSTEM_ERROR_CODE,
				Success:    false,
				Message:    "error scanning row",
			})
	})

	t.Run("Success Case", func(t *testing.T) {
		mock.ExpectQuery(query).
			WithArgs(1).
			WillReturnRows(mockFindUser)
		executeRequest(t,
			req,
			models.TestingHandlerExpected{
				StatusCode: http.StatusOK,
				Code:       constants.SUCCESS_CODE,
				Success:    true,
				Message:    constants.EMPTY_VALUE,
			})
	})
}
