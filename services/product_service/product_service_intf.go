package productservice

import (
	"github.com/ayoadeoye1/insta-shop-screening/data/requests"
	"github.com/ayoadeoye1/insta-shop-screening/data/responses"
	"github.com/ayoadeoye1/insta-shop-screening/models"
)

type ProductService interface {
	Create(product requests.CreateProductReq)
	Update(product models.Products)
	Delete(productId int)
	FindById(productId int) (product responses.ProductResponse, err error)
	FindAll() (products []responses.ProductResponse, err error)
}
