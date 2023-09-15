package routes

import (
	"desafio-final/cmd/server/handler/ping"

	"github.com/gin-gonic/gin"
)

// Interfaz Router para definir los métodos
type Router interface {
	MapRoutes()
}

// Router de Gin
type router struct {
	engine *gin.Engine
}

// NewRouter crea un nuevo Router de Gin
func NewRouter(engine *gin.Engine) Router {
	return &router{
		engine: engine,
	}
}

// MapRoutes mapea las rutas de la aplicación
func (r *router) MapRoutes() {
	r.buildPingRoutes()
}

// buildPingRoutes mapea las rutas del ping
func (r *router) buildPingRoutes() {
	// Creamos un nuevo controlador
	pingController := ping.NewControladorPing()

	r.engine.GET("/ping", pingController.Ping())
}
