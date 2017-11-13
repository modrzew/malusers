package scraper

import (
	"fmt"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers/core"
)

var usersCache = make(map[string]*core.User)
var usersToFetch = make(map[string]*core.User)
var mux sync.Mutex

// PopulateCache fetches all users from database into cache
func PopulateCache(db *gorm.DB) {
	fmt.Println("Populating cache...")
	mux.Lock()
	defer mux.Unlock()
	var users []core.User
	db.Model(&core.User{}).Find(&users)
	for i := range users {
		user := users[i]
		usersCache[user.Username] = &user
		if !user.Fetched && !user.Fetching {
			usersToFetch[user.Username] = &user
		}
	}
	fmt.Printf("%d users in cache\n", len(users))
}

func getOrCreateUser(username string, db *gorm.DB) *core.User {
	mux.Lock()
	defer mux.Unlock()
	displayName := username
	username = strings.ToLower(username)
	if user, ok := usersCache[username]; ok {
		return user
	}
	user := &core.User{Username: username, DisplayName: displayName}
	db.Create(&user)
	usersCache[username] = user
	usersToFetch[username] = user
	return user
}

func removeFromToFetch(username string) {
	mux.Lock()
	delete(usersToFetch, username)
	mux.Unlock()
}

// CacheStats contains information about current scraping process, to be displayed in the command line
type CacheStats struct {
	Fetched int
	ToFetch int
}

// GetStatsFromCache returns information about current scraping process
func GetStatsFromCache() *CacheStats {
	mux.Lock()
	defer mux.Unlock()
	toFetch := len(usersToFetch)
	return &CacheStats{
		Fetched: len(usersCache) - toFetch,
		ToFetch: toFetch,
	}
}

// GetUsersToFetchFromCache returns next batch of users to download
func GetUsersToFetchFromCache(limit int) []*core.User {
	mux.Lock()
	defer mux.Unlock()
	var users []*core.User
	for _, user := range usersToFetch {
		if len(users) > limit {
			return users
		}
		users = append(users, user)
	}
	return users
}

// AddUsersToFetchFromDatabase adds to cache all users from database that are not yet fetched
func AddUsersToFetchFromDatabase(db *gorm.DB) {
	mux.Lock()
	defer mux.Unlock()
	var users []core.User
	db.Model(&core.User{}).Where("fetched = ?", false).Find(&users)
	for i := range users {
		user := users[i]
		if _, ok := usersCache[user.Username]; !ok {
			usersCache[user.Username] = &user
		}
		if !user.Fetched && !user.Fetching {
			usersToFetch[user.Username] = &user
		}
	}
}
