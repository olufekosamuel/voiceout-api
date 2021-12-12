package helpers

import (
	"log"
	"math/rand"
	"mime/multipart"
	"os"

	b64 "encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

func ConnectAws() *session.Session {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				accessKey,
				secretKey,
				"",
			),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return sess

}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func UploadFile(file *multipart.FileHeader, directory string) (string, error) {
	sess := ConnectAws()

	uploader := s3manager.NewUploader(sess)

	myBucket := os.Getenv("BUCKET_NAME")

	dec := b64.StdEncoding.EncodeToString([]byte(file.Filename))

	filename := ""
	if len(directory) != 0 {
		filename = directory + "/" + dec + RandStringBytes(5)
	} else {
		filename = dec + RandStringBytes(5)
	}

	actualFile, err := file.Open()

	if err != nil {
		return "", err
	}

	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   actualFile,
	})

	if err != nil {
		return "", err
	}

	return up.Location, nil
}
