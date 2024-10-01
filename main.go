package main

import (
	"log"

	"github.com/Rajendro1/Talenzen/config"
	"github.com/Rajendro1/Talenzen/db/pgd"
	"github.com/Rajendro1/Talenzen/route"
)

func init() {
	config.PgDbWrite, config.Err = pgd.WriteConnectPostgresDB()
	if config.Err != nil {
		log.Fatal("Write Postgres Connection Err: ", config.Err)
	}
	config.PgDbRead, config.Err = pgd.ReadConnectPostgresDB()
	if config.Err != nil {
		log.Fatal("Read Postgres Connection Err: ", config.Err)
	}
	config.WG.Add(1)
	go func() {
		defer config.WG.Done()
		if bbMigrationErr := pgd.DbMigration(config.PgDbWrite); bbMigrationErr != nil {
			log.Println("DbMigration Error: ", bbMigrationErr)
		}
	}()

}
func main() {
	route.HandleRequests()
	config.WG.Wait()
}
