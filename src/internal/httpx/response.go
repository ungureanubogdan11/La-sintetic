package httpx

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ungureanubogdan11/La-sintetic/src/internal/apperror"
)

type ErrorBody struct {
	Code    apperror.Kind `json:"code"`
	Message string        `json:"message"`
	Details any           `json:"details,omitempty"`
}

type PageMeta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

func NewPageMeta(page, pageSize int, total int64) PageMeta {
	pages := 0
	if pageSize > 0 {
		pages = int((total + int64(pageSize) - 1) / int64(pageSize))
	}
	return PageMeta{Page: page, PageSize: pageSize, Total: total, TotalPages: pages}
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, gin.H{"data": data})
}

func Page(c *gin.Context, data any, meta PageMeta) {
	c.JSON(http.StatusOK, gin.H{"data": data, "meta": meta})
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

var statusByKind = map[apperror.Kind]int{
	apperror.KindInvalid:      http.StatusBadRequest,
	apperror.KindUnauthorized: http.StatusUnauthorized,
	apperror.KindForbidden:    http.StatusForbidden,
	apperror.KindNotFound:     http.StatusNotFound,
	apperror.KindConflict:     http.StatusConflict,
	apperror.KindInternal:     http.StatusInternalServerError,
}

func Error(c *gin.Context, err error) {
	e := apperror.As(err)
	status, ok := statusByKind[e.Kind]
	if !ok {
		status = http.StatusInternalServerError
	}
	if status >= 500 {
		slog.ErrorContext(c.Request.Context(), "request failed",
			"method", c.Request.Method, "path", c.Request.URL.Path, "err", err)
	}
	c.AbortWithStatusJSON(status, gin.H{"error": ErrorBody{
		Code:    e.Kind,
		Message: e.Message,
		Details: e.Details,
	}})
}
