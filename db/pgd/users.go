package pgd

import (
	"context"

	"github.com/Rajendro1/Talenzen/config"
	"github.com/Rajendro1/Talenzen/model"
)

func CreateUser(user model.User) error {
	_, err := config.PgDbWrite.Exec(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

func UserExists(userID int) bool {
	var exists bool
	err := config.PgDbRead.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", userID).Scan(&exists)
	return err == nil && exists
}

func TaskExists(taskID int) bool {
	var exists bool
	err := config.PgDbRead.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM tasks WHERE id = $1)", taskID).Scan(&exists)
	return err == nil && exists
}
func GetUserByEmail(email string) (model.User, error) {
	var user model.User
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	err := config.PgDbRead.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
