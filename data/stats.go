package data

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers/core"
)

// StatsEntry holds data for single "cell" in stats response
type StatsEntry struct {
	Count  int `json:"count"`
	Mean   int `json:"mean"`
	Median int `json:"median"`
}

// StatsKind holds data for single group in stats response
type StatsKind struct {
	Completed StatsEntry `json:"completed"`
	Dropped   StatsEntry `json:"dropped"`
	TotalDays StatsEntry `json:"total_days"`
}

// Stats holds data for whole response
type Stats map[string]StatsKind

// StatsRow is used to map DB results to response
type StatsRow struct {
	Users        int
	BirthYear    int
	Gender       string
	CompletedSum int
	CompletedAvg int
	DroppedSum   int
	DroppedAvg   int
	DaysSum      int
	DaysAvg      int
}

// GenerateStatsTable recreates stats table in database
func GenerateStatsTable(db *gorm.DB) {
	db.DropTableIfExists(&core.GlobalStats{})
	db.CreateTable(&core.GlobalStats{})
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

// GetGlobalStats extracts appropriate stats from database
func GetGlobalStats(db *gorm.DB, kind string, filter string) Stats {
	stats := make(Stats)
	var groupBy string
	if filter == "year" {
		groupBy = "birth_year"
	} else {
		groupBy = "gender"
	}
	query := db.Model(&core.GlobalStats{}).Select(groupBy).Select(fmt.Sprintf(`
		%[1]s, SUM(users) AS users,
		SUM(%[2]s_completed_sum) AS completed_sum,
		ROUND(AVG(%[2]s_completed_avg)) AS completed_avg,
		SUM(%[2]s_dropped_sum) AS dropped_sum,
		ROUND(AVG(%[2]s_dropped_avg)) AS dropped_avg,
		SUM(%[2]s_days_sum) AS days_sum,
		ROUND(AVG(%[2]s_days_avg)) AS days_avg
	`, groupBy, kind)).Group(groupBy)
	var results []StatsRow
	query.Scan(&results)
	for _, row := range results {
		value := StatsKind{
			Completed: StatsEntry{
				Count: row.CompletedSum,
				Mean:  row.CompletedAvg,
			},
			Dropped: StatsEntry{
				Count: row.DroppedSum,
				Mean:  row.DroppedAvg,
			},
			TotalDays: StatsEntry{
				Count: row.DaysSum,
				Mean:  row.DaysAvg,
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
