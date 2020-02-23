package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"log"
	"os"
	"time"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	time.Sleep(time.Duration(50) * time.Millisecond)
	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc.Identity.CognitoIdentityPoolID)
	tableName := os.Getenv("SOME_VAR")

	globalVar := os.Getenv("GLOBAL_VAR")
	log.Printf("gobal var is %s", globalVar)

	notExistsInTemplate := os.Getenv("NOT_EXISTS_IN_TEMPLATE")
	log.Printf("not exists in tempalte can be read %s", notExistsInTemplate)

	notExists := os.Getenv("NOT exists")
	log.Printf("NOT exists: %s", notExists)

	globalParameter := os.Getenv("SomeVar")
	log.Printf("gobal parameter: %s", globalParameter)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", tableName),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
