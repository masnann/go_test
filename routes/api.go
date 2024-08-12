package routes

import (
	"test/handler"
	userhandler "test/handler/userHandler"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo, handler handler.Handler) {

	//public := e.Group("/api/v1/public")
	userHandler := userhandler.NewUserHandler(handler)

	private := e.Group("/api/v1/private")

	userGroup := private.Group("/user")
	userGroup.POST("/findbyid", userHandler.FindUserByID)

}
