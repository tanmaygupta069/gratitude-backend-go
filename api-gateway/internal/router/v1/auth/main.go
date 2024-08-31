package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tanmaygupta069/api-gateway/internal/middleware/auth"
	"github.com/tanmaygupta069/api-gateway/internal/pkg"
)


func RegisterV1AuthRoutes(rg *gin.RouterGroup){
	auth:=rg.Group("/auth")
	{
		auth.Use(auth_test.GenericAccess()) //for IDOR requests

		postClient,_:=pkg.GetPostServiceClient()
		service := pkg.NewServiceServer(*postClient)

		auth.POST("/updatePost",service.UpdatePost)
		auth.DELETE("/deletePost",service.DeletePost)
	}

}