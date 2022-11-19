package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techlateef/tech-lateef-gol/models"
)

type RegistrationInput struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegistrationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// SAVING  THE CREDENTIAL INTO THE DATABASE

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration succes"})

}
