package main

import (
	"fmt"

	"github.com/asciimoo/colly"
)

func getStats(animeChannel chan *AnimeStats, mangaChannel chan *MangaStats, username string) {
	c := colly.NewCollector()

	c.OnHTML("div.stats.anime", func(e *colly.HTMLElement) {
		stats := ExtractAnimeStats(e)
		animeChannel <- stats
	})

	c.OnHTML("div.stats.manga", func(e *colly.HTMLElement) {
		stats := ExtractMangaStats(e)
		mangaChannel <- stats
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
		names := ExtractFriendNames(e)
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

// GetUser obtains stats for single user and their friends
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
	animeStatsChannel := make(chan *AnimeStats)
	mangaStatsChannel := make(chan *MangaStats)
	go getStats(animeStatsChannel, mangaStatsChannel, username)
	animeStats := <-animeStatsChannel
	animeStats.Username = username
	db.Create(animeStats)
	mangaStats := <-mangaStatsChannel
	mangaStats.Username = username
	db.Create(mangaStats)

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
