package configs

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Service *s3.Client

func S3Config() {

	STORAGE_ENDPOINT := os.Getenv("STORAGE_ENDPOINT")
	STORAGE_REGION := os.Getenv("STORAGE_REGION")
	STORAGE_ACCESS_KEY := os.Getenv("STORAGE_ACCESS_KEY")
	STORAGE_SECRET_KEY := os.Getenv("STORAGE_SECRET_KEY")

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           STORAGE_ENDPOINT,
			SigningRegion: STORAGE_REGION,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(STORAGE_ACCESS_KEY, STORAGE_SECRET_KEY, "")),
	)
	if err != nil {
		fmt.Println("Storage s3", err)
		panic("Failed to initialize storage s3")
	}

	S3Service = s3.NewFromConfig(cfg)
}
