package responses

type ProductResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Currency    string   `json:"currency"`
	Category    string   `json:"category"`
	Brand       *string  `json:"brand"`
	Stock       uint     `json:"stock"`
	Images      []string `json:"images"`
}
