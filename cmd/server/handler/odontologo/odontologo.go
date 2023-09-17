package odontologo

import (
	odontologo "desafio-final/internal/domain/odontologo"
	"desafio-final/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service odontologo.Service
}

func NewControladorOdontologo(service odontologo.Service) Controlador {
	return Controlador{
		service: service,
	}
}

/* --------------------------------- CREATE --------------------------------- */
func (c *Controlador) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request odontologo.RequestOdontologo

		err := ctx.Bind(&request)

		// If Status: Bad Request
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		response, err := c.service.Create(ctx, request)
		// If Status: Internal Server Error
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}
		// If Status: OK - Created
		web.Success(ctx, http.StatusCreated, response)
	}

}

/* --------------------------------- GET ALL -------------------------------- */
func (c *Controlador) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		odontologos, err := c.service.GetAll(ctx)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}

		web.Success(ctx, http.StatusOK, odontologos)
	}
}

/* -------------------------------- GET BY ID ------------------------------- */
func (c *Controlador) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Id Inválido")
			return
		}

		response, err := c.service.GetById(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}

		web.Success(ctx, http.StatusOK, response)
	}
}
