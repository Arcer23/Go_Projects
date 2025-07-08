package main

import (
	"authapp/database"
	"authapp/models"
	"authapp/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)
func main(){	

	err := godotenv.Load();
	if err!=nil {
		log.Fatal("Error loading in the .env file");
	}

	database.ConnectDB();
	database.DB.AutoMigrate(&models.User{})

	e := echo.New()
	routes.InitRoutes(e)
	 
	e.Logger.Fatal(e.Start(":8080"))
}
