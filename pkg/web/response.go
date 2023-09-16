package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, response{Data: data})
}

func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := errorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}

	Response(c, status, err)
}
