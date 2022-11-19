package main

import (
	"github.com/gin-gonic/gin"
	"github.com/techlateef/tech-lateef-gol/config"
	"github.com/techlateef/tech-lateef-gol/controllers"
	"github.com/techlateef/tech-lateef-gol/repository"
	"github.com/techlateef/tech-lateef-gol/services"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	userRepository repository.UserRepository = repository.NewUserRepository(db)
	userService    services.UserService      = services.NewUserServices(userRepository)

	userController controllers.UserController = controllers.NewUserController(userService)
)

func main() {
	defer config.SetupDatabaseConnection()
	r := gin.Default()
	authRoutes := r.Group("User")

	{
		authRoutes.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"Message": "Welcom",
			})
		})
		authRoutes.POST("/create", userController.CreateUser)
		authRoutes.PUT("/:id", userController.UpdateUser)
		authRoutes.GET("/AllUsers", userController.GetAllUsers)
		authRoutes.GET("/:id", userController.GetUserById)
		authRoutes.DELETE("/:id", userController.DeleteUser)
	}

	r.Run()

}
