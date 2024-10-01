package config

import (
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	JWTSecret    = "4znwRLTg-KB1PZHplJpoTceIIv44u9LA"
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
	ReadPostgresDatabaseName = "taskmanagement"

	WritePostgresHost         = "localhost"
	WritePostgresPort         = "5432"
	WritePostgresUser         = "postgres"
	WritePostgresPassword     = "postgres"
	WritePostgresDatabaseName = "taskmanagement"

	Err error
	WG  sync.WaitGroup
)
