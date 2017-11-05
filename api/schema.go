package api

import (
	"time"

	"github.com/modrzew/malusers/core"
)

// UserStats contains response structure for single user
type UserStats struct {
	Username        string          `json:"username"`
	Age             int             `json:"age"`
	LastUpdate      time.Time       `json:"lastUpdate"`
	AnimeStats      core.AnimeStats `json:"animeStats"`
	MangaStats      core.MangaStats `json:"mangaStats"`
	Ranking         Ranking         `json:"ranking"`
	TotalUsersCount int             `json:"totalUsers"`
}

// Ranking contains ranking information for single user. To be used in API response.
type Ranking struct {
	Anime AnimeRanking `json:"anime"`
	Manga MangaRanking `json:"manga"`
}

// AnimeRanking contains anime ranking information for single user. To be used in API response.
type AnimeRanking struct {
	Completed int     `json:"completed"`
	Dropped   int     `json:"dropped"`
	TotalDays float64 `json:"totalDays"`
	Episodes  int     `json:"episodes"`
}

// MangaRanking contains manga ranking information for single user. To be used in API response.
type MangaRanking struct {
	Completed int     `json:"completed"`
	Dropped   int     `json:"dropped"`
	TotalDays float64 `json:"totalDays"`
	Chapters  int     `json:"chapters"`
	Volumes   int     `json:"volumes"`
}

// DBRankingToSchema maps DB structure into API, JSONable structure
func DBRankingToSchema(fromDb core.Ranking) Ranking {
	return Ranking{
		Anime: AnimeRanking{
			Completed: fromDb.CompletedAnime,
			Dropped:   fromDb.DroppedAnime,
			TotalDays: fromDb.TotalDaysAnime,
			Episodes:  fromDb.EpisodesAnime,
		},
		Manga: MangaRanking{
			Completed: fromDb.CompletedManga,
			Dropped:   fromDb.DroppedManga,
			TotalDays: fromDb.TotalDaysManga,
			Chapters:  fromDb.ChaptersManga,
			Volumes:   fromDb.VolumesManga,
		},
	}
}
