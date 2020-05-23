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

const TableWithTwoKey = "twoKeyTable"

func TestCreateEnv(t *testing.T) {
	t.Run("create db with partition key and sort key", func(t *testing.T) {
		svc := CreateLocalDB()

		hashKey := &dynamodb.KeySchemaElement{
			AttributeName: aws.String("partitionKey"),
			KeyType:       aws.String(dynamodb.KeyTypeHash),
		}
		rangeKey := &dynamodb.KeySchemaElement{
			AttributeName: aws.String("sortKey"),
			KeyType:       aws.String(dynamodb.KeyTypeRange),
		}
		input := &dynamodb.CreateTableInput{
			TableName: aws.String(TableWithTwoKey),
			KeySchema: []*dynamodb.KeySchemaElement{hashKey, rangeKey},
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("partitionKey"),
					AttributeType: aws.String(dynamodb.ScalarAttributeTypeN),
				},
				{
					AttributeName: aws.String("sortKey"),
					AttributeType: aws.String(dynamodb.ScalarAttributeTypeN),
				},
			},
			BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		}
		table, err := svc.CreateTable(input)

		if err != nil {
			t.Log(err.Error())
			t.Fatal("error when create table with partition key and range key")

		}

		t.Logf("create %v success", table.String())
	})
}

func TestPartitionKeyAndSortKeyQuery(t *testing.T) {

	t.Run("get only by sort key", func(t *testing.T) {
		/**
		  Get must base on partition key and sort key
		*/
		svc := CreateLocalDB()
		putItemToTwoKeyTable("1", "1", svc)
		putItemToTwoKeyTable("2", "1", svc)

		input := &dynamodb.GetItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"sortKey":      {N: aws.String("1")},
				"partitionKey": {N: aws.String("1")},
			},
			TableName: aws.String(TableWithTwoKey),
		}
		item, err := svc.GetItem(input)

		if err != nil {
			t.Fatalf("query fail!!!,%v", err)
		}

		t.Logf("item is %v", item.Item)

	})

}

func putItemToTwoKeyTable(pKey string, sKey string, svc *dynamodb.DynamoDB) {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"partitionKey": {
				N: aws.String(pKey),
			},
			"sortKey": {
				N: aws.String(sKey),
			},
		},
		TableName: aws.String(TableWithTwoKey),
	}
	_, _ = svc.PutItem(input)
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
		svc := CreateRemoteDB()

		result, err := svc.Query(&dynamodb.QueryInput{
			TableName:              aws.String(UserTableName),
			IndexName:              aws.String("user_name_index"),
			KeyConditionExpression: aws.String("#NAME = :name"),
			ExpressionAttributeNames: map[string]*string{
				"#NAME": aws.String("name"),
			},
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":name": {
					S: aws.String("chandler"),
				},
			},
			Limit: aws.Int64(1),
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
