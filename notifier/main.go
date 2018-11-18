package main

import (
    "fmt"
    "context"
    "github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context) (error) {
    fmt.Println(ctx)

    return nil
}

func main() {
  lambda.Start(Handler)
}
