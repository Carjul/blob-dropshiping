package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    . "github.com/Carjul/web-service-gin/internal/services"
    "github.com/Carjul/web-service-gin/internal/models"
)

func HomeHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "message": "Welcome to the Home Page!",
    })
}


func GetArticulosHandler(c *gin.Context) {
    response := GetArticulos()
    c.JSON(http.StatusOK, gin.H{
        "data": response,
    })
}

func GetArticuloHandler(c *gin.Context) {
    id := c.Param("id")
    response := GetArticulo(id);
    c.JSON(http.StatusOK, gin.H{
        "data": response,
    })
}

func CreateArticuloHandler(c *gin.Context) {
    var articulo models.Articulo
    c.BindJSON(&articulo)
    CreateArticulo(articulo)

    c.JSON(http.StatusOK, gin.H{
        "message": "CreateArticuloHandler",
    })
}

func UpdateArticuloHandler(c *gin.Context) {
    id := c.Param("id")
    var articulo models.Articulo
    c.BindJSON(&articulo)
    UpdateArticulo(id, articulo)

    c.JSON(http.StatusOK, gin.H{
        "message": "UpdateArticuloHandler",
    })
}

func DeleteArticuloHandler(c *gin.Context) {
    id := c.Param("id")
    DeleteArticulo(id)
    c.JSON(http.StatusOK, gin.H{
        "message": "DeleteArticuloHandler",
    })
}