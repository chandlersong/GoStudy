package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"testing"
)

type Item struct {
	Year   int
	Title  string
	Plot   string
	Rating float64
}

func TestCreateItem(t *testing.T) {

	svc := CreateLocalDB()
	fmt.Printf("svc is %v \n", svc)

	item := Item{
		Year:   2015,
		Title:  "The Big New Movie",
		Plot:   "Nothing happens at all.",
		Rating: 1.0,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		t.Fail()
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(MovieTableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		t.Fail()
	}
	fmt.Println("create Item save successful ")
}
