package routes

import (
   . "github.com/Carjul/web-service-gin/internal/handlers"
    "github.com/gin-gonic/gin"
    
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/", HomeHandler)
    r.GET("/articulos", GetArticulosHandler)
    r.POST("/articulo", CreateArticuloHandler)
    r.GET("/articulo/:id", GetArticuloHandler)
    r.PUT("/articulo/:id", UpdateArticuloHandler)
    r.DELETE("/articulo/:id", DeleteArticuloHandler)
}
