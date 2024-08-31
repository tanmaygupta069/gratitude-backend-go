package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tanmaygupta069/api-gateway/internal/middleware"
	"github.com/tanmaygupta069/api-gateway/internal/router/v1/auth"
	"github.com/tanmaygupta069/api-gateway/internal/router/v1/unauth"

)

func GetRouter() *gin.Engine{
	router := gin.Default()

	router.Use(middleware.RateLimitMiddleware())

	router.Use(middleware.IPBlacklistingMiddleware())

	router.Use(middleware.CORSMiddleware())

	router.GET("/health", func(c *gin.Context) {
        c.String(200, "Hello, HTTPS!")
    })

	v1:=router.Group("/v1")
	{
		auth.RegisterV1AuthRoutes(v1)
		unauth.RegisterV1UnauthRoutes(v1)
	}

	return router
}