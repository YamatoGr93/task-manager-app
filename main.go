package main

import (
	"task-manager-app/db"     // Adjust this import path according to your structure
	"task-manager-app/routes" // Adjust this import path according to your structure

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.InitDB()
	defer db.CloseDB()

	r := gin.Default()

	// Register task routes
	routes.RegisterRoutes(r, db.DB)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
