package internal

import (
	"context"
	"fmt"
	"log"

	pb "github.com/tanmaygupta069/post-service/generated"
)

type ControllerInterface interface {
	CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error)
	UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error)
	DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeleteResponse, error)
	GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error)
	GetFeed(ctx context.Context, req *pb.GetFeedRequest) (*pb.GetFeedResponse, error)
	pb.UnimplementedPostServiceServer
}

type ControllerImplementation struct {
	postService ServiceInterface
	pb.UnimplementedPostServiceServer
}

func NewControllerServer(postService ServiceInterface)*ControllerImplementation{
	return &ControllerImplementation{
		postService:postService,
	}
}

func (s *ControllerImplementation) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	post, err := s.postService.CreatePost(&PostRequest{
		UserId:      req.UserId,
		Content:     req.Content,
		IsAnonoymus: req.IsAnonymous,
	})

	if err != nil {
		log.Printf("Failed to create post: %v", err)
		return nil, err
	}

	Post := &pb.Post{
		PostId:      post.PostID,
		UserId:      post.UserID,
		Content:     post.Content,
		IsAnonymous: post.IsAnonymous,
		Timestamp:   post.CreatedAt,
	}
	response := &pb.PostResponse{Post: Post}
	return response, nil
}

func (s *ControllerImplementation) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.PostResponse, error) {
	post, err := s.postService.UpdatePost(&PostRequest{
		UserId:      req.UserId,
		PostId: 	 req.PostId,
		Content:     req.Content,
		IsAnonoymus: req.IsAnonymous,
	})
	fmt.Printf("in controller post service")
	if err != nil {
		log.Printf("Failed to create post: %v", err)
		return nil, err
	}

	Post := &pb.Post{
		PostId:      post.PostID,
		UserId:      post.UserID,
		Content:     post.Content,
		IsAnonymous: post.IsAnonymous,
		Timestamp:   post.CreatedAt,
	}
	response := &pb.PostResponse{Post: Post}
	return response, nil
}

func (s *ControllerImplementation) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeleteResponse, error) {
	err := s.postService.DeletePost(&PostRequest{
		UserId:      req.UserId,
		PostId: 	 req.PostId,
	})
    if err != nil {
        log.Printf("Failed to delete post: %v", err)
        return &pb.DeleteResponse{Success: false}, err
    }
    return &pb.DeleteResponse{Success: true}, nil
}

func (s *ControllerImplementation) GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	posts, lastEvaluatedKey, err := s.postService.GetPosts(&GetPostsRequest{
		UserIdWhose: req.UserIdWhose,
		UserIdWhom: req.UserIdWhom,
		PageSize: int(req.PageSize),
		LastEvaluatedKey: req.LastEvaluatedKey,
	})
    if err != nil {
        log.Printf("Failed to get posts: %v", err)
        return nil, err
    }

    // Convert internal Post model to gRPC Post model
    var grpcPosts []*pb.Post
    for _, post := range posts {
        grpcPosts = append(grpcPosts, &pb.Post{
            PostId:      post.PostID,
            UserId:      post.UserID,
            Content:     post.Content,
            IsAnonymous: post.IsAnonymous,
            Timestamp:   post.CreatedAt,
        })
    }

    response := &pb.GetPostsResponse{
        Posts:            grpcPosts,
        LastEvaluatedKey: lastEvaluatedKey,
    }
 
    return response, nil
}

func (s *ControllerImplementation) GetFeed(ctx context.Context, req *pb.GetFeedRequest) (*pb.GetFeedResponse, error) {
    posts, lastEvaluatedKey, err := s.postService.GetFeed(&GetPostsRequest{
		UserIdWhom: req.UserId,
		PageSize:  int(req.PageSize),
		LastEvaluatedKey: req.LastEvaluatedKey,
	})

    if err != nil {
        log.Printf("Failed to get feed: %v", err)
        return nil, err
    }

    // Convert internal Post model to gRPC Post model
    var grpcPosts []*pb.Post
    for _, post := range posts {
        grpcPosts = append(grpcPosts, &pb.Post{
            PostId:      post.PostID,
            UserId:      post.UserID,
            Content:     post.Content,
            IsAnonymous: post.IsAnonymous,
            Timestamp:   post.CreatedAt,
        })
    }

    response := &pb.GetFeedResponse{
        Posts:            grpcPosts,
        LastEvaluatedKey: lastEvaluatedKey,
    }

    return response, nil
}