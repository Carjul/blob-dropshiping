package routes

import (
   . "github.com/Carjul/web-service-gin/internal/handlers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/", HomeHandler)
    r.GET("/ping", PingHandler)
}
