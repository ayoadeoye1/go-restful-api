package productrepository

import "github.com/ayoadeoye1/insta-shop-screening/models"

type ProductRepo interface {
	Add(product models.Products) error
	Edit(product models.Products) error
	FindAll() ([]models.Products, error)
	FindById(userId int) (*models.Products, error)
	Remove(userId int) error
}
