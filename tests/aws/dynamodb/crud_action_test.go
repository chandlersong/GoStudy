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

type User struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	Platform      string `json:"platform"`
	CurrentDevice string `json:"currentDevice"`
}

func TestDynamodbCrud(t *testing.T) {
	t.Run("save items to db", func(t *testing.T) {

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
	})

	t.Run("query by index", func(t *testing.T) {
		svc := CreateLocalDB()

		tableName := "User"

		result, err := svc.Query(&dynamodb.QueryInput{
			TableName: aws.String(tableName),
			IndexName: aws.String("user_name_index"),
			KeyConditions: map[string]*dynamodb.Condition{
				"name": {
					ComparisonOperator: aws.String("EQ"),
					AttributeValueList: []*dynamodb.AttributeValue{
						{
							S: aws.String("chandler"),
						},
					},
				}},
		})

		if err != nil {
			t.Fatal(err.Error())
		}

		for a, i := range result.Items {
			t.Logf("first item is %v", a)

			user := User{}

			err = dynamodbattribute.UnmarshalMap(i, &user)
			t.Logf("user is %v", user)

		}
	})

}
