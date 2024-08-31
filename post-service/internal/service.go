package internal

import (
	// "context"
	// "log"
	// "net"
    "time"
    "github.com/google/uuid"
)

type PostRequest struct{
    PostId string
    UserId string
    Content string
    IsAnonoymus bool
    CreatedAt string
}

type GetPostsRequest struct{
    UserIdWhose string
    UserIdWhom string
    PageSize int
    LastEvaluatedKey string
}

type ServiceInterface interface {
    CreatePost(req *PostRequest)(*Post,error)
	UpdatePost(req *PostRequest)(*Post,error)
	DeletePost(req *PostRequest)(error)
	GetPosts(req *GetPostsRequest)([]*Post,string, error)
	GetFeed(req *GetPostsRequest)([]*Post,string, error)
    GeneratePostId()string
}

type ServiceImplementation struct{
    repo PostRepositoryInterface
    post Post
}

func NewPostServiceServer(repo PostRepositoryInterface) *ServiceImplementation {
    return &ServiceImplementation{
        repo:repo,
    }
}

func (s *ServiceImplementation) CreatePost(req *PostRequest) (*Post,error) {

    post:=&Post{
        UserID : req.UserId,
        PostID : s.GeneratePostId(),
        Content : req.Content,
        IsAnonymous : req.IsAnonoymus,
        CreatedAt : time.Now().Format(time.RFC3339),
    }

    _,err := s.repo.CreatePost(post)
    if err != nil {
        return nil, err
    }

    return post, nil
}

func (s *ServiceImplementation) UpdatePost(req *PostRequest) (*Post,error) {
    post:=&Post{
        UserID : req.UserId,
        PostID : req.PostId,
        Content : req.Content,
        IsAnonymous : req.IsAnonoymus,
        CreatedAt : time.Now().Format(time.RFC3339),
    }

    _,err := s.repo.UpdatePost(post)
    if err != nil {
        return nil, err
    }

    return post, nil
}

func (s *ServiceImplementation) DeletePost(req *PostRequest) error {
    post:=&Post{
        UserID : req.UserId,
        PostID : req.PostId,
    }
    err:=s.repo.DeletePost(post)
    if err!=nil{
        return err
    }
    return nil
}

func (s *ServiceImplementation) GetPosts(req *GetPostsRequest) ([]*Post, string, error) {
    post:=&GetPosts{
        PageSize: req.PageSize,
        UserID: req.UserIdWhose,
        LastEvaluatedKey: req.LastEvaluatedKey,
    }
    return s.repo.GetPosts(post)
}

func (s *ServiceImplementation) GetFeed(req *GetPostsRequest) ([]*Post, string, error) {
    post:=&GetPosts{
        PageSize: req.PageSize,
        UserID: req.UserIdWhom,
        LastEvaluatedKey: req.LastEvaluatedKey,
    }
    return s.repo.GetFeed(post)
}



func (s *ServiceImplementation) GeneratePostId()string{
    return uuid.New().String()
}


