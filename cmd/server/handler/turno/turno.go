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
//  Turno godoc
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
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}
		// If Status: OK - Created
		web.Success(ctx, http.StatusCreated, response)
	}

}

/* --------------------------------- GET ALL -------------------------------- */
//  Turno godoc
//	@Summary		turno example
//	@Description	Get all turnos
//	@Tags			turno
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		500	{object}	web.errorResponse
//	@Router			/turnos [get]
func (c *Controlador) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		turnos, err := c.service.GetAll(ctx)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}

		web.Success(ctx, http.StatusOK, turnos)
	}
}

/* -------------------------------- GET BY ID ------------------------------- */
//  Turno godoc
//	@Summary		turno example
//	@Description	Get turno by id
//	@Tags			turno
//	@Param			id	path	int	true	"id del turno"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/turnos/:id [get]
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

/* ------------------------------- UPDATE ALL ------------------------------- */
//  Turno godoc
//	@Summary		turno example
//	@Description	Update turno by id
//	@Tags			turno
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/turnos/:id [put]
func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request turno.RequestTurno

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Id Inválido")
			return
		}

		response, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}

		web.Success(ctx, http.StatusOK, response)
	}
}

/* --------------------------------- UPDATE NAME --------------------------------- */
//  Turno godoc
//	@Summary		turno example
//	@Description	Update turno name by id
//	@Tags			turno
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/turnos/:id [patch]
func (c *Controlador) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request turno.RequestTurno

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Id Inválido")
			return
		}

		response, err := c.service.UpdateDescripcion(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"id":          response.ID,
			"descripcion": response.Descripcion,
		})
	}
}

/* --------------------------------- DELETE --------------------------------- */
//  Turno godoc
//	@Summary		turno example
//	@Description	Delete turno by id
//	@Tags			turno
//	@Param			id	path	int	true	"id del turno"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/turnos/:id [delete]
func (c *Controlador) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Id Inválido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"message": "turno eliminado correctamente",
		})
	}
}
