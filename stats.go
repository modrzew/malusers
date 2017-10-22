package malusers

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

type StatsEntry struct {
	Count  int `json:"count"`
	Mean   int `json:"mean"`
	Median int `json:"median"`
}

type StatsKind struct {
	Completed StatsEntry `json:"completed"`
	Dropped   StatsEntry `json:"dropped"`
	TotalDays StatsEntry `json:"total_days"`
}

type Stats map[string]StatsKind

func GenerateStatsTable(db *gorm.DB) {
	db.DropTableIfExists(&GlobalStats{})
	db.CreateTable(&GlobalStats{})
	db.Exec(`
		INSERT INTO global_stats
		(
			created_at, updated_at,
			users, birth_year, gender,
			anime_completed_sum, anime_completed_avg,
			anime_dropped_sum, anime_dropped_avg,
			anime_days_sum, anime_days_avg,
			manga_completed_sum, manga_completed_avg,
			manga_dropped_sum, manga_dropped_avg,
			manga_days_sum, manga_days_avg
		)
		SELECT
			NOW(), NOW(),
			COUNT(*) users, EXTRACT(YEAR FROM birthday) AS birth_year, gender,
			SUM(anime_stats.completed) AS anime_completed_sum,
			ROUND(AVG(anime_stats.completed)) AS anime_completed_avg,
			SUM(anime_stats.dropped) AS anime_dropped_sum,
			ROUND(AVG(anime_stats.dropped)) AS anime_dropped_avg,
			SUM(anime_stats.days) AS anime_days_sum,
			ROUND(AVG(anime_stats.days)) AS anime_days_avg,
			SUM(manga_stats.completed) AS manga_completed_sum,
			ROUND(AVG(manga_stats.completed)) AS manga_completed_avg,
			SUM(manga_stats.dropped) AS manga_dropped_sum,
			ROUND(AVG(manga_stats.dropped)) AS manga_dropped_avg,
			SUM(manga_stats.days) AS manga_days_sum,
			ROUND(AVG(manga_stats.days)) AS manga_days_avg
		FROM users
		JOIN anime_stats ON users.username=anime_stats.username
		JOIN manga_stats ON users.username=manga_stats.username
		WHERE
			birthday IS NOT NULL
			AND EXTRACT(YEAR FROM birthday)>1900
			AND EXTRACT(YEAR FROM birthday)<2020
		GROUP BY birth_year, gender
		ORDER BY birth_year;
	`)
}

func GetAnimeStats(db *gorm.DB, filter string) Stats {
	stats := make(Stats)
	db.LogMode(true)
	var groupBy string
	if filter == "year" {
		groupBy = "birth_year"
	} else {
		groupBy = "gender"
	}
	query := db.Model(&GlobalStats{}).Select(groupBy).Select(fmt.Sprintf(`
		%s, SUM(users) AS users,
		SUM(anime_completed_sum) AS anime_completed_sum,
		ROUND(AVG(anime_completed_avg)) AS anime_completed_avg,
		SUM(anime_dropped_sum) AS anime_dropped_sum,
		ROUND(AVG(anime_dropped_avg)) AS anime_dropped_avg,
		SUM(anime_days_sum) AS anime_days_sum,
		ROUND(AVG(anime_days_avg)) AS anime_days_avg
	`, groupBy)).Group(groupBy)
	var results []GlobalStats
	query.Find(&results)
	for _, row := range results {
		value := StatsKind{
			Completed: StatsEntry{
				Count: row.AnimeCompletedSum,
				Mean:  row.AnimeCompletedAvg,
			},
			Dropped: StatsEntry{
				Count: row.AnimeDroppedSum,
				Mean:  row.AnimeDroppedAvg,
			},
			TotalDays: StatsEntry{
				Count: row.AnimeDaysSum,
				Mean:  row.AnimeDaysAvg,
			},
		}
		if filter == "year" {
			stats[strconv.Itoa(row.BirthYear)] = value
		} else if filter == "gender" {
			stats[row.Gender] = value
		}
	}
	return stats
}
