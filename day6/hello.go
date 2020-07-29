package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
)

func main() {
	lambda.Start(hello)
}


type MyEvent struct {
	Name string `json:"uname"`
	Age  int    `json:"uage"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func hello(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

