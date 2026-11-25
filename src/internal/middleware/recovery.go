package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ungureanubogdan11/La-sintetic/src/internal/apperror"
	"github.com/ungureanubogdan11/La-sintetic/src/internal/httpx"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, rec any) {
		httpx.Error(c, apperror.Internal(fmt.Errorf("panic: %v", rec)))
	})
}
