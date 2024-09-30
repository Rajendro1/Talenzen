package main

import (
	"github.com/Rajendro1/Talenzen/config"
	"github.com/Rajendro1/Talenzen/db/pgd"
	"github.com/Rajendro1/Talenzen/route"
)

func init() {
	config.PgDbWrite = pgd.SetupReadDatabase()
	config.PgDbWrite = pgd.SetupWriteDatabase()

}
func main() {
	route.HandleRequests()
}
