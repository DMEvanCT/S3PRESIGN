package S3Presign

import "time"

type S3BucketInfo struct {
	// Name of the bucket to retreive items from
	BucketName string `json:"BucketName"`
	// Name of item from the bucket
	ItemName string `json:"ItemName"`
	// Region the bucket exists in
	BucketRegion string `json:"BucketRegion"`
	// Validity of the token in minutes.
	Validity time.Duration

}

type S3TempUrlReturn struct {
	TempS3URL string `json:"TempS3URL"`
}