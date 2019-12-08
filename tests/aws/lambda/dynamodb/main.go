package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	Year   int     `json:"Year"`
	Title  string  `json:"Title"`
	Plot   string  `json:"Plot"`
	Rating float64 `json:"Rating"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Body size = %d.\n", len(request.Body))
	fmt.Printf("Body content = %v.\n", request.Body)

	var item Item
	json.Unmarshal([]byte(request.Body), &item)
	response := "request body is : #{item}"
	fmt.Printf(response)
	fmt.Printf("body content: %v \n", request.Body)
	fmt.Printf("item is : %v \n", item)
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Music"),
	}

	svc := createDBClient()
	_, errInput := svc.PutItem(input)
	if errInput != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(errInput.Error())
	}

	return events.APIGatewayProxyResponse{Body: string(item.Year), StatusCode: 200}, nil
}

func createDBClient() *dynamodb.DynamoDB {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
	})

	if err != nil {
		fmt.Printf("The returned error is %s.\n", err)
	}
	fmt.Printf("sess is %v \n", sess)
	return dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
}

func main() {
	lambda.Start(HandleRequest)
}
