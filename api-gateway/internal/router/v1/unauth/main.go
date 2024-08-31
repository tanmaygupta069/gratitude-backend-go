package unauth

import (
	"github.com/gin-gonic/gin"
	// "github.com/tanmaygupta069/api-gateway/internal/middleware/auth"
	"github.com/tanmaygupta069/api-gateway/internal/pkg"
)


func RegisterV1UnauthRoutes(rg *gin.RouterGroup){
	unauth:=rg.Group("/unauth")
	{
		postClient,_:=pkg.GetPostServiceClient()
		service := pkg.NewServiceServer(*postClient)
		
		unauth.POST("/createPost",service.CreatePost)
		unauth.GET("/getFeed",service.GetFeed)
		unauth.POST("/getPosts",service.GetPosts)
	}
}