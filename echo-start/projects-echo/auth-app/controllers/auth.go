package controllers

import (
	"authapp/database"
	"authapp/models"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to hash password"})
	}
	user.Password = string(hash)

	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusConflict, echo.Map{"message": "User is already registered in the database"})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "User Registered"})
}

func Login(c echo.Context) error{
	req := new(models.User)
	if err := c.Bind(req); err != nil {
		return err
	}
	var user models.User
	database.DB.First(&user,"email=?",req.Email);

	if user.ID==0{
		return c.JSON(http.StatusUnauthorized, echo.Map{"message":"Invalid Login Credentials"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password));

	if err!= nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message":"Invalid Password"})
	}

	claims := jwt.MapClaims{
		"user_id" : user.ID,
		"exp" : time.Now().Add(time.Hour*72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token":t});
}

func Profile(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message":"you are authenticated user and now you can view the profile"})
}