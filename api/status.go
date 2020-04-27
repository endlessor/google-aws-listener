package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
