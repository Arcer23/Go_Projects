package routes

import (
	"authapp/controllers"
	"authapp/middlewares"

	"github.com/labstack/echo/v4"
)


func InitRoutes(e *echo.Echo){
	e.POST("/register", controllers.Register);
	e.POST("/login", controllers.Login);
	e.GET("/profile",controllers.Profile,middlewares.JWTmiddleware);
}