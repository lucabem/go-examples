package S3Utils

import (
	"context"
	"fmt"
	"os"

	AWSUtils "myApp/src/utils/aws"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetObject(bucket string, key string) (*s3.GetObjectOutput, error) {
	sdkConfig, err := AWSUtils.GetSDKConfig()
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
	}
	s3Client := s3.NewFromConfig(sdkConfig)

	fmt.Printf("Let's get key=%s file on bucket=%s\n", bucket, key)

	obj, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	return obj, err
}

func PutObject(fileName string, bucket string, prefix string) error {
	sdkConfig, err := AWSUtils.GetSDKConfig()
	s3Client := s3.NewFromConfig(sdkConfig)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
	} else {
		defer file.Close()
		_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(prefix),
			Body:   file,
		})
		if err != nil {
			fmt.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				fileName, bucket, prefix, err)
		}
	}
	return err

}
