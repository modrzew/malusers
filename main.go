package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func openDb() *gorm.DB {
	config := ReadConfig()
	connection := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		config.Host,
		config.Port,
		config.Database,
		config.Username,
		config.Password,
		config.SSLMode,
	)
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	return db
}

func overseer(mainDb *gorm.DB, active chan bool, maxConcurrent int) {
	for i := 0; i < maxConcurrent; i++ {
		active <- true
	}
	for {
		if len(active) > 0 {
			users := getUsersToFetch(len(active))
			for i := range users {
				<-active
				user := users[i]
				go GetUser(user.Username, mainDb, active)
			}
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func monitor(db *gorm.DB, active chan bool, maxConcurrent int) {
	for {
		stats := getStatsFromCache()
		fetching := maxConcurrent - len(active)
		fmt.Printf("\rTo fetch: %d, fetching: %d, fetched: %d", stats.toFetch, fetching, stats.fetched)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	config := ReadConfig()
	db := openDb()
	defer db.Close()
	db.AutoMigrate(&AnimeStats{}, &MangaStats{}, &Relation{}, &User{})

	// Reset all fetching statuses
	db.Model(&User{}).UpdateColumn("fetching", false)

	populateCache(db)

	active := make(chan bool, config.MaxConcurrent)
	go monitor(db, active, config.MaxConcurrent)
	go overseer(db, active, config.MaxConcurrent)

	// Maybe trigger first user?
	var inDb *int
	if db.Model(&User{}).Count(&inDb); *inDb == 0 {
		db.Create(&User{Username: "sweetmonia"})
	}

	// Don't quit
	finished := make(chan bool)
	<-finished
}
