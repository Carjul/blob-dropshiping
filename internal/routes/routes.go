package routes

import (
   . "github.com/Carjul/web-service-gin/internal/handlers"
    "github.com/gin-gonic/gin"
    
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/", HomeHandler)
    r.GET("/api/v1/articulos", GetArticulosHandler)
    r.POST("/api/v1/articulo", CreateArticuloHandler)
    r.GET("/api/v1/articulo/:id", GetArticuloHandler)
    r.PUT("/api/v1/articulo/:id", UpdateArticuloHandler)
    r.DELETE("/api/v1/articulo/:id", DeleteArticuloHandler)
}
