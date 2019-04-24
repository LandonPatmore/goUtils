package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	session2 "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Creates a dynamo client
func CreateDynamoClient() (*dynamodb.DynamoDB, error) {
	rKey := ReadAWSEnv()

	session, err := session2.NewSession(&aws.Config{
		Region: aws.String(rKey),
		// Credentials aren't here because we pass in ENV variables and the sdk auto detects them
	})

	if err != nil {
		return nil, err
	}

	// Create DynamoDB client
	client := dynamodb.New(session)

	return client, nil
}

// Takes items in, marshals them, and then sends them to the database
func PutItem(dynamoClient *dynamodb.DynamoDB, table string, v interface{}) error {
	av, err := dynamodbattribute.MarshalMap(v)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}

	_, err = dynamoClient.PutItem(input)

	if err != nil {
		return err
	}

	fmt.Print(v)
	return nil
}

func GetAllItems(dynamoClient *dynamodb.DynamoDB, table string, v interface{}) (err error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(table),
	}

	result, err := dynamoClient.Scan(params)

	if err != nil {
		return nil
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &v)

	if err != nil {
		return err
	}

	return nil
}

func GetRowCount(dynamoClient *dynamodb.DynamoDB, table string) (count int64, err error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(table),
	}

	result, err := dynamoClient.Scan(params)

	if err != nil {
		return 0, nil
	}

	return *result.Count, nil
}
