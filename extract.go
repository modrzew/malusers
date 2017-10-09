package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"

	"github.com/PuerkitoBio/goquery"
)

// ExtractAnimeStats gets anime stats from single user's profile page
func ExtractAnimeStats(elem *goquery.Selection) *AnimeStats {
	result := new(AnimeStats)
	// Days and mean time
	baseStats := elem.Find("div.stat-score div.di-tc")
	baseStats.Each(func(_ int, stat *goquery.Selection) {
		// Remove label
		label := stat.Find("span.fw-n")
		labelText := strings.ToLower(strings.TrimSpace(label.Text()))
		label.Remove()
		value, err := strconv.ParseFloat(strings.Replace(strings.TrimSpace(stat.Text()), ",", "", -1), 64)
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
	stats := elem.Find("ul.stats-status li")
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
	stats = elem.Find("ul.stats-data li")
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
func ExtractMangaStats(elem *goquery.Selection) *MangaStats {
	result := new(MangaStats)
	// Days and mean time
	baseStats := elem.Find("div.stat-score div.di-tc")
	baseStats.Each(func(_ int, stat *goquery.Selection) {
		// Remove label
		label := stat.Find("span.fw-n")
		labelText := strings.ToLower(strings.TrimSpace(label.Text()))
		label.Remove()
		value, err := strconv.ParseFloat(strings.Replace(strings.TrimSpace(stat.Text()), ",", "", -1), 64)
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
	stats := elem.Find("ul.stats-status li")
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
	stats = elem.Find("ul.stats-data li")
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
func ExtractFriendNames(elem *goquery.Selection) []string {
	names := []string{}
	sel := elem.Find("div.friendBlock strong")
	for i := range sel.Nodes {
		elem := sel.Eq(i)
		names = append(names, elem.Text())
	}
	return names
}

// ExtractBasicInfo gets basic info about the user
func ExtractBasicInfo(elem *goquery.Selection) *BasicInfo {
	info := &BasicInfo{}
	elem.Find("li").Each(func(_ int, stat *goquery.Selection) {
		label := strings.ToLower(strings.TrimSpace(stat.Find("span.user-status-title").Text()))
		value := strings.TrimSpace(stat.Find("span.user-status-data").Text())
		switch label {
		case "birthday":
			if strings.Index(value, ",") == -1 {
				value = value + ", 1900"
			}
			parsed, err := time.Parse("Jan 2, 2006", value)
			if err == nil {
				info.Birthday = pq.NullTime{Time: parsed, Valid: true}
			} else {
				info.Birthday = pq.NullTime{Valid: false}
			}
		case "gender":
			info.Gender = getGender(value)
		}
	})
	return info
}

func getGender(value string) string {
	switch strings.ToLower(value) {
	case "male":
		return "M"
	case "female":
		return "F"
	case "non-binary":
		return "X"
	default:
		return ""
	}
}
