package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kvncrtr/vendex/db"
	"github.com/kvncrtr/vendex/routes"
	_ "github.com/lib/pq"
)

func main() {
	// Init Database
	db.InitDB()
	defer db.DB.Close()

	// Start Server
	server := gin.Default()

	// Execute Routes
	routes.RegisterRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	// Listen and serve
	server.Run(":" + port)
}

/*
	result := fmt.Sprintf(`

		"errorMessage": %s

	`, err.Error())
	fmt.Println(result)
*/
