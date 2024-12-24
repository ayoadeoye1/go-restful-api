package requests

import "mime/multipart"

type CreateProductReq struct {
	Name        string                  `form:"name" binding:"required"`
	Description string                  `form:"description" binding:"required"`
	Price       int                     `form:"price" binding:"required"`
	Currency    string                  `form:"currency" binding:"required"`
	Category    string                  `form:"category" binding:"required"`
	Brand       *string                 `form:"brand" binding:"required"`
	Stock       uint                    `form:"stock" binding:"required"`
	Images      []*multipart.FileHeader `form:"images" binding:"required"`
}

type UpdateProductReq struct {
	ID          uint                    `gorm:"primaryKey;autoIncrement"`
	Name        *string                 `form:"name"`
	Description *string                 `form:"description"`
	Price       *int                    `form:"price"`
	Currency    *string                 `form:"currency"`
	Category    *string                 `form:"category"`
	Brand       *string                 `form:"brand"`
	Stock       *uint                   `form:"stock"`
	Images      []*multipart.FileHeader `form:"images"`
}
