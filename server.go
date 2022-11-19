package main

import (
	"github.com/gin-gonic/gin"
	"github.com/techlateef/tech-lateef-gol/controllers"
	"github.com/techlateef/tech-lateef-gol/models"
)

func main() {

	models.ConnectDatabase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)

	r.Run()

}
