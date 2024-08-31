package internal

import (
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/tanmaygupta069/post-service/internal/pkg/dynamo"
	"github.com/aws/aws-sdk-go/aws"
)

type PostRepositoryInterface interface {
    CreatePost(post *Post) (*dynamodb.PutItemOutput,error)
	UpdatePost(post *Post) (*dynamodb.UpdateItemOutput,error)
	DeletePost(post *Post) (error)
    GetPosts(post *GetPosts) ([]*Post,string, error)
    GetFeed(post *GetPosts)([]*Post,string, error)
    // Other CRUD methods here
}

type DynamoPostRepositoryImplementation struct {
	db dynamo.DynamoServiceImplementation
    tableName string
}

func NewDynamoDBPostRepository(db dynamo.DynamoServiceImplementation, tableName string) *DynamoPostRepositoryImplementation {
    return &DynamoPostRepositoryImplementation{
        db: db,
        tableName: tableName,
    }
}

func (d *DynamoPostRepositoryImplementation) CreatePost(post *Post)(*dynamodb.PutItemOutput,error) {
    item, err := dynamodbattribute.MarshalMap(post)
    if err != nil {
        return nil,err
    }

    input := &dynamodb.PutItemInput{
        TableName: &d.tableName,
        Item:      item,
	}
	output,err:=d.db.Create(input)
	if err!=nil{
		return nil,err
	}
    return output,nil
}

func (d *DynamoPostRepositoryImplementation) UpdatePost(post *Post)(*dynamodb.UpdateItemOutput,error) {
	key, err := dynamodbattribute.MarshalMap(map[string]string{
        "PostID": post.PostID,
    })
    if err != nil {
        return nil, err
	}

	updateValues, err := dynamodbattribute.MarshalMap(map[string]interface{}{
        ":content":     post.Content,
        ":isAnonymous": post.IsAnonymous,
        ":timestamp":   post.CreatedAt,
    })
	if err != nil {
        return nil, err
    }

    input := &dynamodb.UpdateItemInput{
        TableName:                 &d.tableName,
        Key:                       key,
        UpdateExpression:          aws.String("SET Content = :content, IsAnonymous = :isAnonymous, CreatedAt = :timestamp"),
        ExpressionAttributeValues: updateValues,
        ReturnValues:              aws.String("ALL_NEW"),
    }

	output,err:=d.db.Update(input)
	if err!=nil{
		return nil,err
	}

    return output,nil
}

func (d *DynamoPostRepositoryImplementation) DeletePost(post *Post) error {
	
    key, err := dynamodbattribute.MarshalMap(map[string]string{
        "PostID": post.PostID,  // Assuming "PostID" is the primary key
        "UserID": post.UserID,  // Assuming "UserID" is part of the primary key or used for authorization
    })
    if err != nil {
        return err
    }

    // Creating the DeleteItemInput
    input := &dynamodb.DeleteItemInput{
        TableName: &d.tableName,
        Key:       key,
    }

	err = d.db.Delete(input)
		if err!=nil{
			return err
		}

    return nil
}

func (d *DynamoPostRepositoryImplementation) GetPosts(post *GetPosts)  ([]*Post, string, error){
    input := &dynamodb.QueryInput{
        TableName:              aws.String(d.tableName),
        IndexName:              aws.String("UserId-Timestamp-index"), // Assuming this is the name of your GSI
        KeyConditionExpression: aws.String("UserId = :userId"),
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":userId": {
                S: aws.String(post.UserID),
            },
        },
        Limit: aws.Int64(int64(post.PageSize)),
        ScanIndexForward:  aws.Bool(false),  // To get the most recent posts first
    }

    key:="LastEvaluatedKey"
    attributeValueMap := map[string]*dynamodb.AttributeValue{
        key: {
            S: aws.String(post.LastEvaluatedKey),
        },
    }
    input.ExclusiveStartKey=attributeValueMap

    // Perform the query
    result, err := d.db.Query(input)
    if err != nil {
        return nil,"",err
    }
    var posts []*Post

    err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &posts)
    if err != nil {
        return nil,"",err
    }

    var value string
    if attributeValue, ok := result.LastEvaluatedKey[key]; ok {
        if attributeValue.S != nil {
            value = *attributeValue.S
        }
    }
    // Return LastEvaluatedKey for pagination
    return posts, value, nil
}

func (d *DynamoPostRepositoryImplementation) GetFeed(post *GetPosts) ([]*Post, string, error) {
    input := &dynamodb.QueryInput{
        TableName:              aws.String(d.tableName),
        IndexName:              aws.String("UserId-Timestamp-index"), // Replace with your GSI name
        KeyConditionExpression: aws.String("UserId = :userId"),
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":userId": {
                S: aws.String(post.UserID),
            },
        },
        Limit:            aws.Int64(int64(post.PageSize)),
        ScanIndexForward: aws.Bool(false), // Order by Timestamp descending
    }

    key:="LastEvaluatedKey"
    attributeValueMap := map[string]*dynamodb.AttributeValue{
        key: {
            S: aws.String(post.LastEvaluatedKey),
        },
    }
    input.ExclusiveStartKey=attributeValueMap

    // Perform the query
    result, err := d.db.Query(input)
    if err != nil {
        return nil,"", err
    }

    var posts []*Post
    err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &posts)
    if err != nil {
        return nil, "", err
    }

    var value string
    if attributeValue, ok := result.LastEvaluatedKey[key]; ok {
        if attributeValue.S != nil {
            value = *attributeValue.S
        }
    }

    return posts, value, nil
}
