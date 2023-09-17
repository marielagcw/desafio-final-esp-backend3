package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := os.Getenv("TOKEN")
		tokenHeader := ctx.GetHeader("Authorization")

		if tokenHeader != token {
			ctx.JSON(http.StatusUnauthorized, "Unauthorized")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		} else {
			ctx.Next()
		}
	}
}
