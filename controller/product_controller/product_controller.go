package productcontroller

import (
	"net/http"
	"strconv"

	"github.com/ayoadeoye1/insta-shop-screening/data/requests"
	"github.com/ayoadeoye1/insta-shop-screening/helper"
	productservice "github.com/ayoadeoye1/insta-shop-screening/services/product_service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService productservice.ProductServiceImpl
}

func NewProductController(productService productservice.ProductServiceImpl) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// CreateProducts godoc
// @Summary Create a new product
// @Description Endpoint to create a new product with details and multiple images
// @Param Authorization header string true "Bearer token"
// @Param name formData string true "Product Name"
// @Param description formData string true "Product Description"
// @Param price formData int true "Product Price"
// @Param currency formData string true "Currency"
// @Param category formData string true "Product Category"
// @Param brand formData string true "Product Brand"
// @Param stock formData int true "Available Stock"
// @Param images formData file true "Product Images (can upload multiple)" multiple
// @Accept multipart/form-data
// @Produce application/json
// @Tags Admin
// @Success 200 {object} responses.Response{message=string} "Product created successfully"
// @Router /api/v1/product/create [post]
func (productController *ProductController) CreateNewProduct(ctx *gin.Context) {
	var createProductRequest requests.CreateProductReq

	err := ctx.ShouldBind(&createProductRequest)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Invalid image upload", err.Error())
		return
	}

	images := form.File["images"]
	createProductRequest.Images = images

	err = productController.productService.Create(createProductRequest, ctx)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "Failed to create product", err.Error())
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Product created successfully", nil)
}

// FindAllProducts godoc
// @Summary Get all products
// @Description Endpoint to get a list of all products
// @Param Authorization header string true "Bearer token"
// @Produce application/json
// @Tags Users
// @Success 200 {array} responses.ProductResponse "List of products"
// @Router /api/v1/product [get]
func (productController *ProductController) FindAllProducts(ctx *gin.Context) {
	products, err := productController.productService.FindAll()
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "Failed to fetch products", err.Error())
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Products fetched successfully", products)
}

// FindProductById godoc
// @Summary Get a product by ID
// @Description Endpoint to get details of a product by its ID
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Product ID"
// @Produce application/json
// @Tags Users
// @Success 200 {object} responses.ProductResponse "Product details"
// @Router /api/v1/product/{id} [get]
func (productController *ProductController) FindProductById(ctx *gin.Context) {
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	product, err := productController.productService.FindById(id)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "Failed to fetch product", err.Error())
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Product found", product)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Endpoint to update an existing product's details
// @Param Authorization header string true "Bearer token"
// @Param id path int false "Product ID"
// @Param name formData string false "Product Name"
// @Param description formData string false "Product Description"
// @Param price formData int false "Product Price"
// @Param currency formData string false "Currency"
// @Param category formData string false "Product Category"
// @Param brand formData string false "Product Brand"
// @Param stock formData int false "Available Stock"
// @Param images formData file false "Product Images" multiple
// @Accept multipart/form-data
// @Produce application/json
// @Tags Admin
// @Success 200 {object} responses.Response{message=string} "Product updated successfully"
// @Router /api/v1/product/{id} [put]
func (productController *ProductController) UpdateProduct(ctx *gin.Context) {
	var updateProductRequest requests.UpdateProductReq
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	err = ctx.ShouldBind(&updateProductRequest)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Invalid image upload", err.Error())
		return
	}

	images := form.File["images"]
	updateProductRequest.Images = images

	updateProductRequest.ID = uint(id)

	err = productController.productService.Update(updateProductRequest, ctx)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "Failed to update product", err.Error())
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Product updated successfully", nil)
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Endpoint to delete a product using its ID
// @Param Authorization header string true "Bearer token"
// @Param id path int true "Product ID"
// @Produce application/json
// @Tags Admin
// @Success 200 {object} responses.Response{message=string} "Product deleted successfully"
// @Router /api/v1/product/{id} [delete]
func (productController *ProductController) DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	err = productController.productService.Delete(id)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "Failed to delete product", err.Error())
		return
	}

	helper.SendSuccess(ctx, http.StatusOK, "Product deleted successfully", nil)
}
