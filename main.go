package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kvncrtr/vendex/db"
	"github.com/kvncrtr/vendex/routes"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Init Database
	db.InitDB()
	defer db.DB.Close()

	// Start Server
	server := gin.Default()

	// Execute Routes
	routes.RegisterRoutes(server)

	// Listen and serve
	server.Run(":8080")
}

/*
	result := fmt.Sprintf(`

		"errorMessage": %s

	`, err.Error())
	fmt.Println(result)
*/
