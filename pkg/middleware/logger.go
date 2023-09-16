package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	basePath = "http://localhost:8080"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL
		time := time.Now()
		method := ctx.Request.Method

		ctx.Next()

		var size int

		if ctx.Writer != nil {
			size = ctx.Writer.Size()
		}

		fmt.Printf("Path: %s%s\nMethod: %s\nTime: %s\nSize: %d\n", basePath, path, method, time, size)
	}
}
