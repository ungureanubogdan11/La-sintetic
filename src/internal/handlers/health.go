package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ungureanubogdan11/La-sintetic/src/internal/httpx"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

func (h *HealthHandler) Register(r gin.IRouter) {
	r.GET("/health", h.Get)
}

func (h *HealthHandler) Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	dbStatus := "up"
	if sqlDB, err := h.db.DB(); err != nil || sqlDB.PingContext(ctx) != nil {
		dbStatus = "down"
	}

	body := gin.H{"status": "ok", "database": dbStatus}
	if dbStatus != "up" {
		body["status"] = "degraded"
		c.JSON(http.StatusServiceUnavailable, gin.H{"data": body})
		return
	}
	httpx.OK(c, body)
}
