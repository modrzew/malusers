package main

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/asciimoo/colly"
)

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
