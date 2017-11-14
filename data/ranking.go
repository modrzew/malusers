package data

import (
	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers/core"
)

// RankingManager is able to recreate ranking tables
type RankingManager struct {
	DB *gorm.DB
}

// RecreateTemporaryRankingTable recreates temporary ranking table
func (m *RankingManager) RecreateTemporaryRankingTable() {
	m.DB.DropTableIfExists(&core.TemporaryRanking{})
	m.DB.CreateTable(&core.TemporaryRanking{})
}

// PopulateTemporaryRankingTable fills temporary ranking table based on raw anime/manga data
func (m *RankingManager) PopulateTemporaryRankingTable() {
	m.DB.Exec(`
		INSERT INTO temporary_rankings
		(
			created_at, user_id,
			completed_anime, dropped_anime, total_days_anime, episodes_anime,
			completed_manga, dropped_manga, total_days_manga, chapters_manga, volumes_manga
		)
		SELECT
			now(),
			u.id,
			rank() over (order by anime.completed desc) AS completed_anime,
			rank() over (order by anime.dropped desc) AS dropped_anime,
			rank() over (order by anime.days desc) AS total_days_anime,
			rank() over (order by anime.episodes desc) AS episodes_anime,
			rank() over (order by manga.completed desc) AS completed_manga,
			rank() over (order by manga.dropped desc) AS dropped_manga,
			rank() over (order by manga.days desc) AS total_days_manga,
			rank() over (order by manga.chapters desc) AS chapters_manga,
			rank() over (order by manga.volumes desc) AS volumes_manga
		FROM users u
			JOIN anime_stats anime ON u.id=anime.user_id
			JOIN manga_stats manga ON u.id=manga.user_id
		WHERE u.fetched=true
	`)
}

// MigrateRankingResults moves data from temporary to target table. Target table will be dropped.
func (m *RankingManager) MigrateRankingResults() {
	m.DB.DropTableIfExists(&core.Ranking{})
	m.DB.CreateTable(&core.Ranking{})
	m.DB.Exec("INSERT INTO rankings SELECT * FROM temporary_rankings")
}
