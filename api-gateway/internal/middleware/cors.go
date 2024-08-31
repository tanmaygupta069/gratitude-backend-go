package middleware

import (
	// "fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tanmaygupta069/api-gateway/internal/constants"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		// Allow all origins, handle special cases with credentials
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(constants.AllowedMethods, ", "))

		// Handle preflight request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

// Helper function to check if the origin is allowed
func isAllowedOrigin(origin string) bool {
	for _, allowedOrigin := range constants.AllowedOrigins {
		if origin == allowedOrigin || origin == "*"{
			return true
		}
	}
	return false
}