package pkg

import (
	"context"
	"fmt"
	// "log"
	"net/http"
	"github.com/gin-gonic/gin"
	// "github.com/tanmaygupta069/api-gateway/config"
	postpb "github.com/tanmaygupta069/post-service/generated"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
)

var RequestData struct {
	Content   string `json:"content"`
	UserIDWhom string `json:"userIdWhom"`
	UserIdWhose string `json:"userIdWhose"`
	Timestamp int64  `json:"timestamp"`
	IsAnonymous bool `json:"isAnonymous"`
	PageSize int32 `json:"pageSize"`
	LastEvaluatedKey string `json:"lastEvaluatedKey"`
	PostId string `json:"postId"`
}

type ServiceInterface interface {
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
	CreatePost(c *gin.Context)
	GetPosts(c *gin.Context)
	GetFeed(c *gin.Context)
}


type ServiceImplementation struct {
	postClient postpb.PostServiceClient
}

func NewServiceServer(postClient postpb.PostServiceClient)*ServiceImplementation{
	return &ServiceImplementation{
		postClient: postClient,
	}
}

func (s ServiceImplementation) UpdatePost(c* gin.Context){

	if err := c.BindJSON(&RequestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	postRequest := &postpb.UpdatePostRequest{
		Content:   RequestData.Content,
		PostId:    RequestData.PostId,
		IsAnonymous: RequestData.IsAnonymous,
		UserId: RequestData.UserIDWhom,
	}

	// Call the Post gRPC service method
	resp, err := s.postClient.UpdatePost(context.Background(), postRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create post",err.Error())})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{"post_id": resp.Post.PostId})
}

func (s ServiceImplementation) DeletePost(c* gin.Context){
	if err := c.BindJSON(&RequestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	deletRequest := &postpb.DeletePostRequest{
		PostId:    RequestData.PostId,
		UserId: RequestData.UserIDWhom,
	}

	// Call the Post gRPC service method
	resp, err := s.postClient.DeletePost(context.Background(), deletRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create post",err.Error())})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{"post_id": resp.Success})
}

func (s ServiceImplementation) CreatePost(c* gin.Context){
	if err := c.BindJSON(&RequestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	createRequest := &postpb.CreatePostRequest{
		UserId: RequestData.UserIdWhose,
		Content: RequestData.Content,
		IsAnonymous: RequestData.IsAnonymous,
	}

	// Call the Post gRPC service method
	resp, err := s.postClient.CreatePost(context.Background(), createRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create post",err.Error())})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{"post_id": resp.Post})
}

func (s ServiceImplementation) GetPosts(c* gin.Context){
	if err := c.BindJSON(&RequestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	getRequest := &postpb.GetPostsRequest{
		UserIdWhom: RequestData.UserIDWhom,
		UserIdWhose: RequestData.UserIdWhose,
		PageSize: RequestData.PageSize,
		LastEvaluatedKey: RequestData.LastEvaluatedKey,	
	}

	// Call the Post gRPC service method
	resp, err := s.postClient.GetPosts(context.Background(), getRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create post",err.Error())})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{"last evaluated key": resp.Posts})
}

func (s ServiceImplementation) GetFeed(c* gin.Context){
	if err := c.BindJSON(&RequestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	getFeedRequest := &postpb.GetFeedRequest{
		UserId: RequestData.UserIDWhom,
		PageSize: RequestData.PageSize,
		LastEvaluatedKey: RequestData.LastEvaluatedKey,	
	}

	// Call the Post gRPC service method
	resp, err := s.postClient.GetFeed(context.Background(), getFeedRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create post",err.Error())})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{"last evaluated key": resp.Posts})
}


