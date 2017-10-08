package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/asciimoo/colly"
)

func extractAnimeStats(elem *colly.HTMLElement) *AnimeStats {
	result := new(AnimeStats)
	// Days and mean time
	baseStats := elem.DOM.Find("div.stat-score div.di-tc")
	baseStats.Each(func(_ int, stat *goquery.Selection) {
		// Remove label
		label := stat.Find("span.fw-n")
		labelText := strings.ToLower(strings.TrimSpace(label.Text()))
		label.Remove()
		value, err := strconv.ParseFloat(stat.Text(), 64)
		if err != nil {
			panic(err)
		}
		switch labelText {
		case "days:":
			result.Days = value
		case "mean score:":
			result.MeanScore = value
		}
	})
	// First column
	stats := elem.DOM.Find("ul.stats-status li")
	stats.Each(func(_ int, stat *goquery.Selection) {
		label := strings.ToLower(strings.TrimSpace(stat.Find("a").Text()))
		value, err := strconv.Atoi(strings.Replace(strings.TrimSpace(stat.Find("span").Text()), ",", "", -1))
		if err != nil {
			panic(err)
		}
		switch label {
		case "watching":
			result.InProgress = value
		case "completed":
			result.Completed = value
		case "on-hold":
			result.OnHold = value
		case "dropped":
			result.Dropped = value
		case "plan to watch":
			result.Planned = value
		}
	})
	// Second column
	stats = elem.DOM.Find("ul.stats-data li")
	stats.Each(func(_ int, stat *goquery.Selection) {
		label := strings.ToLower(strings.TrimSpace(stat.Find("span").First().Text()))
		value, err := strconv.Atoi(strings.Replace(strings.TrimSpace(stat.Find("span").Last().Text()), ",", "", -1))
		if err != nil {
			panic(err)
		}
		switch label {
		case "rewatched":
			result.Rewatched = value
		case "episodes":
			result.Episodes = value
		}
	})
	return result
}

func getStats(channel chan *AnimeStats, username string) {
	c := colly.NewCollector()

	c.OnHTML("div.stats.anime", func(e *colly.HTMLElement) {
		stats := extractAnimeStats(e)
		channel <- stats
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s\n", r.URL.String())
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
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		if e.DOM.Find("div.friendBlock").Length() >= 100 {
			go getFriends(channel, username, offset+100)
		} else {
			close(channel)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s\n", r.URL.String())
	})

	url := fmt.Sprintf("https://myanimelist.net/profile/%s/friends?offset=%d", username, offset)
	c.Visit(url)
}

func getUser(db *gorm.DB, username string) {
	user := new(User)
	db.Where(User{Username: username}).FirstOrCreate(&user)
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

	friendsChannel := make(chan []string)
	go getFriends(friendsChannel, username, 0)
	for friendsPage := range friendsChannel {
		for i := range friendsPage {
			friendName := friendsPage[i]
			friend := new(User)
			db.Where(User{Username: friendName}).FirstOrCreate(&friend)
			relation := NewRelation(user, friend)
			db.Where(relation).FirstOrCreate(relation)
		}
	}
	user.Fetched = true
	user.Fetching = false
	db.Save(&user)
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&AnimeStats{}, &MangaStats{}, &Relation{}, &User{})

	go getUser(db, "sweetmonia")
	go getUser(db, "mikeone")

	time.Sleep(time.Millisecond * 10)

	for {
		var notFetched *int
		var fetching *int
		db.Model(&User{}).Where(&User{Fetched: false}).Count(&notFetched)
		db.Model(&User{}).Where(&User{Fetching: true}).Count(&fetching)
		if *fetching == 0 {
			break
		}
		fmt.Printf("To fetch: %d, fetching: %d\n", *notFetched, *fetching)
		time.Sleep(time.Second)
	}
}
