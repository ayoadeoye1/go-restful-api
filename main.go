package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ayoadeoye1/insta-shop-screening/config"
	productcontroller "github.com/ayoadeoye1/insta-shop-screening/controller/product_controller"
	usercontroller "github.com/ayoadeoye1/insta-shop-screening/controller/user_controller"
	"github.com/ayoadeoye1/insta-shop-screening/helper"
	"github.com/ayoadeoye1/insta-shop-screening/models"
	productrepository "github.com/ayoadeoye1/insta-shop-screening/repository/product_repository"
	userrepository "github.com/ayoadeoye1/insta-shop-screening/repository/user_repository"
	"github.com/ayoadeoye1/insta-shop-screening/router"
	productservice "github.com/ayoadeoye1/insta-shop-screening/services/product_service"
	userservice "github.com/ayoadeoye1/insta-shop-screening/services/user_service"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	//docs
	_ "github.com/ayoadeoye1/insta-shop-screening/docs"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	db := config.ConnectDB()
	validate := validator.New()

	db.Table("users").AutoMigrate(&models.Users{})
	usersRepository := userrepository.NewUserRepoImpl(db)
	usersService := userservice.NewUserServiceImpl(usersRepository, validate)
	userController := usercontroller.NewUserController(*usersService)

	db.Table("products").AutoMigrate(&models.Products{})
	productsRepository := productrepository.NewProductRepoImpl(db)
	productsService := productservice.NewProductServiceImpl(productsRepository, validate)
	productController := productcontroller.NewProductController(*productsService)

	routes := router.SetupRouter(userController, productController)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: routes,
	}

	fmt.Printf("\n Server listening on: http://localhost:%d \n", port)

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
