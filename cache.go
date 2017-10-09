package main

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
)

var usersCache = make(map[string]*User)
var usersToFetch = make(map[string]*User)
var mux sync.Mutex

func populateCache(db *gorm.DB) {
	fmt.Println("Populating cache...")
	mux.Lock()
	defer mux.Unlock()
	var users []User
	db.Model(&User{}).Find(&users)
	for i := range users {
		user := users[i]
		usersCache[user.Username] = &user
		if !user.Fetched && !user.Fetching {
			usersToFetch[user.Username] = &user
		}
	}
	fmt.Printf("%d users in cache\n", len(users))
}

func getOrCreateUser(username string, db *gorm.DB) *User {
	mux.Lock()
	defer mux.Unlock()
	if user, ok := usersCache[username]; ok {
		return user
	}
	user := &User{Username: username}
	db.Create(&user)
	usersCache[user.Username] = user
	usersToFetch[user.Username] = user
	return user
}

func removeFromToFetch(username string) {
	mux.Lock()
	delete(usersToFetch, username)
	mux.Unlock()
}

type stats struct {
	fetched int
	toFetch int
}

func getStatsFromCache() *stats {
	mux.Lock()
	defer mux.Unlock()
	toFetch := len(usersToFetch)
	return &stats{
		fetched: len(usersCache) - toFetch,
		toFetch: toFetch,
	}
}

func getUsersToFetch(limit int) []*User {
	mux.Lock()
	defer mux.Unlock()
	var users []*User
	for _, user := range usersToFetch {
		if len(users) > limit {
			return users
		}
		users = append(users, user)
	}
	return users
}
