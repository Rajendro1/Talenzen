package pgd

import (
	"context"

	"github.com/Rajendro1/Talenzen/config"
	"github.com/Rajendro1/Talenzen/model"
)

func CreateTask(task model.Task) error {
	_, err := config.PgDbWrite.Exec(context.Background(), "INSERT INTO tasks (title, description, assigned_user, status, due_date) VALUES ($1, $2, $3, $4, $5)", task.Title, task.Description, task.AssignedUser, task.Status, task.DueDate)
	return err
}

func GetTasks() ([]model.Task, error) {
	rows, err := config.PgDbRead.Query(context.Background(), "SELECT id, title, description, assigned_user, status, due_date FROM tasks")
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
func GetTasksByID(taskId string) (model.Task, error) {
	var t model.Task
	if err := config.PgDbRead.QueryRow(context.Background(), "SELECT id, title, description, assigned_user, status, due_date FROM tasks WHERE id = $1", taskId).Scan(&t.ID, &t.Title, &t.Description, &t.AssignedUser, &t.Status, &t.DueDate); err != nil {
		return t, err
	}
	return t, nil
}
func UpdateTask(t model.Task) error {
	_, err := config.PgDbWrite.Exec(context.Background(), "UPDATE tasks SET title = $1, description = $2, assigned_user = $3, status = $4, due_date = $5 WHERE id = $6",
		t.Title, t.Description, t.AssignedUser, t.Status, t.DueDate, t.ID)
	return err
}

func DeleteTask(id int) error {
	_, err := config.PgDbWrite.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	return err
}

func AssignTask(userID, taskID int, assign_by string) error {
	_, err := config.PgDbWrite.Exec(context.Background(), "INSERT INTO map_users_with_tasks (user_id, task_id, assign_by) VALUES ($1, $2, $3)", userID, taskID)
	return err
}

func GetTasksByUserID(userID int) ([]model.Task, error) {
	tasks := []model.Task{}
	query := `SELECT t.id, t.title, t.description, t.assigned_user, t.status, t.due_date
              FROM tasks t
              INNER JOIN map_users_with_tasks m ON t.id = m.task_id
              WHERE m.user_id = $1`
	rows, err := config.PgDbRead.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.AssignedUser, &t.Status, &t.DueDate); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
