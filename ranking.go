package malusers

import (
	"github.com/jinzhu/gorm"
)

type RankingManager struct {
	DB *gorm.DB
}

func (m *RankingManager) RecreateTemporaryRankingTable() {
	m.DB.DropTableIfExists(&TemporaryRanking{})
	m.DB.CreateTable(&TemporaryRanking{})
}

func (m *RankingManager) PopulateTemporaryRankingTable() {
	m.DB.Exec(`
		INSERT INTO temporary_rankings
		(
			created_at, updated_at,
			username, completed_anime, dropped_anime, total_days_anime, episodes_anime,
			completed_manga, dropped_manga, total_days_manga, chapters_manga, volumes_manga
		)
		SELECT
			now(),
			now(),
			u.username,
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
			JOIN anime_stats anime ON u.username=anime.username
			JOIN manga_stats manga ON u.username=manga.username
		WHERE u.fetched=true
	`)
}

func (m *RankingManager) MigrateRankingResults() {
	m.DB.DropTableIfExists(&Ranking{})
	m.DB.CreateTable(&Ranking{})
	m.DB.Exec("INSERT INTO rankings SELECT * FROM temporary_rankings")
}
