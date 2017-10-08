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
		"host=%s port=%d user=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
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

func overseer(db *gorm.DB, maxConcurrent int) {
	finished := make(chan bool, maxConcurrent)
	for {
		var users []User
		db.Limit(5).Find(&users, "Fetching = ? AND Fetched = ?", false, false)
		for i := range users {
			finished <- true
			user := users[i]
			go GetUser(user.Username, finished)
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func monitor(db *gorm.DB, finished chan bool) {
	for {
		var notFetched *int
		var fetching *int
		db.Model(&User{}).Where(&User{Fetched: false}).Count(&notFetched)
		db.Model(&User{}).Where(&User{Fetching: true}).Count(&fetching)
		fmt.Printf("\rTo fetch: %d, fetching: %d\n", *notFetched, *fetching)
		time.Sleep(time.Second)
	}
}

func main() {
	config := ReadConfig()
	db := openDb()
	defer db.Close()
	db.AutoMigrate(&AnimeStats{}, &MangaStats{}, &Relation{}, &User{})

	// Reset all fetching statuses
	db.Model(&User{}).UpdateColumn("fetching", false)

	finished := make(chan bool)
	go monitor(db, finished)
	go overseer(db, config.MaxConcurrent)

	// Maybe trigger first user?
	var inDb *int
	if db.Model(&User{}).Count(&inDb); *inDb == 0 {
		go GetUser("sweetmonia", make(chan bool))
	}

	// Don't quit
	<-finished
}
