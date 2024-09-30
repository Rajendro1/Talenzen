package pgd

import (
	"database/sql"

	"github.com/Rajendro1/Talenzen/config"
	"github.com/Rajendro1/Talenzen/model"
)

func CreateUser(db *sql.DB, user model.User) error {
	_, err := config.PgDbWrite.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}
