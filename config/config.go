package config

import "database/sql"

const (
	DBUser       = "postgres"
	DBPassword   = "password"
	DBName       = "taskmanagement"
	DBHost       = "localhost"
	DBPort       = "5432"
	JWTSecret    = "your_jwt_secret_here"
	SMTPServer   = "smtp.gmail.com"
	SMTPPort     = "587"
	SMTPUsername = "your.email@gmail.com"
	SMTPPassword = "yourpassword"
)

var (
	PgDbWrite *sql.DB
	PgDbRead  *sql.DB
)

func GetReadDatabaseConnectionString() string {
	return "host=" + DBHost + " user=" + DBUser + " password=" + DBPassword + " dbname=" + DBName + " sslmode=disable"
}
func GetWriteDatabaseConnectionString() string {
	return "host=" + DBHost + " user=" + DBUser + " password=" + DBPassword + " dbname=" + DBName + " sslmode=disable"
}
