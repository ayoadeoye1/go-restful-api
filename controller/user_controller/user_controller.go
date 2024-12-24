package usercontroller

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ayoadeoye1/insta-shop-screening/data/requests"
	"github.com/ayoadeoye1/insta-shop-screening/data/responses"
	"github.com/ayoadeoye1/insta-shop-screening/helper"
	userservice "github.com/ayoadeoye1/insta-shop-screening/services/user_service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService userservice.UserServiceImpl
}

func NewUserController(userService userservice.UserServiceImpl) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUsers godoc
// @Summary Sign-Up New User
// @Description An endpoint for a new user to sign-up
// @Param users body requests.CreateUserReq true "Create Users"
// @Accept json
// @Produce application/json
// @Tags Users
// @Success 200 {object} responses.Response{}
// @Router /api/v1/user/signup [post]
func (userController *UserController) CreateUser(ctx *gin.Context) {
	createUserRequest := requests.CreateUserReq{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Internal Error, Please Try Again", err.Error())
		return
	}

	validate := validator.New()

	err = validate.Struct(createUserRequest)
	if err != nil {
		// validationErrors := helper.FormatValidationErrors(err)
		// helper.SendError(ctx, http.StatusBadRequest, "Validation Error", validationErrors)
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := helper.FormatValidationErrors(validationErrs)
			helper.SendError(ctx, http.StatusBadRequest, "Validation Error", formattedErrors)
		} else {
			helper.SendError(ctx, http.StatusBadRequest, "Invalid JSON input", err.Error())
		}
		return
	}

	loginUser, err := userController.userService.FindByEmail(createUserRequest.Email)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Error occured while searching for user email", err.Error())
		return
	}

	if loginUser != (responses.UserResponse{}) {
		helper.SendError(ctx, http.StatusBadRequest, "Account with email already exist", nil)
		return
	}

	salt, err := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Internal Error, Please Try Again", err.Error())
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), salt)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Internal Error, Please Try Again", err.Error())
		return
	}

	createUserRequest.Password = string(hash)

	userController.userService.Create(createUserRequest)

	helper.SendSuccess(ctx, http.StatusOK, "Sign Up Successful", nil)
}

// CreateUsers godoc
// @Summary Create New Admin User
// @Description An endpoint for a new admin user to sign-up
// @Param users body requests.CreateUserReq true "Create Users"
// @Accept json
// @Produce application/json
// @Tags Admin
// @Success 200 {object} responses.Response{}
// @Router /api/v1/user/signup/admin [post]
func (userController *UserController) CreateAdminUser(ctx *gin.Context) {
	createUserRequest := requests.CreateUserReq{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Internal Error, Please Try Again", err.Error())
		return
	}

	validate := validator.New()

	err = validate.Struct(createUserRequest)
	if err != nil {
		// validationErrors := helper.FormatValidationErrors(err)
		// helper.SendError(ctx, http.StatusBadRequest, "Validation Error", validationErrors)
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := helper.FormatValidationErrors(validationErrs)
			helper.SendError(ctx, http.StatusBadRequest, "Validation Error", formattedErrors)
		} else {
			helper.SendError(ctx, http.StatusBadRequest, "Invalid JSON input", err.Error())
		}
		return
	}

	loginUser, err := userController.userService.FindByEmail(createUserRequest.Email)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Error occured while searching for user email", err.Error())
		return
	}

	if loginUser != (responses.UserResponse{}) {
		helper.SendError(ctx, http.StatusBadRequest, "Account with email already exist", nil)
		return
	}

	salt, err := strconv.Atoi(os.Getenv("BCRYPT_SALT"))
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Internal Error, Please Try Again", err.Error())
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), salt)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Internal Error, Please Try Again", err.Error())
		return
	}

	createUserRequest.Password = string(hash)

	userController.userService.CreateAdmin(createUserRequest)

	helper.SendSuccess(ctx, http.StatusOK, "Created New Admin Account Successfully", nil)
}

// UserSignIn godoc
// @Summary Sign-In User
// @Description An endpoint for a user to sign-in
// @Param users body requests.LoginReq true "User SignIn"
// @Accept json
// @Produce application/json
// @Tags Users
// @Success 200 {object} responses.Response{}
// @Router /api/v1/user/signin [post]
func (userController *UserController) SignIn(ctx *gin.Context) {
	LoginReq := requests.LoginReq{}
	err := ctx.ShouldBindJSON(&LoginReq)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Internal Error, Please Try Again", err.Error())
		return
	}

	validate := validator.New()

	err = validate.Struct(LoginReq)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := helper.FormatValidationErrors(validationErrs)
			helper.SendError(ctx, http.StatusBadRequest, "Validation Error", formattedErrors)
		} else {
			helper.SendError(ctx, http.StatusBadRequest, "Invalid JSON input", err.Error())
		}
		return
	}

	loginUser, err := userController.userService.FindByEmail(LoginReq.Email)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Error occured while searching for user", err.Error())
		return
	}

	if loginUser == (responses.UserResponse{}) {
		helper.SendError(ctx, http.StatusBadRequest, "Account with email does not exist", nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(LoginReq.Password))
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Incorrect password", err.Error())
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   loginUser.ID,
		"mail": loginUser.Email,
		"role": loginUser.Role,
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Signing token error", err.Error())
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Sign Up Successful", tokenString)
}

// UserSignIn godoc
// @Summary Get Users
// @Description An endpoint to get Users
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Produce application/json
// @Tags Admin
// @Success 200 {object} []responses.Response{}
// @Router /api/v1/user/fetchall [get]
func (userController *UserController) GetUsers(ctx *gin.Context) {
	users, err := userController.userService.FindAll()
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Error Fetching Users", err.Error())
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Users Fetch Successful", users)
}
