package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/bwmarrin/snowflake"
	"reflect"
	"strings"
	"testing"
)

type Item struct {
	Year   int
	Title  string
	Plot   string
	Rating float64
}

type User struct {
	Id            int64  `json:"id"`
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

	/***
	put check non-key attribute can't work
	for more detail please check https://github.com/aws/aws-sdk-js/issues/1611
	*/
	t.Run("save with condition non-key attribute", func(t *testing.T) {
		svc := CreateLocalDB()
		fmt.Printf("svc is %v \n", svc)

		user := User{
			Id:            6311018584275884701,
			Name:          "chandlerZ",
			Password:      "passowrd",
			Platform:      "bf",
			CurrentDevice: "bfbf",
		}

		node, _ := snowflake.NewNode(1)
		user.Id = node.Generate().Int64()
		av, _ := dynamodbattribute.MarshalMap(user)

		input := &dynamodb.PutItemInput{
			Item:                av,
			TableName:           aws.String(UserTableName),
			ConditionExpression: aws.String(" attribute_exists(#N) and #N <> :name"),
			ExpressionAttributeNames: map[string]*string{
				"#N": aws.String("name"),
			},
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":name": {
					S: aws.String(user.Name),
				},
			},
		}

		_, err := svc.PutItem(input)
		if err != nil {
			fmt.Println("Got error calling PutItem:")
			s := err.Error()
			fmt.Println(s)
			errType := reflect.TypeOf(err)
			fmt.Printf("errorType is %v \n", errType)
			exception := dynamodb.ErrCodeConditionalCheckFailedException

			if !strings.HasPrefix(err.Error(), exception) {
				t.Fail()
			}
			return
		}
		t.Fail()

	})

	t.Run("query by index", func(t *testing.T) {
		svc := CreateLocalDB()

		result, err := svc.Query(&dynamodb.QueryInput{
			TableName: aws.String(UserTableName),
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

			user := new(User)

			err = dynamodbattribute.UnmarshalMap(i, &user)
			t.Logf("user is %v", user)

		}
	})

}
