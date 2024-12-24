package orderservice

import (
	"github.com/ayoadeoye1/insta-shop-screening/data/requests"
	"github.com/ayoadeoye1/insta-shop-screening/data/responses"
	"github.com/gin-gonic/gin"
)

type OrderService interface {
	Create(orders requests.CreateOrderRequest, ctx *gin.Context) error
	GetUserOrders(userId uint) ([]responses.OrderResponse, error)
	UpdateStatus(order *requests.UpdateOrderRequest) error
	CancelOrder(orderId int, userId uint) error
}
