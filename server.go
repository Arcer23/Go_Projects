package main

import (
	"net/http"

	"todo/database"

	"github.com/labstack/echo/v4"
)

func main(){
	server := echo.New();
	server.GET("/home",func(c echo.Context)error{
		return c.JSON(http.StatusOK,"This is Pranish and I am building a todo app with golang and echo")
	})

	database.DatabaseInit();
	gorm:= database.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping();
	server.Logger.Fatal(server.Start(":8080"))

}