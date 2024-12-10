package main

import (
	"graded-challenge-2-andiadisasmita/database"
	"graded-challenge-2-andiadisasmita/routes"
)

func main() {
	// Connect to the database
	database.ConnectDB()

	// Initialize routes
	r := routes.InitRoutes()

	// Start server
	r.Run(":8080") // Default port
}
