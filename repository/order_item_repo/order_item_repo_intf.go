package orderitemrepo

import "github.com/ayoadeoye1/insta-shop-screening/models"

type OrderItemRepo interface {
	Add(orderitem models.OrderItem) error
	FindAll(orderId uint) ([]models.OrderItem, error)
	Remove(orderItemId int) error
}
