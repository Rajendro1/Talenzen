package auth

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func LoginHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func RegisterHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}