package orderitemservice

import (
	"github.com/ayoadeoye1/insta-shop-screening/data/requests"
	"github.com/ayoadeoye1/insta-shop-screening/data/responses"
)

type OrderService interface {
	Create(orderitem requests.CreateOrderItemRequest)
	FindAll() (orders []responses.OrderItemResponse, err error)
	Delete(orderitemId int)
}
