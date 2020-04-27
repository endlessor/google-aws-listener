package fileUploader

import (
	"google-rtb/pkg/svc/aws/s3"
	"google-rtb/pkg/svc/requestsCounter"
	"google-rtb/pkg/util"
	"fmt"
	"time"
)

var fileName = generateFileName()

func GetFileName() string {
	if util.IsChunkFull() {
		requestsCounter.RestartCounter()

		go s3.ProcessFileUpload(fileName)

		fileName = generateFileName()
	}

	return fileName
}

func generateFileName() string {
	return fmt.Sprintf("%d.json", makeTimestamp())
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / 1000 //timestamp in milliseconds
}
