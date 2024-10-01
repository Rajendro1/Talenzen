package model

import (
	"github.com/golang-jwt/jwt"
)

type (
	Task struct {
		ID           int    `json:"id"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		AssignedUser int    `json:"assigned_user"`
		Status       string `json:"status"`
		DueDate      string `json:"due_date"`
	}
	TaskUsers struct {
		ID           int    `json:"id"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		AssignedUser int    `json:"assigned_user"`
		Status       string `json:"status"`
		DueDate      string `json:"due_date"`

		UserID int    `json:"user_id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
	}

	User struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"-"`
	}
	UserInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	Claims struct {
		UserID int    `json:"user_id"`
		Email  string `json:"email"`
		jwt.StandardClaims
	}
)
