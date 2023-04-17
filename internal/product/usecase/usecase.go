package usecase

import (
	"CrowdProject/internal/models"
	"CrowdProject/internal/product"
	"context"
)

type ProductUseCase struct {
	repo product.Repository
}

func NewProductUseCase(repo product.Repository) *ProductUseCase {
	return &ProductUseCase{repo}
}

func (p ProductUseCase) CreateProduct(ctx context.Context, product *models.Product) error {
	err := p.repo.CreateProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductUseCase) GetDetailProduct(ctx context.Context, id string) (*models.Product, error) {
	detailProduct, err := p.repo.GetDetailProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return detailProduct, nil
}

func (p ProductUseCase) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	products, err := p.repo.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p ProductUseCase) UpdateProduct(ctx context.Context, id string) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductUseCase) DeleteProduct(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
