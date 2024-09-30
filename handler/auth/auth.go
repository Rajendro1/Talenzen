package auth

import (
	"net/http"
	"strings"

	"github.com/Rajendro1/Talenzen/db/pgd"
	"github.com/Rajendro1/Talenzen/model"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func RegisterHandler(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data", "details": err.Error()})
		return
	}

	if err := pgd.CreateUser(newUser); err != nil {
		// Check for a unique violation error (example for PostgreSQL)
		if strings.Contains(err.Error(), "unique_violation") {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
