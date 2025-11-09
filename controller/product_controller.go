package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) *ProductController {
	return &ProductController{
		productUsecase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		response := model.Response{
			Message: "Erro ao buscar produtos: " + err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	// Garantir que sempre retorne um array vazio se não houver produtos
	// Isso evita retornar null
	if products == nil {
		products = []model.Product{}
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {

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

func (p *ProductController) GetProductById(ctx *gin.Context) {

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

func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "ID do produto não pode ser nulo",
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

	err = p.productUsecase.DeleteProduct(productId)
	if err != nil {
		// Verifica se é erro de produto não encontrado
		if strings.Contains(err.Error(), "não encontrado") {
			response := model.Response{
				Message: err.Error(),
			}
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		response := model.Response{
			Message: "Erro ao deletar produto: " + err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := model.Response{
		Message: "Produto deletado com sucesso",
	}
	ctx.JSON(http.StatusOK, response)
}
