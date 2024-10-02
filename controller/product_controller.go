package controller

import (
	"ApiGo/model"
	"ApiGo/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: usecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata frita",
			Price: 20,
		},
	}
	ctx.JSON(http.StatusOK, products)
}
