package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"
)

func printKey() string {
	keyData, err := ioutil.ReadFile("rsa/rsa")

	if err != nil {
		return "not find file"
	}

	key, _ := jwt.ParseRSAPrivateKeyFromPEMWithPassword(keyData, "test")

	if err != nil {
		return "not load key"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(key)
	return tokenString
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	time.Sleep(time.Duration(50) * time.Millisecond)
	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc.Identity.CognitoIdentityPoolID)
	tableName := os.Getenv("SOME_VAR")
	fmt.Printf("table name,%v", tableName)
	globalVar := os.Getenv("GLOBAL_VAR")
	log.Printf("gobal var is %s", globalVar)

	notExistsInTemplate := os.Getenv("NOT_EXISTS_IN_TEMPLATE")
	log.Printf("not exists in tempalte can be read %s", notExistsInTemplate)

	notExists := os.Getenv("NOT exists")
	log.Printf("NOT exists: %s", notExists)

	globalParameter := os.Getenv("SomeVar")
	log.Printf("gobal parameter: %s", globalParameter)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", printKey()),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
