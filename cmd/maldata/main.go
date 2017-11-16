package main

import (
	"fmt"
	"os"

	"github.com/modrzew/malusers/core"
	"github.com/modrzew/malusers/data"
)

func main() {
	command := os.Args[1]
	db := core.OpenDb()
	if command == "ranking" {
		manager := &data.RankingManager{DB: db}
		fmt.Println("Recreating temporary table")
		manager.RecreateTemporaryRankingTable()
		fmt.Println("Populating temporary table")
		manager.PopulateTemporaryRankingTable()
		fmt.Println("Moving temporary to permanent")
		manager.MigrateRankingResults()
		fmt.Println("Done")
	} else if command == "stats" {
		fmt.Println("Regenerating global stats")
		data.GenerateStatsTable(db)
	} else if command == "fetched" {
		fmt.Println("Marking users to refresh")
		data.MarkUsersToFetch(db)
		fmt.Println("Done")
	} else {
		fmt.Println("Please use either `maldata ranking`, `maldata fetched` or `maldata stats`.")
	}
}
