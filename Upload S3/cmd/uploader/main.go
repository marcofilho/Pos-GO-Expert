package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	session, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-west-1"),
			Credentials: credentials.NewStaticCredentials(
				"AKID",
				"SECRET",
				"",
			),
		})
	if err != nil {
		log.Fatal(err)
	}
	s3Client = s3.New(session)
	s3Bucket = "Pos-Go-Expert"

}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		uploadFileAWS(files[0].Name())
	}
}

func uploadFileAWS(fileName string) {
	completeFileName := fmt.Sprintf("./tmp/%s" + fileName)
	fmt.Printf("Uploading file %s to AWS S3 Bucket %s\n", completeFileName, s3Bucket)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Sprintf("Error opening file %s\n", fileName)
		return
	}
	defer file.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
	}
}
