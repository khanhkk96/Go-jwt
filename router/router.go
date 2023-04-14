package router

import (
	"golang-jwttoken/controllers"
	"golang-jwttoken/middleware"
	"golang-jwttoken/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(userRepository repository.UserRepository, authenticationController *controllers.AuthenticationController, userController *controllers.UserController) *gin.Engine {
	service := gin.Default()
	//add swagger
	service.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	service.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome")
	})

	router := service.Group("/api")

	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	userRouter := router.Group("/user")
	userRouter.GET("/", middleware.DeserializeUser(userRepository), userController.GetUsers)
	userRouter.PUT("/:userId", middleware.DeserializeUser(userRepository), userController.UpdateUser)
	userRouter.DELETE("/:userId", middleware.DeserializeUser(userRepository), userController.Remove)
	userRouter.DELETE("/delete/:userId", middleware.DeserializeUser(userRepository), userController.Delete)

	return service
}
