package main

import (
	"fmt"

	"github.com/modrzew/malusers"
)

func main() {
	db := malusers.OpenDb()
	manager := &malusers.RankingManager{DB: db}
	fmt.Println("Recreating temporary table")
	manager.RecreateTemporaryRankingTable()
	fmt.Println("Populating temporary table")
	manager.PopulateTemporaryRankingTable()
	fmt.Println("Moving temporary to permanent")
	manager.MigrateRankingResults()
	fmt.Println("Done")
}
