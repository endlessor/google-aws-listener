package streamer

import (
	"encoding/json"
	"fmt"
	"google-rtb/config"
	"google-rtb/model"
	"google-rtb/pkg/logger"
	"google-rtb/pkg/svc/fileUploader"
	"google-rtb/pkg/svc/requestsCounter"
	"os"
)

func ProcessRequestBody(requestBody model.RequestBody) {
	filename := fileUploader.GetFileName()

	requestsCounter.Increment()

	fullFileName := fmt.Sprintf("%s/%s", config.Cfg.FileUploader.LocalStorageDir, filename)

	f, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		params.Add("fileName:", fullFileName)
		logger.ErrorP("unable to open file:", params)

		return
	}

	defer f.Close()

	jsonContent, err := json.Marshal(requestBody)

	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		params.Add("requestBody:", requestBody)
		logger.ErrorP("unable to parse requestBody:", params)

		return
	}

	if _, err = f.WriteString(fmt.Sprintf("%s\n", string(jsonContent))); err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		params.Add("fileName:", fullFileName)
		params.Add("content:", jsonContent)
		logger.ErrorP("unable to write to file:", params)

		return
	}
}
