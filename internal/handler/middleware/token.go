package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// middleware to check secret api token in request
func (m *Middleware) ApiToken(ctx *gin.Context) {
	if ctx.GetHeader("X-API-Key") != m.token {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, 
			gin.H{
				"error": "Invalid or missing api key",
		})
	}

	ctx.Next()
}