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

func overseer(mainDb *gorm.DB, dbsChannel chan *gorm.DB, maxConcurrent int) {
	for i := 0; i < maxConcurrent; i++ {
		dbsChannel <- openDb()
	}
	for {
		if len(dbsChannel) > 0 {
			users := getUsersToFetch(len(dbsChannel))
			for i := range users {
				db := <-dbsChannel
				user := users[i]
				go GetUser(user.Username, db, dbsChannel)
			}
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func monitor(db *gorm.DB, dbsChannel chan *gorm.DB, maxConcurrent int) {
	for {
		stats := getStatsFromCache()
		fetching := maxConcurrent - len(dbsChannel)
		fmt.Printf("\rTo fetch: %d, fetching: %d, fetched: %d", stats.toFetch, fetching, stats.fetched)
		time.Sleep(time.Millisecond * 200)
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

	dbsChannel := make(chan *gorm.DB, config.MaxConcurrent)
	go monitor(db, dbsChannel, config.MaxConcurrent)
	go overseer(db, dbsChannel, config.MaxConcurrent)

	// Maybe trigger first user?
	var inDb *int
	if db.Model(&User{}).Count(&inDb); *inDb == 0 {
		db.Create(&User{Username: "sweetmonia"})
	}

	// Don't quit
	finished := make(chan bool)
	<-finished
}
