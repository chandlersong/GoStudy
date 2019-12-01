package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
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

func TestCreateLocalClient(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-1"),
		Endpoint: aws.String("http://cloudTest:18080"),
	})

	if err != nil {
		fmt.Printf("The returned error is %s.\n", err)
	}
	fmt.Printf("sess is %v \n", sess)
	svc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
	fmt.Printf("svc is %v \n", svc)
}

func TestCreateTable(t *testing.T) {
	svc := CreateLocalDB()
	fmt.Printf("svc is %v \n", svc)
	tableName := "Movies"

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Year"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Title"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Year"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Title"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, errCreateTable := svc.CreateTable(input)
	if errCreateTable != nil {
		fmt.Printf("Got error calling CreateTable: %s.\n", errCreateTable)
		fmt.Println(errCreateTable)
	}

	fmt.Println("Created the table", tableName)
}

func TestListTable(t *testing.T) {
	svc := CreateLocalDB()
	fmt.Printf("svc is %v \n", svc)
	fmt.Printf("Tables:\n")
	input := &dynamodb.ListTablesInput{}

	for {
		// Get the list of tables
		result, err := svc.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		for _, n := range result.TableNames {
			fmt.Println(*n)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}
	}
}
