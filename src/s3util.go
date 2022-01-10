// Created by Luis Enrique Fuentes Plata

package src

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"os"
	"path/filepath"
)

// S3PutObjectAPI defines the interface for the PutObject function.
type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

// PutFile uploads a file to an Amazon Simple Storage Service (Amazon S3) bucket
// Inputs:
//     c is the context of the method call, which includes the AWS Region
//     api is the interface that defines the method call
//     input defines the input arguments to the service call.
// Output:
//     If success, a PutObjectOutput object containing the result of the service call and nil
//     Otherwise, nil and an error from the call to PutObject
func PutFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}

// UploadS3BucketFile function used outside off this package, that allows to upload a JSON file stored in tmp
// the tmp dir is created when a GET requests happens.
// Inputs:
//     objectName is a string with the name of the object that will be uploaded.
func UploadS3BucketFile(objectName string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		log.Println("Error", err)
	}

	client := s3.NewFromConfig(cfg)

	file, err := os.Open(localDir + prefix + objectName + extension)
	if err != nil {
		log.Println("Unable to open file" + objectName)
	}
	defer file.Close()

	refPrefix := prefix + objectName + extension

	input := &s3.PutObjectInput{
		Bucket: &awsS3Bucket,
		Key:    &refPrefix,
		Body:   file,
	}

	_, err = PutFile(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got error uploading file:")
		fmt.Println(err)
		return
	}

}

// GetS3BucketFile downloads the content of an object in S3.
// Inputs:
//     key: This is the S3 object key to be downloaded.
// Output:
//     If success return true otherwise false.
func GetS3BucketFile(key string) bool {

	awsS3Key := key

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		log.Println("Error:", err)
	}

	client := s3.NewFromConfig(cfg)
	newManager := manager.NewDownloader(client)

	// this uses concurrency for faster speed through a paginator.
	paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket: &awsS3Bucket,
		Prefix: &awsS3Key,
	})

	var isFileValid = false

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatalln("error:", err)
		}
		if len(page.Contents) > 0 {
			isFileValid = true
			for _, obj := range page.Contents {
				//fmt.Println(aws.ToString(obj.Key))
				if err := downloadToFile(newManager, localDir, awsS3Bucket, aws.ToString(obj.Key)); err != nil {
					log.Fatalln("error:", err)
				}
			}
		}
	}
	return isFileValid
}

// downloadToFile downloads the file locally into tmp directory.
// Inputs:
//     downloader: manager.Downloader pointer
//	   targetDirectory: Where the file will be stored locally.
//	   bucket: AWS S3 bucket.
//     key: This is the S3 object key to be downloaded.
// Output:
//     If success return error if there is an issue.
func downloadToFile(downloader *manager.Downloader, targetDirectory string, bucket string, key string) error {
	// Create the directories in the path
	file := filepath.Join(targetDirectory, key)
	if err := os.MkdirAll(filepath.Dir(file), 0775); err != nil {
		return err
	}

	// Set up the local file
	fd, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	// Download the file
	_, err = downloader.Download(context.TODO(), fd, &s3.GetObjectInput{Bucket: &bucket, Key: &key})

	return err
}
