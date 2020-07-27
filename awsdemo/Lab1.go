package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")})

	if err != nil {
		fmt.Println("In err", err)
	} else {
		fmt.Println("Session Created Successfully")
		fmt.Println(sess)
		svc := dynamodb.New(sess)
		listtablesoutput, err := svc.ListTables(&dynamodb.ListTablesInput{})
		if err != nil {
			fmt.Println("In error from ListTables", err)
		} else {
			fmt.Println("List of Tables Created ")
			fmt.Println(listtablesoutput)
		}
	}

}
