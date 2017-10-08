package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

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
			result.days = value
		case "mean score:":
			result.meanScore = value
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
			result.inProgress = value
		case "completed":
			result.completed = value
		case "on-hold":
			result.onhold = value
		case "dropped":
			result.dropped = value
		case "plan to watch":
			result.planned = value
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
			result.rewatched = value
		case "episodes":
			result.episodes = value
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

func getFriends(channel chan []string, username string) {
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

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s\n", r.URL.String())
	})

	url := fmt.Sprintf("https://myanimelist.net/profile/%s/friends", username)
	c.Visit(url)
}

func main() {
	statsChannel := make(chan *AnimeStats)
	go getStats(statsChannel, "mikeone")
	stats := <-statsChannel
	fmt.Printf("%+v\n", stats)

	friendsChannel := make(chan []string)
	go getFriends(friendsChannel, "mikeone")
	friends := <-friendsChannel
	fmt.Printf("%+v\n", friends)
}
