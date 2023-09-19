package paciente

import (
	paciente "desafio-final/internal/domain/paciente"
	"desafio-final/pkg/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service paciente.Service
}

func NewControladorPaciente(service paciente.Service) Controlador {
	return Controlador{
		service: service,
	}
}

/* --------------------------------- CREATE --------------------------------- */
// Paciente godoc
//	@Summary		paciente example
//	@Description	Create a new paciente
//	@Tags			paciente
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/pacientes [post]
func (c *Controlador) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request paciente.RequestPaciente

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
// Paciente godoc
//	@Summary		paciente example
//	@Description	Get all pacientes
//	@Tags			paciente
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]Paciente
//	@Failure		400	{object}	web.error
//	@Failure		500	{object}	web.
//	@Router			/pacientes [get]
func (c *Controlador) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		response, err := c.service.GetAll(ctx)
		// If Status: Internal Server Error
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}
		// If Status: OK
		web.Success(ctx, http.StatusOK, response)
	}
}
