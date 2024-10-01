package config

import "github.com/jackc/pgx/v5/pgxpool"

const (
	JWTSecret    = "60GqY3XPu9YShUosfYvQR8b8Yo1wM1x51SvjsUGVoezDm1cf4VvkGSSKBqgn6ZUL"
	SMTPServer   = "smtp.gmail.com"
	SMTPPort     = "587"
	SMTPUsername = "rajandroprosad@gmail.com"
	SMTPPassword = "rajendro@1"
)

var (
	PgDbWrite *pgxpool.Pool
	PgDbRead  *pgxpool.Pool

	ReadPostgresHost         = "localhost"
	ReadPostgresPort         = "5432"
	ReadPostgresUser         = "postgres"
	ReadPostgresPassword     = "postgres"
	ReadPostgresDatabaseName = "deposite_withdrawal"

	WritePostgresHost         = "localhost"
	WritePostgresPort         = "5432"
	WritePostgresUser         = "postgres"
	WritePostgresPassword     = "postgres"
	WritePostgresDatabaseName = "deposite_withdrawal"

	Err error
)
