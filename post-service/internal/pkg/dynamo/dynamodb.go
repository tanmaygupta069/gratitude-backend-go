package dynamo

import (
	"fmt"
	// "log"
	"sync"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/tanmaygupta069/post-service/config"
)

var dynamoClient *dynamodb.DynamoDB

var cfg,_ = config.GetConfig()

var once sync.Once

type DynamoInterface interface{
	Create(input *dynamodb.PutItemInput)(*dynamodb.PutItemOutput, error)
	Update(input *dynamodb.UpdateItemInput)(*dynamodb.UpdateItemOutput,error)
	Delete(input *dynamodb.DeleteItemInput)(error)
	Query(input *dynamodb.QueryInput)
	GetPosts(input *dynamodb.PutItemInput)
	GetFeed(input *dynamodb.PutItemInput)
}

type DynamoServiceImplementation struct{
}

func NewDynamoClient() *DynamoServiceImplementation {
    return &DynamoServiceImplementation{
	}
}

func GetDynamoClient()*dynamodb.DynamoDB{
	if(dynamoClient==nil){
		InitializeDynamoClient()
	}
	return dynamoClient
}
func InitializeDynamoClient(){
    once.Do(func(){
		sess := session.Must(session.NewSession(&aws.Config{
			Region:   aws.String(cfg.DynamoDBConfig.Region),
			Endpoint: aws.String(fmt.Sprintf("http://%s:%s",cfg.DynamoDBConfig.Host,cfg.DynamoDBConfig.Port)), // Local DynamoDB endpoint
		}))
		dynamoClient=dynamodb.New(sess)	
	})
}

func (d *DynamoServiceImplementation) Create(input *dynamodb.PutItemInput)(*dynamodb.PutItemOutput, error){
	output,err := dynamoClient.PutItem(input)
    if err != nil {
        return nil,err
    }
	return output,nil
}

func (d *DynamoServiceImplementation) Update(input *dynamodb.UpdateItemInput)(*dynamodb.UpdateItemOutput, error){
	output, err := dynamoClient.UpdateItem(input)
    if err != nil {
        return nil, err
    }
	return output, nil
}

func (d *DynamoServiceImplementation) Delete(input *dynamodb.DeleteItemInput)(error){
	_,err := dynamoClient.DeleteItem(input)
    if err != nil {
        return err
    }
	return nil
}

func (d *DynamoServiceImplementation) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
    res,err:=dynamoClient.Query(input)
	if err!=nil{
		return nil,err
	}
	return res,nil
}


// Implement methods for CreatePost, UpdatePost, DeletePost, GetPosts, and GetFeed
