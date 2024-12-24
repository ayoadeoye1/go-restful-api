package orderitemrepo

import (
	"github.com/ayoadeoye1/insta-shop-screening/models"
	"gorm.io/gorm"
)

type OrderItemRepoImpl struct {
	Db *gorm.DB
}

func NewOrderItemRepoImpl(Db *gorm.DB) OrderItemRepo {
	return &OrderItemRepoImpl{Db: Db}
}

func (o *OrderItemRepoImpl) Add(orderItems models.OrderItem) error {
	result := o.Db.Create(&orderItems)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (o *OrderItemRepoImpl) FindAll(orderItemId uint) (orderItems []models.OrderItem, err error) {
	result := o.Db.Where("order_id = ?", orderItemId).Find(&orderItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return orderItems, nil
}

func (o *OrderItemRepoImpl) Remove(orderItemId int) error {
	result := o.Db.Delete(&models.OrderItem{}, orderItemId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
