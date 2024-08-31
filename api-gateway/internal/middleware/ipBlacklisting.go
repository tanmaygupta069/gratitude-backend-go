package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tanmaygupta069/api-gateway/internal/constants"
)

func IPBlacklistingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		for _, blockedIP := range constants.BlackListedIps {
			if strings.TrimSpace(clientIP) == blockedIP {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"message": "Your IP is blacklisted.",
				})
				return
			}
		}
		fmt.Println(clientIP)
		c.Next()
	}
}
