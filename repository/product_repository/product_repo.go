package productrepository

import (
	"errors"

	"github.com/ayoadeoye1/insta-shop-screening/models"
	"gorm.io/gorm"
)

type ProductRepoImpl struct {
	Db *gorm.DB
}

func NewProductRepoImpl(Db *gorm.DB) ProductRepo {
	return &ProductRepoImpl{Db: Db}
}

func (p *ProductRepoImpl) Add(products models.Products) error {
	result := p.Db.Create(&products)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *ProductRepoImpl) Edit(products models.Products) error {
	result := p.Db.Save(&products)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *ProductRepoImpl) FindAll() ([]models.Products, error) {
	var products []models.Products
	result := p.Db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (p *ProductRepoImpl) FindById(productId int) (*models.Products, error) {
	var product models.Products
	result := p.Db.First(&product, productId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &product, nil
}

func (p *ProductRepoImpl) Remove(productId int) error {
	result := p.Db.Delete(&models.Products{}, productId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
