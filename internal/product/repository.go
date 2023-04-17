package product

import (
	"CrowdProject/internal/models"
	"context"
)

type Repository interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	GetDetailProduct(ctx context.Context, id string) (*models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	UpdateProduct(ctx context.Context, id string) (*models.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}
