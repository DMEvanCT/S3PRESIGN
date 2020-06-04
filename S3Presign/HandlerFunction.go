package S3Presign

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"time"
)

func Handler(request events.APIGatewayProxyRequest )  ( events.APIGatewayProxyResponse, error) {
	var BucketInfo S3BucketInfo

		err := json.Unmarshal([]byte(request.Body), &BucketInfo)
		if err != nil {
			ApiResponse := events.APIGatewayProxyResponse{Body: BucketInfo.BucketName, StatusCode: 502}
			return ApiResponse, err
		}

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(BucketInfo.BucketRegion)},
		)

		s3svc := s3.New(sess)

		req, _ := s3svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(BucketInfo.BucketName),
			Key:    aws.String(BucketInfo.ItemName),
		})

		urltoreturn, err := req.Presign(BucketInfo.Validity * time.Minute)

		if err != nil {
			ApiResponse := events.APIGatewayProxyResponse{Body: BucketInfo.BucketName, StatusCode: 502}
			return ApiResponse, err

			log.Println("Failed to sign request", err)
		}

		log.Println("The URL is", urltoreturn)
		returnurl := S3TempUrlReturn{
			TempS3URL: urltoreturn,
		}



		rbody, err := json.Marshal(returnurl)
		if err != nil {
			log.Println("Error")
		}

		ApiResponse := events.APIGatewayProxyResponse{Body: string(rbody), StatusCode: 200}

		return ApiResponse, nil
	}


