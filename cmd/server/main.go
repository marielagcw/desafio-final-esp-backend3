package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"desafio-final/cmd/server/routes"
	"desafio-final/pkg/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

	// Connect to the database
	db := connectDB()

	// Router
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())

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
