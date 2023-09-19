package routes

import (
	"database/sql"
	handlerOdontologo "desafio-final/cmd/server/handler/odontologo"
	handlerPaciente "desafio-final/cmd/server/handler/paciente"
	"desafio-final/cmd/server/handler/ping"
	odontologo "desafio-final/internal/domain/odontologo"
	paciente "desafio-final/internal/domain/paciente"
	"desafio-final/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// Router interface defines the methods that any router must implement.
type Router interface {
	MapRoutes()
}

// router is the Gin router.
type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
}

// NewRouter creates a new Gin router.
func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

// MapRoutes maps all routes.
func (r *router) MapRoutes() {
	r.setGroup()
	r.buildPingRoutes()
	r.buildOdontologoRoutes()
	r.buildPacienteRoutes()
}

/* --------------------------------- GROUPS --------------------------------- */
// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

/* ---------------------------------- PING ---------------------------------- */
// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControladorPing()

	r.routerGroup.GET("/ping", pingController.Ping())
}

/* ------------------------------- ODONTOLOGOS ------------------------------ */
// buildOdontologoRoutes maps all routes for the odontologo domain.
func (r *router) buildOdontologoRoutes() {
	// Create a new odontologo controller.
	repository := odontologo.NewRepository(r.db)
	service := odontologo.NewService(repository)
	odontologoController := handlerOdontologo.NewControladorOdontologo(service)

	r.routerGroup.POST("/odontologos", middleware.Authenticate(), odontologoController.Create())
	r.routerGroup.GET("/odontologos", middleware.Authenticate(), odontologoController.GetAll())
	r.routerGroup.GET("/odontologos/:id", middleware.Authenticate(), odontologoController.GetById())
	r.routerGroup.PUT("/odontologos/:id", middleware.Authenticate(), odontologoController.Update())
	r.routerGroup.PATCH("/odontologos/:id", middleware.Authenticate(), odontologoController.UpdateName())
	r.routerGroup.DELETE("/odontologos/:id", middleware.Authenticate(), odontologoController.Delete())
}

/* ------------------------------- PACIENTES ------------------------------ */
// buildPacienteRoutes maps all routes for the paciente domain.
func (r *router) buildPacienteRoutes() {
	// Create a new odontologo controller.
	repository := paciente.NewRepository(r.db)
	service := paciente.NewService(repository)
	odontologoController := handlerPaciente.NewControladorPaciente(service)

	r.routerGroup.POST("/pacientes", middleware.Authenticate(), odontologoController.Create())
}
