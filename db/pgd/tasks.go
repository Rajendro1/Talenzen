package pgd

import (
	"github.com/Rajendro1/Talenzen/config"
	"github.com/Rajendro1/Talenzen/model"
)

func CreateTask(task model.Task) error {
	_, err := config.PgDbWrite.Exec("INSERT INTO tasks (title, description, assigned_user, status, due_date) VALUES ($1, $2, $3, $4, $5)", task.Title, task.Description, task.AssignedUser, task.Status, task.DueDate)
	return err
}
