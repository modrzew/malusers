package main

import (
	"fmt"
	"os"

	"github.com/modrzew/malusers"
)

func main() {
	command := os.Args[1]
	db := malusers.OpenDb()
	if command == "ranking" {
		manager := &malusers.RankingManager{DB: db}
		fmt.Println("Recreating temporary table")
		manager.RecreateTemporaryRankingTable()
		fmt.Println("Populating temporary table")
		manager.PopulateTemporaryRankingTable()
		fmt.Println("Moving temporary to permanent")
		manager.MigrateRankingResults()
		fmt.Println("Done")
	} else if command == "stats" {
		fmt.Println("Regenerating global stats")
		malusers.GenerateStatsTable(db)
	}
}
