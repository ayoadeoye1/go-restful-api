package responses

import "time"

type OrderResponse struct {
	ID          uint                `json:"id"`
	UserID      uint                `json:"user_id"`
	UserName    string              `json:"user_name"`
	TotalAmount int                 `json:"total_amount"`
	Address     string              `json:"address"`
	Status      string              `json:"status"`
	Items       []OrderItemResponse `json:"items"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}
