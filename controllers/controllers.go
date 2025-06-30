package controllers

import (
	"net/http"
	"todo/database"
	"todo/models"

	"github.com/labstack/echo/v4"

)

func GetTodos(c echo.Context) error {
	var todos []models.Todo
	database.DB().Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

func CreateTodo()