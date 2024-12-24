package requests

type CreateOrderRequest struct {
	Address string                   `json:"address" binding:"required"`
	Items   []CreateOrderItemRequest `json:"items" binding:"required"`
}

type UpdateOrderRequest struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	// Address *string                   `json:"address"`
	// Items   []*CreateOrderItemRequest `json:"items"`
	Status *string `json:"status" binding:"oneof=Pending Processing Completed Cancelled"`
}
