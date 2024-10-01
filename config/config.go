package config

import (
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	JWTSecret    = "1ab5413f57b98068b9a2243e7fb7ee428954686c919bbfe160f50a51b5bf6f670760473dc17f209b344d18b49ae04ae366f4963411bae7f1b5dd987cc9a76fd5e39d64a20e7df8bee9673e6bb33e331e470aec9614046b0d95fd1565c8c2a5bd6512a42e1fcd9d8e34df876ca2f3d12c83a583118e4ad3fffdf030729d4c57e467c33eb83b30860174a3352f5738e011e0dcffde8413b11e25315d4fe4e07288cbeb198d353b2673230cb24071cfdf88305c74a0af8bd7da58aeac17bbbe4caa6a86c4e13a12fd5b8f550e7174a35636ced531db1ff251a4e7a5ad1d64351489b63b7479d8c19f48b90677e67f4d2557f9d76d11d43ca71468b946dff73f2082"
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
