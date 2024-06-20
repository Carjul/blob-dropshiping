package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/Carjul/web-service-gin/internal/routes"
	. "github.com/Carjul/web-service-gin/config"
	"github.com/joho/godotenv"
	"log"
	"os"


)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	PORT := os.Getenv("PORT")
    r := gin.Default()

	r.Static("/public", "../public")

	r.LoadHTMLGlob("../public/templates/*")
	
    SetupRoutes(r) 

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	
	if err := Dbconect(); err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer DesconectarDB()

    if err := r.Run(":"+PORT); err != nil && PORT == ""{
        log.Fatalf("Error starting server: %s", err)
    }
}
