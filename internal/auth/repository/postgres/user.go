package postgres

import (
	"CrowdProject/internal/models"
	postgres "CrowdProject/pkg/client/postgresql"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	*postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	logrus.Info("Making the SQL query")
	sql, args, err := r.Builder.
		Insert("auth").
		Columns("username, password").
		Values(user.Username, user.Password).
		ToSql()

	if err != nil {
		logrus.Info(err)
		return err
	}

	logrus.Info("Execute SQL query")
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		logrus.Info(err)
		return err
	}
	return nil
}

func (r UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	logrus.Info("Making the SQL query")
	sql, args, err := r.Builder.
		Select("auth_id, username, password").
		Where(squirrel.Eq{"username": username, "password": password}).
		From("auth").ToSql()
	if err != nil {
		return nil, err
	}

	logrus.Info("Execute SQL query")
	user := models.User{}
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		logrus.Info(err)
		return nil, err
	}

	return &user, nil
}
