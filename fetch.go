package main

import (
	"fmt"

	"github.com/asciimoo/colly"
)

func getStats(channel chan *AnimeStats, username string) {
	c := colly.NewCollector()

	c.OnHTML("div.stats.anime", func(e *colly.HTMLElement) {
		stats := ExtractAnimeStats(e)
		channel <- stats
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Printf("visiting %s\n", r.URL.String())
	})

	url := fmt.Sprintf("https://myanimelist.net/profile/%s", username)
	c.Visit(url)
}

func getFriends(channel chan []string, username string, offset int) {
	c := colly.NewCollector()

	c.OnHTML("div.majorPad", func(e *colly.HTMLElement) {
		names := []string{}
		sel := e.DOM.Find("div.friendBlock strong")
		for i := range sel.Nodes {
			elem := sel.Eq(i)
			names = append(names, elem.Text())
		}
		channel <- names
		if e.DOM.Find("div.friendBlock").Length() >= 100 {
			getFriends(channel, username, offset+100)
		} else {
			close(channel)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Printf("visiting %s\n", r.URL.String())
	})

	url := fmt.Sprintf("https://myanimelist.net/profile/%s/friends?offset=%d", username, offset)
	c.Visit(url)
}

// GetUser dd
func GetUser(username string, finished chan bool) {
	db := openDb()
	defer db.Close()

	user := new(User)
	db.Where(&User{Username: username}).Attrs(&User{Fetching: true}).FirstOrCreate(&user)
	if user.Fetched {
		return
	}
	user.Fetched = false
	user.Fetching = true
	db.Save(&user)
	statsChannel := make(chan *AnimeStats)
	go getStats(statsChannel, username)
	stats := <-statsChannel
	stats.Username = username
	db.Create(stats)

	friendsChannel := make(chan []string, 1)
	go getFriends(friendsChannel, username, 0)
	for friendsPage := range friendsChannel {
		for i := range friendsPage {
			friendName := friendsPage[i]
			friend := new(User)
			db.Where(&User{Username: friendName}).FirstOrCreate(&friend)
			relation := NewRelation(user, friend)
			db.Where(relation).FirstOrCreate(relation)
		}
	}
	user.Fetched = true
	user.Fetching = false
	db.Save(&user)
	<-finished
}
