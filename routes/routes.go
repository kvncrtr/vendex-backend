package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kvncrtr/vendex/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// Middlewares
	server.Use(middlewares.CorsMiddleware)

	// Employee Route Group
	employeeRoutes := server.Group("/")
	employeeRoutes.Use(middlewares.AuthenticateEmployee)
	employeeRoutes.Use(middlewares.ClassAorBMiddleware)
	employeeRoutes.POST("/signup", makeEmployeeProfile)
	employeeRoutes.PUT("/employee/:id", updateEmployee)
	employeeRoutes.DELETE("/employee/:id", terminateEmployee)

	partRoutes := server.Group("/")
	partRoutes.Use(middlewares.AuthenticateEmployee)
	partRoutes.Use(middlewares.ClassAorBMiddleware)
	partRoutes.POST("/parts", InsertNewPart)
	partRoutes.GET("/parts", FetchAllParts)
	partRoutes.GET("/parts/:id", GetPartById)
	partRoutes.PUT("/parts/:id", UpdatePart)
	partRoutes.DELETE("/parts/:id", DeletePart)

	// Routes
	server.GET("/employee", fetchAllEmployees)
	server.GET("/employee/:id", getEmployee)
	server.POST("/login", LoginEmployee)
}
