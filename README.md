# S3Presign

#### Used to create presigned URL from S3
This application is a serverless lambda function that runs on
aws. The api endpoint is a get request. 

Development
```bash 
REDACTED
```

####Requires the following to be sent via JSON
```json
{
"BucketName": "examplebucketname",
"ItemName": "example.jpg",
"BucketRegion": "us-east-1",
"Validity": "10"
}
```

Note: Validity is measured in minutes. 

##Build:
```bash 
make 
serverless deploy -v 
```
