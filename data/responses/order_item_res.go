package responses

type OrderItemResponse struct {
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    uint   `json:"quantity"`
	Price       int    `json:"price"`
	Subtotal    int    `json:"subtotal"`
}
