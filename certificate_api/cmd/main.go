package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/certificate_api/server/routes"
)

func main() {
	r := gin.Default()

	// Initialize all routes
	routes.RegisterRoutes(r)

	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}
