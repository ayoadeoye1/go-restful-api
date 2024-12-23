package router

import (
	usercontroller "github.com/ayoadeoye1/insta-shop-screening/controller/user_controller"
	"github.com/ayoadeoye1/insta-shop-screening/middleware"
	"github.com/gin-gonic/gin"

	//swagger files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userController *usercontroller.UserController) *gin.Engine {
	routes := gin.Default()

	// routes.Use(cors.New(cors.Config{
	//     AllowOrigins:     []string{"*"}, github.com/gin-contrib/cors
	//     AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//     AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
	//     ExposeHeaders:    []string{"Content-Length"},
	//     AllowCredentials: true,
	// }))

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := routes.Group("/api/v1")
	userRouter := router.Group("/user")
	userRouter.POST("/signup/admin", userController.CreateAdminUser)
	userRouter.POST("/signup", userController.CreateUser)
	userRouter.POST("/signin", userController.SignIn)

	userRouter.Use(middleware.AdminAuth)
	{
		userRouter.GET("/fetchall", userController.GetUsers)
	}

	return routes
}
