package router

import (
	ordercontroller "github.com/ayoadeoye1/insta-shop-screening/controller/order_controller"
	productcontroller "github.com/ayoadeoye1/insta-shop-screening/controller/product_controller"
	usercontroller "github.com/ayoadeoye1/insta-shop-screening/controller/user_controller"
	"github.com/ayoadeoye1/insta-shop-screening/middleware"
	"github.com/gin-gonic/gin"

	//swagger files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userController *usercontroller.UserController, productController *productcontroller.ProductController, orderController *ordercontroller.OrderController) *gin.Engine {
	routes := gin.Default()

	// routes.Use(cors.New(cors.Config{
	//     AllowOrigins:     []string{"*"}, github.com/gin-contrib/cors
	//     AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//     AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
	//     ExposeHeaders:    []string{"Content-Length"},
	//     AllowCredentials: true,
	// }))

	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.Static("/uploads", "./public/uploads")

	version := routes.Group("/api/v1")

	userRouter := version.Group("/user")
	{
		userRouter.POST("/signup/admin", userController.CreateAdminUser)
		userRouter.POST("/signup", userController.CreateUser)
		userRouter.POST("/signin", userController.SignIn)

		userRouter.Use(middleware.AdminAuth)
		{
			userRouter.GET("/fetchall", userController.GetUsers)
		}
	}

	productRouter := version.Group("/product")
	{
		productRouter.GET("/", productController.FindAllProducts)
		productRouter.GET("/:id", productController.FindProductById)

		productRouter.Use(middleware.AdminAuth)
		{
			productRouter.POST("/create", productController.CreateNewProduct)
			productRouter.PUT("/:id", productController.UpdateProduct)
			productRouter.DELETE("/:id", productController.DeleteProduct)
		}
	}

	orderRouter := version.Group("/order")
	{
		orderRouter.Use(middleware.UserAuth)
		{
			orderRouter.POST("/create", orderController.CreateOrder)
			orderRouter.GET("/all", orderController.GetUserOrders)
			orderRouter.PUT("/cancel/:id", orderController.CancelOrder)
		}
		orderRouter.Use(middleware.AdminAuth)
		{
			orderRouter.PUT("/change/status", orderController.UpdateOrderStatus)
		}
	}

	return routes
}
