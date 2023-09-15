package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controlador struct {
}

func NewControladorPing() *Controlador {
	return &Controlador{}
}

func (c *Controlador) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
