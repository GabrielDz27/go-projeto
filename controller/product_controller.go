package model

import (
	"go-api/model"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//Usecase
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:            7,
			name_product:  "Computador",
			price_product: 1000,
		},
	}
}
