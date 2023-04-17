package server

import (
	"CrowdProject/internal/auth"
	authhttp "CrowdProject/internal/auth/delivery/http"
	authrepo "CrowdProject/internal/auth/repository/postgres"
	authusecase "CrowdProject/internal/auth/usecase"
	"CrowdProject/internal/product"

	producthttp "CrowdProject/internal/product/delivery/http"
	productrepo "CrowdProject/internal/product/repository/postgres"
	productusecase "CrowdProject/internal/product/usecase"

	"CrowdProject/internal/config"
	postgres "CrowdProject/pkg/client/postgresql"

	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	authUC auth.UseCase

	productUC product.UseCase
}

func NewApp(cfg *config.Config) *App {
	logrus.Info("Connecting to DB")
	db, err := postgres.New(fmt.Sprintf("postgresql://%s:%s@%s:%s?sslmode=disable",
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host, cfg.PostgreSQL.Port))
	if err != nil {
		logrus.Infof("Problem with connection to db, err=%s", err)
	}

	userRepo := authrepo.NewUserRepository(db)

	productRepo := productrepo.NewRepository(db)

	return &App{
		httpServer: nil,
		authUC: authusecase.NewAuthUseCase(
			userRepo,
			cfg.Auth.HashSalt,
			[]byte(cfg.Auth.SigningKey),
			cfg.Auth.TokenTtl),
		productUC: productusecase.NewProductUseCase(productRepo),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Set up http handlers
	// SignUp/SignIn endpoints
	authhttp.RegisterHTTPEndpoints(router, a.authUC)
	producthttp.RegisterHTTPEndpoints(router, a.productUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
