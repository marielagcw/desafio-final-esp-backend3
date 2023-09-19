package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"desafio-final/cmd/server/routes"
	"desafio-final/docs"
	"desafio-final/pkg/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const port = ":8080"

/* --------------------------------- SWAGGER -------------------------------- */
//	@title			Desafio Final - Certified Tech developer
//	@version		1.0
//	@description	This is a proyect for Desafio Final
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {

	// Recover from panic
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// Load the environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	db := connectDB()

	// Router
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())
	engine.Use(middleware.Authenticate())

	/* --------------------------------- Swagger -------------------------------- */
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// use ginSwagger middleware to serve the API docs
	engine.GET("/swagger/*any", middleware.Authenticate(), ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Server
	runApp(db, engine)

	// Close database connection
	defer db.Close()

}

/* --------------------------------- RUN APP -------------------------------- */
func runApp(db *sql.DB, engine *gin.Engine) {
	// Run the application
	router := routes.NewRouter(engine, db)
	// Map all routes
	router.MapRoutes()
	if err := engine.Run(port); err != nil {
		panic(err)
	}

}

/* -------------------------------- DATABASE -------------------------------- */
func connectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string
	dbUsername = "root"
	dbPassword = "root"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "db_desafio_final"

	// Create data source
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open connection
	db, err := sql.Open("mysql", dataSource)

	// Check error
	if err != nil {
		panic(err)
	}

	// Check connection
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// Return connection
	return db
}
