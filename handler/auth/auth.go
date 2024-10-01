package auth

import (
	"net/http"
	"strings"

	"github.com/Rajendro1/Talenzen/db/pgd"
	"github.com/Rajendro1/Talenzen/middleware"
	"github.com/Rajendro1/Talenzen/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *gin.Context) {
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	user, err := pgd.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found", "details": err.Error()})
		return
	}

	// Check if the hashed password and the given password match
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials", "details": "Incorrect password"})
		return
	}

	// Generate JWT token
	token, err := middleware.CreateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func RegisterHandler(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data", "details": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password", "details": err.Error()})
		return
	}
	newUser.Password = string(hashedPassword)

	// Create user in the database
	if err := pgd.CreateUser(newUser); err != nil {
		if strings.Contains(err.Error(), "unique_violation") {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
