package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const MovieTableName = "Movies"

func CreateLocalDB() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-1"),
		Endpoint: aws.String("http://localhost:18080"),
	})
	if err != nil {
		fmt.Printf("The returned error is %s.\n", err)
	}
	svc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	return svc
}
