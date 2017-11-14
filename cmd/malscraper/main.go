package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers/core"
	"github.com/modrzew/malusers/scraper"
)

func overseer(mainDb *gorm.DB, active chan bool, maxConcurrent int) {
	for i := 0; i < maxConcurrent; i++ {
		active <- true
	}
	for {
		if len(active) > 0 {
			users := scraper.GetUsersToFetchFromCache(len(active))
			if len(users) > 0 {
				for i := range users {
					<-active
					user := users[i]
					go scraper.GetUser(user.Username, mainDb, active)
				}
			} else {
				// Wait for a bit, and query database
				time.Sleep(time.Second * 60)
				scraper.AddUsersToFetchFromDatabase(mainDb)
			}
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func monitor(db *gorm.DB, active chan bool, maxConcurrent int) {
	for {
		stats := scraper.GetStatsFromCache()
		fetching := maxConcurrent - len(active)
		fmt.Printf("\rTo fetch: %d, fetching: %d, fetched: %d", stats.ToFetch, fetching, stats.Fetched)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	config := core.ReadConfig()
	db := core.OpenDb()
	defer db.Close()
	db.AutoMigrate(&core.AnimeStats{}, &core.MangaStats{}, &core.Relation{}, &core.User{})

	// Reset all fetching statuses
	db.Model(&core.User{}).UpdateColumn("fetching", false)

	scraper.PopulateCache(db)

	active := make(chan bool, config.Scraper.MaxConcurrent)
	go monitor(db, active, config.Scraper.MaxConcurrent)
	go overseer(db, active, config.Scraper.MaxConcurrent)

	// Maybe trigger first user?
	var inDb *int
	if db.Model(&core.User{}).Count(&inDb); *inDb == 0 {
		db.Create(&core.User{Username: "sweetmonia"})
	}

	// Don't quit
	finished := make(chan bool)
	<-finished
}
