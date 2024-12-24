package productservice

import (
	"errors"
	"fmt"

	"github.com/ayoadeoye1/insta-shop-screening/data/requests"
	"github.com/ayoadeoye1/insta-shop-screening/data/responses"
	"github.com/ayoadeoye1/insta-shop-screening/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ayoadeoye1/insta-shop-screening/models"
	productrepository "github.com/ayoadeoye1/insta-shop-screening/repository/product_repository"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository productrepository.ProductRepo
	Validate          *validator.Validate
}

func NewProductServiceImpl(productRepository productrepository.ProductRepo, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

func (p *ProductServiceImpl) Create(products requests.CreateProductReq, ctx *gin.Context) error {
	err := p.Validate.Struct(products)
	if err != nil {
		return err
	}

	uploadedImageURLs, err := helper.ProcessUploadedImages(products.Images, ctx)
	if err != nil {
		return fmt.Errorf("failed to upload images: %w", err)
	}

	productModel := models.Products{
		Name:        products.Name,
		Description: products.Description,
		Price:       products.Price,
		Currency:    products.Currency,
		Category:    products.Category,
		Brand:       products.Brand,
		Stock:       products.Stock,
		Images:      uploadedImageURLs,
	}

	err = p.ProductRepository.Add(productModel)
	if err != nil {
		return fmt.Errorf("failed to save product: %w", err)
	}

	return nil
}

func (p *ProductServiceImpl) FindAll() ([]responses.ProductResponse, error) {
	products, err := p.ProductRepository.FindAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []responses.ProductResponse{}, errors.New("error fetching all products")
		}
		return []responses.ProductResponse{}, err
	}

	var productResponses []responses.ProductResponse

	for _, product := range products {
		productResponse := responses.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Currency:    product.Currency,
			Category:    product.Category,
			Brand:       product.Brand,
			Stock:       product.Stock,
			Images:      product.Images,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}

func (p *ProductServiceImpl) Update(product requests.UpdateProductReq, ctx *gin.Context) error {
	err := p.Validate.Struct(product)
	if err != nil {
		return err
	}

	existingProduct, err := p.ProductRepository.FindById(int(product.ID))
	if err != nil {
		return fmt.Errorf("failed to find product with ID %d: %w", product.ID, err)
	}

	if product.Name != nil {
		existingProduct.Name = *product.Name
	}
	if product.Description != nil {
		existingProduct.Description = *product.Description
	}
	if product.Price != nil {
		existingProduct.Price = *product.Price
	}
	if product.Currency != nil {
		existingProduct.Currency = *product.Currency
	}
	if product.Category != nil {
		existingProduct.Category = *product.Category
	}
	if product.Brand != nil {
		existingProduct.Brand = product.Brand
	}
	if product.Stock != nil {
		existingProduct.Stock = *product.Stock
	}

	if len(product.Images) > 0 {
		uploadedImageURLs, err := helper.ProcessUploadedImages(product.Images, ctx)
		if err != nil {
			return fmt.Errorf("failed to upload images: %w", err)
		}
		existingProduct.Images = uploadedImageURLs
	}

	err = p.ProductRepository.Edit(*existingProduct)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}

func (p *ProductServiceImpl) FindById(productId int) (responses.ProductResponse, error) {
	product, err := p.ProductRepository.FindById(productId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return responses.ProductResponse{}, errors.New("product not found")
		}
		return responses.ProductResponse{}, err
	}

	productResponse := responses.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Currency:    product.Currency,
		Category:    product.Category,
		Brand:       product.Brand,
		Stock:       product.Stock,
		Images:      product.Images,
	}

	return productResponse, nil
}

func (p *ProductServiceImpl) Delete(productId int) error {
	err := p.ProductRepository.Remove(productId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	return nil
}
