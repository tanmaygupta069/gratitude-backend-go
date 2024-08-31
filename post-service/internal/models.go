package internal

type Post struct {
    PostID      string
    UserID      string `json:"userId"`
    Content     string `json:"content"`
    IsAnonymous bool   `json:"isAnonymous"`
    CreatedAt   string
}

type GetPosts struct{
    PageSize int
    UserID string
    LastEvaluatedKey string
}
