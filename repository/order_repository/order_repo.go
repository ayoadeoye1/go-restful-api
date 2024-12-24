package orderrepository

import (
	"errors"

	"github.com/ayoadeoye1/insta-shop-screening/models"
	"gorm.io/gorm"
)

type OrderRepoImpl struct {
	Db *gorm.DB
}

func NewOrderRepoImpl(Db *gorm.DB) OrderRepo {
	return &OrderRepoImpl{Db: Db}
}

func (o *OrderRepoImpl) Add(orders models.Order) error {
	result := o.Db.Create(&orders)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (o *OrderRepoImpl) Edit(orders models.Order) error {
	result := o.Db.Save(&orders)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (o *OrderRepoImpl) FindAll(userId uint) ([]models.Order, error) {
	var orders []models.Order
	result := o.Db.Where("user_id = ?", userId).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (o *OrderRepoImpl) FindById(OrderId int) (*models.Order, error) {
	var order models.Order
	result := o.Db.First(&order, OrderId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &order, nil
}

func (o *OrderRepoImpl) Remove(OrderId int) error {
	result := o.Db.Delete(&models.Order{}, OrderId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
