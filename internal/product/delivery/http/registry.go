package http

import (
	"CrowdProject/internal/product"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc product.UseCase) {
	h := NewHandler(uc)

	router.POST("", h.CreateProduct)
	router.GET("", h.GetAllProducts)
	router.GET("/:id", h.GetProduct)

}
