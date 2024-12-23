package userservice

import (
	"github.com/ayoadeoye1/insta-shop-screening/data/requests"
	"github.com/ayoadeoye1/insta-shop-screening/data/responses"
	"github.com/ayoadeoye1/insta-shop-screening/models"
)

type UserService interface {
	Create(user requests.CreateUserReq)
	CreateAdmin(user requests.CreateUserReq)
	Update(user models.Users)
	Delete(userId int)
	FindById(userId int) (user responses.UserResponse, err error)
	FindByEmail(userEmail string) (user responses.UserResponse, err error)
	FindAll() (users []responses.UserResponse, err error)
}
