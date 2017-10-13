package main

import "github.com/modrzew/malusers"

func main() {
	db := malusers.OpenDb()
	manager := &malusers.RankingManager{DB: db}
	manager.RecreateTemporaryRankingTable()
	manager.PopulateTemporaryRankingTable()
	manager.MigrateRankingResults()
}
