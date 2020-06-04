package main

import (
	"github.com/ARepublik/S3URLGENERATOR/S3Presign"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(S3Presign.Handler)
}