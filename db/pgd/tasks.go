package pgd

import (
	"context"
	"database/sql"

	"github.com/Rajendro1/Talenzen/config"
	"github.com/Rajendro1/Talenzen/model"
)

func CreateTask(task model.Task) (int, error) {
	var id int
	err := config.PgDbWrite.QueryRow(context.Background(), "INSERT INTO tasks (title, description, assigned_user, status, due_date) VALUES ($1, $2, $3, $4, $5) RETURNING id", task.Title, task.Description, task.AssignedUser, task.Status, task.DueDate).Scan(&id)
	return id, err
}

func GetTasks() ([]model.Task, error) {
	var time sql.NullString
	rows, err := config.PgDbRead.Query(context.Background(), "SELECT id, title, description, assigned_user, status, due_date FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.AssignedUser, &t.Status, &time); err != nil {
			return nil, err
		}
		t.DueDate = time.String
		tasks = append(tasks, t)
	}
	return tasks, nil
}
func GetTasksByID(taskId string) (model.Task, error) {
	var t model.Task
	var time sql.NullString

	if err := config.PgDbRead.QueryRow(context.Background(), "SELECT id, title, description, assigned_user, status, due_date FROM tasks WHERE id = $1", taskId).Scan(&t.ID, &t.Title, &t.Description, &t.AssignedUser, &t.Status, &time); err != nil {
		return t, err
	}
	t.DueDate = time.String

	return t, nil
}
func UpdateTask(t model.Task) error {
	_, err := config.PgDbWrite.Exec(context.Background(), "UPDATE tasks SET title = $1, description = $2, assigned_user = $3, status = $4, due_date = $5 WHERE id = $6",
		t.Title, t.Description, t.AssignedUser, t.Status, t.DueDate, t.ID)
	return err
}

func DeleteTask(id int) error {
	// Start a transaction
	tx, err := config.PgDbWrite.Begin(context.Background())
	if err != nil {
		return err // Error when starting a transaction
	}

	// First, try to delete the mapping from the join table
	if _, err := tx.Exec(context.Background(), "DELETE FROM map_users_with_tasks WHERE task_id = $1", id); err != nil {
		tx.Rollback(context.Background()) // Roll back in case of error
		return err
	}

	// Then, delete the task from the tasks table
	if _, err := tx.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id); err != nil {
		tx.Rollback(context.Background()) // Roll back in case of error
		return err
	}

	// Commit the transaction
	if err := tx.Commit(context.Background()); err != nil {
		return err // Error when committing the transaction
	}

	return nil
}

func AssignTask(userID, taskID, assign_by int) error {
	_, err := config.PgDbWrite.Exec(context.Background(), "INSERT INTO map_users_with_tasks (user_id, task_id, assign_by) VALUES ($1, $2, $3)", userID, taskID, assign_by)
	return err
}

func GetTasksByUserID(userID int) ([]model.Task, error) {
	tasks := []model.Task{}
	var time sql.NullString
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
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.AssignedUser, &t.Status, &time); err != nil {
			return nil, err
		}
		t.DueDate = time.String
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTasksByTaskId(taskId int) (model.TaskUsers, error) {
	var t model.TaskUsers
	var time sql.NullTime // More accurate handling of datetime fields
	query := `SELECT t.id, t.title, t.description, t.assigned_user, t.status, t.due_date, m.user_id, u.email
              FROM tasks t
              INNER JOIN map_users_with_tasks m ON t.id = m.task_id
              INNER JOIN users u ON u.id = m.user_id
              WHERE t.id = $1`
	if err := config.PgDbRead.QueryRow(context.Background(), query, taskId).Scan(&t.ID, &t.Title, &t.Description, &t.AssignedUser, &t.Status, &time, &t.UserID, &t.Email); err != nil {
		return t, err
	}
	if time.Valid {
		t.DueDate = time.Time.Format("2006-01-02") // Convert to string if valid
	}
	return t, nil
}
