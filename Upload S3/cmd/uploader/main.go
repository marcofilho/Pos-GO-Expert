package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
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

	uploadController := make(chan struct{}, 100)
	errorController := make(chan string, 10)

	go func() {
		for {
			select {
			case fileName := <-errorController:
				uploadController <- struct{}{}
				wg.Add(1)
				go uploadFileAWS(fileName, uploadController, errorController)
			}
		}
	}()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		wg.Add(1)
		uploadController <- struct{}{}
		go uploadFileAWS(files[0].Name(), uploadController, errorController)
	}
	wg.Wait()
}

func uploadFileAWS(fileName string, uploadController <-chan struct{}, errorController chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s" + fileName)
	fmt.Printf("Uploading file %s to AWS S3 Bucket %s\n", completeFileName, s3Bucket)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Sprintf("Error opening file %s\n", fileName)
		<-uploadController
		errorController <- completeFileName
		return
	}
	defer file.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s\n", fileName)
		<-uploadController
		errorController <- completeFileName
		return
	}

	fmt.Printf("File %s uploaded successfully\n", fileName)
	<-uploadController
}
