package http

import (
	"CrowdProject/internal/product"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc product.UseCase) {
	h := NewHandler(uc)

	productEndpoints := router.Group("/products")
	{
		productEndpoints.POST("", h.CreateProduct)
		productEndpoints.GET("", h.GetAllProducts)
		productEndpoints.GET("/:id", h.GetProduct)
	}

}
