package userservice

import (
	"errors"

	"github.com/ayoadeoye1/insta-shop-screening/data/requests"
	"github.com/ayoadeoye1/insta-shop-screening/data/responses"
	"gorm.io/gorm"

	"github.com/ayoadeoye1/insta-shop-screening/models"
	userrepository "github.com/ayoadeoye1/insta-shop-screening/repository/user_repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository userrepository.UserRepo
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository userrepository.UserRepo, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (u *UserServiceImpl) Create(users requests.CreateUserReq) error {
	err := u.Validate.Struct(users)
	if err != nil {
		return err
	}

	userModel := models.Users{
		Email:    users.Email,
		Password: users.Password,
		Role:     "user",
	}
	u.UserRepository.Add(userModel)
	return nil
}

func (u *UserServiceImpl) CreateAdmin(users requests.CreateUserReq) error {
	err := u.Validate.Struct(users)
	if err != nil {
		return err
	}

	userModel := models.Users{
		Email:    users.Email,
		Password: users.Password,
		Role:     "admin",
	}
	u.UserRepository.Add(userModel)
	return nil
}

func (u *UserServiceImpl) FindByEmail(userEmail string) (responses.UserResponse, error) {
	user, err := u.UserRepository.FindByEmail(userEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.UserResponse{}, errors.New("user not found")
		}
		return responses.UserResponse{}, err
	}

	userResponse := responses.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}

	return userResponse, nil
}

func (u *UserServiceImpl) FindAll() ([]responses.UserResponse, error) {
	users, err := u.UserRepository.FindAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []responses.UserResponse{}, errors.New("error fetching all users")
		}
		return []responses.UserResponse{}, err
	}

	var userResponses []responses.UserResponse

	for _, user := range users {
		userResponse := responses.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Role:  user.Role,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}
