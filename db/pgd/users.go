package pgd

import (
	"github.com/Rajendro1/Talenzen/config"
	"github.com/Rajendro1/Talenzen/model"
)

func CreateUser(user model.User) error {
	_, err := config.PgDbWrite.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func UserExists(userID int) bool {
	var exists bool
	err := config.PgDbRead.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", userID).Scan(&exists)
	return err == nil && exists
}

func TaskExists(taskID int) bool {
	var exists bool
	err := config.PgDbRead.QueryRow("SELECT EXISTS(SELECT 1 FROM tasks WHERE id = $1)", taskID).Scan(&exists)
	return err == nil && exists
}
