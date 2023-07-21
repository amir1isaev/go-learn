package core

import (
	"context"
	"errors"
	"fmt"
	"go-learn/internal/config"
	"go-learn/internal/delivery/handlers"
	"go-learn/internal/delivery/http"
	repository "go-learn/internal/repository/postgres"
	"go-learn/internal/usecase/user"
	"go-learn/pkg/database/postgres"

	"github.com/gin-gonic/gin"
)

type Core struct {
	isActive   bool
	ctx        context.Context
	httpServer *http.Server
	pgDB       *postgres.DB
}

func New(ctx context.Context) *Core {
	return &Core{ctx: ctx}
}

func (c *Core) Run() error {
	cfg := config.Get()

	db, err := postgres.NewDatabase(cfg.Storage.Postgres)

	if err != nil {
		return fmt.Errorf("postgres.NewDatabase: %w", err)
	}

	c.pgDB = db

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	repo := repository.New(c.pgDB)

	userUC := user.NewUseCase(repo)

	httpHandlers := handlers.NewHandlers(
		userUC,
	)

	c.httpServer = http.NewServer(router, httpHandlers)

	fmt.Println("httpServer", "http://localhost:8081")

	if err := c.httpServer.Start(); err != nil {
		return fmt.Errorf("c.httpServer.Start: %w", err)
	}

	c.isActive = true
	return nil
}

func (c *Core) Stop() error {
	if !c.isActive {
		return errors.New("core is inactive")
	}

	c.pgDB.Close()
	if err := c.httpServer.Stop(c.ctx); err != nil {
		return fmt.Errorf("c.httpServer.Stop: %w", err)
	}

	return nil
}
