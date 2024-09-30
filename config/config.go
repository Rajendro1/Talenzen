package config

import "database/sql"

const (
	DBUser       = "postgres"
	DBPassword   = "password"
	DBName       = "taskmanagement"
	DBHost       = "localhost"
	DBPort       = "5432"
	JWTSecret    = "b0ba52c740b85a998a1654e9264d1d95f3e69c330d037719f46c84b6e1b746985e922e70f0aa27e071bdf22b81769d9a3a9cd7c34bd66ea74acd52a43f736c77e3ede41df25211b7fb00e35ed69a2851f65978989fe22ea3ccb5816c0de9ed33fbeb3d60f44af9626b024a40f2eb93e6a44057d7c9eb0fc0262f7cc716f4b1c01c29eda250407999dec9b979523a0d358c2d5d890828d846588baeed2ea05623dee7abe5ddfe5acf287a839f9cf470cda4af1a11b94b7c2b62d7e66b678dfd2c8eb311bf1f349659f13b05c635254d27af9524a7b14ac4eec6bc377f7b5cadbc005d633da3550ece889ddcaaec66d1807833786086afc99ad73f92c1aadeffdd"
	SMTPServer   = "smtp.gmail.com"
	SMTPPort     = "587"
	SMTPUsername = "rajandroprosad@gmail.com"
	SMTPPassword = "rajendro@1"
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
