package main

import (
	"log"
	"os"

	"desafio-final/cmd/server/routes"
	"desafio-final/pkg/middleware"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

func main() {

	// Recover from panic
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// Router
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())

	// Server
	runApp(engine)
}

func runApp(engine *gin.Engine) {
	// Corremos el servidor
	router := routes.NewRouter(engine)
	// Mapeamos las rutas
	router.MapRoutes()
	if err := engine.Run(port); err != nil {
		panic(err)
	}

}
