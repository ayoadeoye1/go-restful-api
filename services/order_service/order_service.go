package orderservice

import (
	"fmt"

	"github.com/ayoadeoye1/insta-shop-screening/data/requests"
	"github.com/ayoadeoye1/insta-shop-screening/data/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ayoadeoye1/insta-shop-screening/models"
	orderitemrepo "github.com/ayoadeoye1/insta-shop-screening/repository/order_item_repo"
	orderrepository "github.com/ayoadeoye1/insta-shop-screening/repository/order_repository"
	"github.com/go-playground/validator/v10"
)

type OrderServiceImpl struct {
	OrderRepository     orderrepository.OrderRepo
	OrderItemRepository orderitemrepo.OrderItemRepo
	Validate            *validator.Validate
	Db                  *gorm.DB
}

func NewOrderServiceImpl(orderRepository orderrepository.OrderRepo, orderItemRepository orderitemrepo.OrderItemRepo, validate *validator.Validate, db *gorm.DB) *OrderServiceImpl {
	return &OrderServiceImpl{
		OrderRepository:     orderRepository,
		OrderItemRepository: orderItemRepository,
		Validate:            validate,
		Db:                  db,
	}
}

func (o *OrderServiceImpl) Create(orders requests.CreateOrderRequest, ctx *gin.Context) error {
	err := o.Validate.Struct(orders)
	if err != nil {
		return err
	}

	userID, ok := ctx.Get("userID")
	if !ok {
		return fmt.Errorf("user ID not found in context")
	}

	var orderItems []models.OrderItem
	totalAmount := 0
	for _, item := range orders.Items {
		// subtotal := int(item.Quantity) * item.Price
		// totalAmount += subtotal

		// orderItems = append(orderItems, models.OrderItem{
		// 	ProductID: item.ProductID,
		// 	Quantity:  item.Quantity,
		// 	Price:     item.Price,
		// 	Subtotal:  subtotal,
		// })

		var product models.Products
		if err := o.Db.Where("id = ?", item.ProductID).First(&product).Error; err != nil {
			return fmt.Errorf("failed to fetch product details for product ID %d: %w", item.ProductID, err)
		}

		subtotal := int(item.Quantity) * int(product.Price)
		totalAmount += subtotal

		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     int(product.Price),
			Subtotal:  subtotal,
		})
	}

	orderModel := models.Order{
		UserID:      userID.(uint),
		TotalAmount: totalAmount,
		Address:     orders.Address,
		Status:      "Pending",
	}

	tx := o.Db.Begin()
	if err := tx.Create(&orderModel).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to save order: %w", err)
	}

	for i := range orderItems {
		orderItems[i].OrderID = orderModel.ID
		if err := tx.Create(&orderItems[i]).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to save order item: %w", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (o *OrderServiceImpl) GetUserOrders(userId uint) (orders []responses.OrderResponse, err error) {
	var dbOrders []models.Order

	dbOrders, err = o.OrderRepository.FindAll(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %w", err)
	}

	for _, order := range dbOrders {
		orderItems, err := o.OrderItemRepository.FindAll(order.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch orderitems: %w", err)
		}

		var items []responses.OrderItemResponse
		for _, item := range orderItems {
			items = append(items, responses.OrderItemResponse{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
				Subtotal:  item.Subtotal,
			})
		}

		orders = append(orders, responses.OrderResponse{
			ID:          order.ID,
			UserID:      order.UserID,
			TotalAmount: order.TotalAmount,
			Address:     order.Address,
			Status:      order.Status,
			Items:       items,
		})
	}

	return orders, nil
}

func (o *OrderServiceImpl) UpdateStatus(order *requests.UpdateOrderRequest) error {
	err := o.Validate.Struct(order)
	if err != nil {
		return err
	}

	existingOrder, err := o.OrderRepository.FindById(int(order.ID))
	if err != nil {
		return fmt.Errorf("failed to find order with ID %d: %w", order.ID, err)
	}

	existingOrder.Status = *order.Status

	err = o.OrderRepository.Edit(*existingOrder)
	if err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	return nil
}

func (o *OrderServiceImpl) CancelOrder(orderId int, userId uint) error {
	existingOrder, err := o.OrderRepository.FindById(orderId)
	if err != nil {
		return fmt.Errorf("failed to find order with ID %d: %w", orderId, err)
	}

	if existingOrder.UserID != userId {
		return fmt.Errorf("you are not authorized to cancel this order, not created by you")
	}

	if existingOrder.Status != "Pending" {
		return fmt.Errorf("OrderId:%d can no longer be tarminated, because it's being processed", orderId)
	}

	existingOrder.Status = "Cancelled"

	err = o.OrderRepository.Edit(*existingOrder)
	if err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	return nil
}
