package postgres

import (
	"CrowdProject/internal/models"
	postgres "CrowdProject/pkg/client/postgresql"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
)

type ProductRepository struct {
	*postgres.Postgres
}

func NewRepository(db *postgres.Postgres) *ProductRepository {
	return &ProductRepository{db}
}

func (p ProductRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	logrus.Info("Create Sql query")
	sql, args, err := p.Builder.
		Insert("product").
		Columns("title, cost, description, author_id, category, rate").
		Values(product.Title, product.Cost, product.Description, product.Author.ID, product.Category, 0.0).
		ToSql()
	logrus.Info(product.Author.ID)
	if err != nil {
		logrus.Errorf("Product - Repository - CreateProduct. Err = %s", err)
		return err
	}

	logrus.Info("Execute Sql query")
	_, err = p.Pool.Exec(ctx, sql, args...)
	if err != nil {
		logrus.Errorf("Product - Repository - CreateProduct. Err = %s", err)
		return err
	}
	return nil
}

func (p ProductRepository) GetDetailProduct(ctx context.Context, id string) (*models.Product, error) {
	logrus.Info("Create Sql query")
	sql, args, err := p.Builder.
		Select("product_id, title, cost, Description, a.username, category, rate").
		Where(squirrel.Eq{"product_id": id}).
		From("product").
		InnerJoin("auth a on a.auth_id = product.author_id").
		ToSql()
	if err != nil {
		logrus.Errorf("Product - Repository - GetDetailProduct. Err = %s", err)
		return nil, err
	}

	logrus.Info(sql)

	product := new(models.Product)
	logrus.Info("Execute Sql query")
	err = p.Pool.QueryRow(ctx, sql, args...).Scan(&product.Id, &product.Title, &product.Cost,
		&product.Description, &product.Author.Username, &product.Category, &product.Rate)
	if err != nil {
		logrus.Errorf("Product - Repository - GetDetailProduct. Err = %s", err)
		return nil, err
	}

	return product, nil
}

func (p ProductRepository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	logrus.Info("Create Sql query")
	sql, args, err := p.Builder.
		Select("product_id, title, cost, author_id, category, rate").
		From("product").
		InnerJoin("auth a on a.auth_id = product.author_id").
		ToSql()

	logrus.Infof("Execute Sql query, sql = %s", sql)
	rows, err := p.Pool.Query(ctx, sql, args...)
	if err != nil {
		logrus.Errorf("Product - Repository - GetAllProducts. Err = %s", err)
		return nil, err
	}

	products := make([]models.Product, 0)

	for rows.Next() {
		product := models.Product{}
		err = rows.Scan(&product.Id, &product.Title, &product.Cost, &product.Author.ID, &product.Category, &product.Rate)
		if err != nil {
			logrus.Errorf("Product - Repository - GetAllProducts. Err = %s", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (p ProductRepository) UpdateProduct(ctx context.Context, id string) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) DeleteProduct(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
