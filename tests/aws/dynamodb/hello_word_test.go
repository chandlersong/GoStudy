package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"testing"
)

/*
how to create a dynamodb client:
and you should create a credentials by the reference
https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
*/
func TestCreateClient(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
	})

	if err != nil {
		fmt.Printf("The returned error is %s.\n", err)
	}
	fmt.Printf("sess is %v \n", sess)
	svc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	fmt.Printf("svc is %v \n", svc)
}
