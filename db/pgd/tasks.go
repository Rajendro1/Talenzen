package pgd

import (
	"github.com/Rajendro1/Talenzen/config"
	"github.com/Rajendro1/Talenzen/model"
)

func CreateTask(task model.Task) error {
	_, err := config.PgDbWrite.Exec("INSERT INTO tasks (title, description, assigned_user, status, due_date) VALUES ($1, $2, $3, $4, $5)", task.Title, task.Description, task.AssignedUser, task.Status, task.DueDate)
	return err
}

func GetTasks() ([]model.Task, error) {
	rows, err := config.PgDbRead.Query("SELECT id, title, description, assigned_user, status, due_date FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.AssignedUser, &t.Status, &t.DueDate); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func UpdateTask(t model.Task) error {
	_, err := config.PgDbWrite.Exec("UPDATE tasks SET title = $1, description = $2, assigned_user = $3, status = $4, due_date = $5 WHERE id = $6",
		t.Title, t.Description, t.AssignedUser, t.Status, t.DueDate, t.ID)
	return err
}

func DeleteTask(id int) error {
	_, err := config.PgDbWrite.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}

func AssignTask(userID, taskID int) error {
	_, err := config.PgDbWrite.Exec("INSERT INTO map_users_with_tasks (user_id, task_id) VALUES ($1, $2)", userID, taskID)
	return err
}
