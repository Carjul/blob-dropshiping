package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/Carjul/web-service-gin/internal/routes"
	
	"log"


)

func main() {
    r := gin.Default()

	r.Static("/public", "../public")

	r.LoadHTMLGlob("../public/templates/*")
	
    SetupRoutes(r) 

    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error starting server: %s", err)
    }
}
