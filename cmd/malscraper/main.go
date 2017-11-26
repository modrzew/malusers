package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers/core"
	"github.com/modrzew/malusers/scraper"
)

func overseer(mainDb *gorm.DB, active chan bool, relations chan []core.Relation, maxConcurrent int) {
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
					go scraper.GetUser(user.Username, mainDb, active, relations)
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

func friendsOverseer(mainDb *gorm.DB, channel chan []core.Relation) {
	for {
		relations := <-channel
		if len(relations) > 0 {
			core.SaveRelations(mainDb, relations)
		}
		time.Sleep(time.Millisecond * 10)
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
	relations := make(chan []core.Relation, config.Scraper.MaxConcurrent)
	go monitor(db, active, config.Scraper.MaxConcurrent)

	// Maybe trigger first user?
	var inDb *int
	if db.Model(&core.User{}).Count(&inDb); *inDb == 0 {
		scraper.GetOrCreateUser("sweetmonia", db)
	}

	go overseer(db, active, relations, config.Scraper.MaxConcurrent)
	go friendsOverseer(db, relations)

	// Don't quit
	finished := make(chan bool)
	<-finished
}
