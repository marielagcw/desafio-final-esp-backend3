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
// Odontologo godoc
//	@Summary		odontologo example
//	@Description	Create a new odontologo
//	@Tags			odontologo
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/odontologos [post]
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
// Odontologo godoc
//	@Summary		odontologo example
//	@Description	Get all odontologos
//	@Tags			odontologo
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		500	{object}	web.errorResponse
//	@Router			/odontologos [get]
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
// Odontologo godoc
//	@Summary		odontologo example
//	@Description	Get odontologo by id
//	@Tags			odontologo
//	@Param			id	path	int	true	"id del odontologo"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/odontologos/:id [get]
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
// Odontologo godoc
//	@Summary		odontologo example
//	@Description	Update odontologo by id
//	@Tags			odontologo
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/odontologos/:id [put]
func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request odontologo.RequestOdontologo

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
// Odontologo godoc
//	@Summary		odontologo example
//	@Description	Update odontologo name by id
//	@Tags			odontologo
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/odontologos/:id [patch]
func (c *Controlador) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request odontologo.RequestOdontologo

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

		response, err := c.service.UpdateName(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"id":     response.ID,
			"nombre": response.Nombre,
		})
	}
}

/* --------------------------------- DELETE --------------------------------- */
// Odontologo godoc
//	@Summary		odontologo example
//	@Description	Delete odontologo by id
//	@Tags			odontologo
//	@Param			id	path	int	true	"id del odontologo"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Failure		400	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/odontologos/:id [delete]
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
			"message": "Odontologo eliminado correctamente",
		})
	}
}
