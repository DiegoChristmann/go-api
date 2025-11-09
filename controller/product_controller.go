package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	insertedProduct, err := p.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

	}

	ctx.JSON(http.StatusCreated, insertedProduct)

}

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "ID do produto nao pode ser nula",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if err == nil && product == nil {
		response := model.Response{
			Message: "O produto nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
