package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	Task struct {
		ID           int       `json:"id"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		AssignedUser int       `json:"assigned_user"`
		Status       string    `json:"status"`
		DueDate      time.Time `json:"due_date"`
	}

	User struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"-"`
	}
	Claims struct {
		UserID int    `json:"user_id"`
		Email  string `json:"email"`
		jwt.StandardClaims
	}
)
