package middleware

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/tanmaygupta069/api-gateway/config"
	"golang.org/x/time/rate"
	"github.com/tanmaygupta069/api-gateway/internal/constants"
)

// RateLimitMiddleware is a Gin middleware function for rate limiting.
func RateLimitMiddleware() gin.HandlerFunc {
	cfg,err:=config.GetConfig()
	if(err!=nil){
		fmt.Println("an error occured in middleware")
	}
	rateLimit := rate.Every(time.Duration(cfg.RateLimit) * time.Second)
	bucketSize := cfg.BucketSize
	l := returnLimiter(rateLimit,bucketSize)
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		limiter := l.AddClient(clientIP)

		if !limiter.Allow() {
			c.JSON(429, gin.H{
				"error": "Too many requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func returnLimiter(ratelimit rate.Limit,bucketSize int) *constants.Limiter{
	return constants.NewLimiter(rate.Limit(ratelimit), bucketSize)
}