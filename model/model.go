package model

import "time"

type Task struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	AssignedUser int       `json:"assigned_user"`
	Status       string    `json:"status"`
	DueDate      time.Time `json:"due_date"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
