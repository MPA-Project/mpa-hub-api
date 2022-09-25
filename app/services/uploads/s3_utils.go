package uploads

import (
	"bytes"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	aws_types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gabriel-vasile/mimetype"
	"myponyasia.com/hub-api/pkg/configs"
)

func UploadS3(file []byte, filepath string, metadata map[string]string) error {
	mtype := mimetype.Detect(file)
	var acl aws_types.ObjectCannedACL = "public-read"
	STORAGE_BUCKET := os.Getenv("STORAGE_BUCKET")
	if _, err := configs.S3Service.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:       aws.String(STORAGE_BUCKET),
		Body:         bytes.NewReader(file),
		Key:          aws.String(filepath),
		ContentType:  aws.String(mtype.String()),
		Metadata:     metadata,
		CacheControl: aws.String("public, no-transform, immutable, max-age=2592000"),
		ACL:          acl,
	}); err != nil {
		return err
	}

	return nil
}

func DeleteS3(filepath string) error {
	if _, err := configs.S3Service.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("STORAGE_BUCKET")),
		Key:    aws.String(filepath),
	}); err != nil {
		return err
	}

	return nil
}
