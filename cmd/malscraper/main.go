package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers"
)

func overseer(mainDb *gorm.DB, active chan bool, maxConcurrent int) {
	for i := 0; i < maxConcurrent; i++ {
		active <- true
	}
	for {
		if len(active) > 0 {
			users := malusers.GetUsersToFetch(len(active))
			for i := range users {
				<-active
				user := users[i]
				go malusers.GetUser(user.Username, mainDb, active)
			}
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func monitor(db *gorm.DB, active chan bool, maxConcurrent int) {
	for {
		stats := malusers.GetStatsFromCache()
		fetching := maxConcurrent - len(active)
		fmt.Printf("\rTo fetch: %d, fetching: %d, fetched: %d", stats.ToFetch, fetching, stats.Fetched)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	config := malusers.ReadConfig()
	db := malusers.OpenDb()
	defer db.Close()
	db.AutoMigrate(&malusers.AnimeStats{}, &malusers.MangaStats{}, &malusers.Relation{}, &malusers.User{})

	// Reset all fetching statuses
	db.Model(&malusers.User{}).UpdateColumn("fetching", false)

	malusers.PopulateCache(db)

	active := make(chan bool, config.Scraper.MaxConcurrent)
	go monitor(db, active, config.Scraper.MaxConcurrent)
	go overseer(db, active, config.Scraper.MaxConcurrent)

	// Maybe trigger first user?
	var inDb *int
	if db.Model(&malusers.User{}).Count(&inDb); *inDb == 0 {
		db.Create(&malusers.User{Username: "sweetmonia"})
	}

	// Don't quit
	finished := make(chan bool)
	<-finished
}
