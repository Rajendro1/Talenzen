package pgd

import (
	"database/sql"

	"github.com/Rajendro1/Talenzen/model"
)

func CreateTask(db *sql.DB, task model.Task) error {
	_, err := db.Exec("INSERT INTO tasks (title, description, assigned_user, status, due_date) VALUES ($1, $2, $3, $4, $5)", task.Title, task.Description, task.AssignedUser, task.Status, task.DueDate)
	return err
}
