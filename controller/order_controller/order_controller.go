package ordercontroller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ayoadeoye1/insta-shop-screening/data/requests"
	"github.com/ayoadeoye1/insta-shop-screening/helper"
	orderservice "github.com/ayoadeoye1/insta-shop-screening/services/order_service"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService orderservice.OrderService
}

func NewOrderController(orderService orderservice.OrderService) *OrderController {
	return &OrderController{
		OrderService: orderService,
	}
}

// CreateOrder godoc
// @Summary Create order
// @Description Endpoint to create new order, after being added to the cart on the UI
// @Param Authorization header string true "Bearer token"
// @Param order body requests.CreateOrderRequest true "Create Order"
// @Produce application/json
// @Tags Users
// @Success 200 {object} responses.Response{message=string} "Order created"
// @Router /api/v1/order/create [post]
func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var orderRequest requests.CreateOrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		helper.SendError(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err := ctrl.OrderService.Create(orderRequest, c)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, "Failed to create order", err.Error())
		return
	}

	helper.SendSuccess(c, http.StatusCreated, "Order created successfully", nil)
}

// FindAllUserOders godoc
// @Summary Get User's orders
// @Description Endpoint to Get all LoggedIn User's orders
// @Param Authorization header string true "Bearer token"
// @Produce application/json
// @Tags Users
// @Success 200 {object} responses.Response{message=string} "User Orders fetched"
// @Router /api/v1/order/all [get]
func (o *OrderController) GetUserOrders(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		helper.SendError(c, http.StatusUnauthorized, "User not authenticated", nil)
		return
	}

	orders, err := o.OrderService.GetUserOrders(userID.(uint))
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, "Failed to fetch orders", err.Error())
		return
	}

	helper.SendSuccess(c, http.StatusOK, "User's orders fetched", orders)
}

// Update Order godoc
// @Summary Update order status
// @Description Endpoint to update status of the order, to keep user posted about the delivery
// @Param Authorization header string true "Bearer token"
// @Param order body requests.UpdateOrderRequest true "Create Order"
// @Produce application/json
// @Tags Admin
// @Success 200 {object} responses.Response{message=string} "Order updated"
// @Router /api/v1/order/change/status [put]
func (ctrl *OrderController) UpdateOrderStatus(c *gin.Context) {
	var orderRequest *requests.UpdateOrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		helper.SendError(c, http.StatusBadRequest, "Invalid request body. Make sure to enter valid Enum string: Pending Processing Completed Cancelled", err.Error())
		return
	}

	err := ctrl.OrderService.UpdateStatus(orderRequest)
	if err != nil {
		helper.SendError(c, http.StatusInternalServerError, "Failed to update order status", err.Error())
		return
	}

	helper.SendSuccess(c, http.StatusCreated, "Order updated successfully", nil)
}

// Cancel Order godoc
// @Summary Cancel order
// @Description Endpoint for user to revoke order request, not to be processed anymore
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Cancel Order"
// @Produce application/json
// @Tags Users
// @Success 200 {object} responses.Response{message=string} "Order updated"
// @Router /api/v1/order/cancel/{id} [put]
func (ctrl *OrderController) CancelOrder(c *gin.Context) {
	orderID := c.Param("id")
	fmt.Print(orderID, "/asdfghjkl")
	id, err := strconv.Atoi(orderID)
	if err != nil {
		helper.SendError(c, http.StatusBadRequest, "Invalid order ID", err.Error())
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		helper.SendError(c, http.StatusUnauthorized, "User not authenticated", nil)
		return
	}

	err = ctrl.OrderService.CancelOrder(id, userID.(uint))
	if err != nil {
		helper.SendError(c, http.StatusBadRequest, "Failed to cancel order", err.Error())
		return
	}

	helper.SendSuccess(c, http.StatusCreated, "Order tarminated successfully", nil)
}
