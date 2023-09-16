package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"user-service/pkg/db"
	"user-service/routes"
)

func main() {
	db.Init()

	r := gin.Default()
	routes.UserRoutes(r)

	err := r.Run(":18080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err.Error())
	}
}
