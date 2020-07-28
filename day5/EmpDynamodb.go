package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// List/Retrieve

//Create / Insert

func getdynamodb() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")})

	if err != nil {
		fmt.Println("In err", err)
		return nil
	} else {
		fmt.Println("Session Created Successfully")
		fmt.Println(sess)
		return dynamodb.New(sess)
	}

}

func Save(emp EmpStruct) {
	svc := getdynamodb()
	fmt.Println("emp ", emp)
	emp1, err := dynamodbattribute.MarshalMap(emp)
	fmt.Println("emp1 ", emp1)
	if err != nil {
		fmt.Println("Got error marshalling")
		fmt.Println(err.Error())
	} else {
		tableName := "emp"
		input := &dynamodb.PutItemInput{
			Item:      emp1,
			TableName: aws.String(tableName),
		}

		_, err = svc.PutItem(input)
		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
		} else {
			fmt.Println("Successfully added ", emp)
		}
	}

}

/*func list() []EmpStruct {
	return
}
*/
