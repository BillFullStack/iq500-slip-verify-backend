package utils

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpStatusCode int, status int, message string, errorText string, data interface{}) {
	rsp := gin.H{"code": status, "message": message, "error": errorText, "payload": data}
	c.JSON(httpStatusCode, rsp)
}
