package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "message": "Welcome to the Home Page!",
    })
}

func PingHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}
