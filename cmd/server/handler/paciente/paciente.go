package paciente

import (
	paciente "desafio-final/internal/domain/paciente"
	"desafio-final/pkg/web"
	"net/http"
	"strconv"

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

/* --------------------------------- GET BY ID ------------------------------- */
// Paciente godoc
//	@Summary		paciente example
//	@Description	Get paciente by id
//	@Tags			paciente
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"paciente id"
//	@Success		200	{object}	Paciente
//	@Failure		400	{object}	web.
//	@Failure		500	{object}	web.
//	@Router			/pacientes/{id} [get]
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
//	@Summary		paciente example
//	@Description	Update paciente by id
//	@Tags			paciente
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"paciente id"
//	@Success		200	{object}	Paciente
//	@Failure		400	{object}	web.
//	@Failure		500	{object}	web.
//	@Router			/pacientes/{id} [put]
func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")

		parsedId, err := strconv.Atoi(id)

		// If Status: Bad Request
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		var request paciente.RequestPaciente

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

/* --------------------------------- PATCH --------------------------------- */
// Paciente godoc
//	@Summary		paciente example
//	@Description	Patch paciente by id
//	@Tags			paciente
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"paciente id"
//	@Success		200	{object}	Paciente
//	@Failure		400	{object}	web.
//	@Failure		500	{object}	web.
//	@Router			/pacientes/{id} [patch]
func (c *Controlador) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")

		parsedId, err := strconv.Atoi(id)

		// If Status: Bad Request
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		var request paciente.RequestPaciente

		err = ctx.Bind(&request)

		// If Status: Bad Request
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "Bad Request")
			return
		}

		response, err := c.service.Patch(ctx, parsedId, request)
		// If Status: Internal Server Error
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "Internal Server Error")
			return
		}
		// If Status: OK
		web.Success(ctx, http.StatusOK, response)
	}
}
