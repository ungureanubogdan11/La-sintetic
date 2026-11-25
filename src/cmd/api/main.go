package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ungureanubogdan11/La-sintetic/src/internal/apperror"
	"github.com/ungureanubogdan11/La-sintetic/src/internal/config"
	"github.com/ungureanubogdan11/La-sintetic/src/internal/database"
	"github.com/ungureanubogdan11/La-sintetic/src/internal/handlers"
	"github.com/ungureanubogdan11/La-sintetic/src/internal/httpx"
	"github.com/ungureanubogdan11/La-sintetic/src/internal/middleware"
)

func main() {
	if err := run(); err != nil {
		slog.Error("fatal", "err", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		return err
	}
	if !cfg.IsDev() {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.Connect(ctx, cfg)
	if err != nil {
		return err
	}
	if sqlDB, err := db.DB(); err == nil {
		defer sqlDB.Close()
	}

	health := handlers.NewHealthHandler(db)

	r := gin.New()
	r.Use(gin.Logger(), middleware.Recovery())
	r.HandleMethodNotAllowed = true
	r.NoRoute(func(c *gin.Context) { httpx.Error(c, apperror.NotFound("route not found")) })
	r.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": httpx.ErrorBody{
			Code: "method_not_allowed", Message: "method not allowed",
		}})
	})

	api := r.Group("/api")
	health.Register(api)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		slog.Info("api listening", "addr", srv.Addr, "env", cfg.Env)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
	}

	slog.Info("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return srv.Shutdown(shutdownCtx)
}
