package api

import (
	"github.com/gin-gonic/gin"
	"google-rtb/model"
	"google-rtb/pkg/logger"
	"google-rtb/pkg/svc/streamer"
	"net/http"
)

func RtbListener(c *gin.Context) {
	var requestBody model.RequestBody

	err := c.BindJSON(&requestBody)

	if err != nil {
		params := &logger.LogParams{}
		params.Add("reason:", err)
		params.Add("requestBody:", requestBody)
		logger.ErrorP("unable to parse requestBody:", params)

		return
	}

	go streamer.ProcessRequestBody(requestBody)

	c.JSON(http.StatusOK, requestBody)
}
