package http

import (
	"CrowdProject/internal/models"
	"CrowdProject/internal/product"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	useCase product.UseCase
}

func NewHandler(useCase product.UseCase) *Handler {
	return &Handler{useCase}
}

func (h Handler) GetAllProducts(c *gin.Context) {
	products, err := h.useCase.GetAllProducts(c.Request.Context())
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h Handler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	detailProduct, err := h.useCase.GetDetailProduct(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, detailProduct)
}

type CreateResponse struct {
	Title       string `json:"title,omitempty"`
	Cost        int    `json:"cost,omitempty"`
	Description string `json:"description,omitempty"`
	Author      int    `json:"author,omitempty"`
	Category    string `json:"category,omitempty"`
}

func (h Handler) CreateProduct(c *gin.Context) {
	inp := new(CreateResponse)

	err := c.BindJSON(inp)
	if err != nil {
		logrus.Errorf("Bad requset, err= %s", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	p := models.Product{
		Title:       inp.Title,
		Cost:        inp.Cost,
		Description: inp.Description,
		Author:      models.User{ID: inp.Author},
		Category:    inp.Category,
	}
	err = h.useCase.CreateProduct(c.Request.Context(), &p)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusCreated)
}
