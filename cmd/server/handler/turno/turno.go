package turno

import (
	turno "desafio-final/internal/domain/turno"
	"desafio-final/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service turno.Service
}

func NewControladorTurno(service turno.Service) Controlador {
	return Controlador{
		service: service,
	}
}

/* --------------------------------- CREATE --------------------------------- */
// Paciente godoc
//	@Summary		turno example
//	@Description	Create a new turno
//	@Tags			turno
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/turnos [post]
func (c *Controlador) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request turno.RequestTurno

		err := ctx.Bind(&request)

		// If Status: Bad Request
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		response, err := c.service.Create(ctx, request)
		// If Status: Internal Server Error
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s. %s", "Internal Server Error", err)
			return
		}
		// If Status: OK - Created
		web.Success(ctx, http.StatusCreated, response)
	}

}

/* --------------------------------- GET ALL -------------------------------- */
// Paciente godoc
//	@Summary		turno example
//	@Description	Get all turnos
//	@Tags			turno
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]Paciente
//	@Failure		400	{object}	web.error
//	@Failure		500	{object}	web.
//	@Router			/turnos [get]
func (c *Controlador) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		response, err := c.service.GetAll(ctx)
		// If Status: Internal Server Error
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s. %s", "Internal Server Error", err)
			return
		}
		// If Status: OK
		web.Success(ctx, http.StatusOK, response)
	}
}

/* --------------------------------- GET BY ID ------------------------------- */
// Paciente godoc
//	@Summary		turno example
//	@Description	Get turno by id
//	@Tags			turno
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"turno id"
//	@Success		200	{object}	Paciente
//	@Failure		400	{object}	web.
//	@Failure		500	{object}	web.
//	@Router			/turnos/{id} [get]
func (c *Controlador) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")

		parsedId, err := strconv.Atoi(id)

		// If Status: Bad Request
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		response, err := c.service.GetById(ctx, parsedId)
		// If Status: Internal Server Error
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}
		// If Status: OK
		web.Success(ctx, http.StatusOK, response)
	}
}

/* --------------------------------- UPDATE --------------------------------- */
// Paciente godoc
//	@Summary		turno example
//	@Description	Update turno by id
//	@Tags			turno
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"turno id"
//	@Success		200	{object}	Paciente
//	@Failure		400	{object}	web.
//	@Failure		500	{object}	web.
//	@Router			/turnos/{id} [put]
func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")

		parsedId, err := strconv.Atoi(id)

		// If Status: Bad Request
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		var request turno.RequestTurno

		err = ctx.Bind(&request)

		// If Status: Bad Request
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		response, err := c.service.Update(ctx, parsedId, request)
		// If Status: Internal Server Error
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}
		// If Status: OK
		web.Success(ctx, http.StatusOK, response)
	}
}

/* --------------------------------- DELETE --------------------------------- */
// Paciente godoc
//	@Summary		turno example
//	@Description	Delete turno by id
//	@Tags			turno
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"turno id"
//	@Success		200	{object}	Paciente
//	@Failure		400	{object}	web.
//	@Failure		500	{object}	web.
//	@Router			/turnos/{id} [delete]
func (c *Controlador) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")

		parsedId, err := strconv.Atoi(id)

		// If Status: Bad Request
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		err = c.service.Delete(ctx, parsedId)
		// If Status: Internal Server Error
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}
		// If Status: No Content
		ctx.Status(http.StatusNoContent)
	}
}

/* --------------------------------- GET ALL BY 'Paciente' Dni -------------------------------- */
// Paciente godoc
//	@Summary		turno example
//	@Description	Get all turnos by Paciente Dni
//	@Tags			turno
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]Paciente
//	@Failure		400	{object}	web.error
//	@Failure		500	{object}	web.
//	@Router			/turnos/{dni} [get]
func (c *Controlador) GetAllByPacienteDni() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		dni := ctx.Param("dni")

		parsedDni, err := strconv.Atoi(dni)

		// If Status: Bad Request
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		response, err := c.service.GetAllByPacienteDni(ctx, parsedDni)
		// If Status: Internal Server Error
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}
		// If Status: OK
		web.Success(ctx, http.StatusOK, response)
	}
}
