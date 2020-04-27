package s3

import (
	"google-rtb/config"
	"google-rtb/pkg/logger"
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	s "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ProcessFileUpload(fileName string) {
	session, err := s.NewSession(&aws.Config{
		Region: aws.String(config.Cfg.AWS.Region),
	})

	if err != nil {
		logger.Error(string(err.Error()))
	}

	err = UploadFileToS3(session, fileName, config.Cfg.FileUploader.LocalStorageDir)

	if err != nil {
		logger.Error(string(err.Error()))
	}
}

func UploadFileToS3(session *s.Session, fileName string, dir string) error {
	file, err := os.Open(fmt.Sprintf("%s/%s", dir, fileName))
	if err != nil {
		return err
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	if err != nil {
		return err
	}

	var size int64 = fileInfo.Size()

	buffer := make([]byte, size)
	file.Read(buffer)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(config.Cfg.AWS.S3.Bucket),
		Key:           aws.String(fileName),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String("application/json"),
	})

	if err == nil {
		filename := fmt.Sprintf("%s/%s", dir, fileName)
		os.Remove(filename)

		logger.Info("local file was uploaded to s3. Filename:  " + filename)
		logger.Info("local file was deleted. Filename: " + filename)
	} else {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		logger.ErrorP("unable to upload file to s3:", params)
	}

	return err
}
