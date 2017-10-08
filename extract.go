package main

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/asciimoo/colly"
)

// ExtractAnimeStats gets anime stats from single user's profile page
func ExtractAnimeStats(elem *colly.HTMLElement) *AnimeStats {
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

// ExtractMangaStats gets manga stats from single user's profile page
func ExtractMangaStats(elem *colly.HTMLElement) *MangaStats {
	result := new(MangaStats)
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
		case "reading":
			result.InProgress = value
		case "completed":
			result.Completed = value
		case "on-hold":
			result.OnHold = value
		case "dropped":
			result.Dropped = value
		case "plan to read":
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
		case "reread":
			result.Rewatched = value
		case "chapters":
			result.Chapters = value
		case "volumes":
			result.Volumes = value
		}
	})
	return result
}

// ExtractFriendNames gets friends names from single page
func ExtractFriendNames(e *colly.HTMLElement) []string {
	names := []string{}
	sel := e.DOM.Find("div.friendBlock strong")
	for i := range sel.Nodes {
		elem := sel.Eq(i)
		names = append(names, elem.Text())
	}
	return names
}
