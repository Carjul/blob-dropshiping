package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/Carjul/web-service-gin/internal/routes"
	. "github.com/Carjul/web-service-gin/config"
	"path/filepath"
	"github.com/joho/godotenv"
	"log"
	"os"


)

func getDirname(filePath string) string {
	return filepath.Dir(filePath)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	PORT := os.Getenv("PORT")
    r := gin.Default()

	basePath := "./public"
	templatePath := filepath.Join(basePath, "templates/*")
	staticPath := filepath.Join(basePath)

	r.Static("/public", staticPath)
	r.LoadHTMLGlob(templatePath)
	
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
