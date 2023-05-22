package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dnnybanh/goxen_backend/pkg/handlers"
	"github.com/dnnybanh/goxen_backend/pkg/middleware"
)

func main() {
	// Initialize the Gin engine
	r := gin.Default()
	r.Use(middleware.RequestLogger())

	// Initialize the GORM database connection
	dsn := "host=localhost user=gorm dbname=gorm password=gorm sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	userHandler := handlers.UserHandler{Db: db}

	// Set the router as the default one provided by Gin
	route := r.Group("/api")
	{
		route.POST("/users", userHandler.CreateUser)	
	}

	// Run the server
	if err := r.Run(); err != nil {
		log.Fatalf("server failed to run: %v", err)
	}
}
