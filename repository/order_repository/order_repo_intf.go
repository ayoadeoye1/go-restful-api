package orderrepository

import "github.com/ayoadeoye1/insta-shop-screening/models"

type OrderRepo interface {
	Add(order models.Order) error
	Edit(order models.Order) error
	FindAll(userId uint) ([]models.Order, error)
	FindById(orderId int) (*models.Order, error)
	Remove(orderId int) error
}
