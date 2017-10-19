package api

import (
	"time"

	"github.com/modrzew/malusers"
)

// UserStats contains response structure for single user
type UserStats struct {
	Username        string              `json:"username"`
	Age             int                 `json:"age"`
	LastUpdate      time.Time           `json:"lastUpdate"`
	AnimeStats      malusers.AnimeStats `json:"animeStats"`
	MangaStats      malusers.MangaStats `json:"mangaStats"`
	Ranking         Ranking             `json:"ranking"`
	TotalUsersCount int                 `json:"totalUsers"`
}

type Ranking struct {
	Anime AnimeRanking `json:"anime"`
	Manga MangaRanking `json:"manga"`
}

type AnimeRanking struct {
	Completed int     `json:"completed"`
	Dropped   int     `json:"dropped"`
	TotalDays float64 `json:"totalDays"`
	Episodes  int     `json:"episodes"`
}

type MangaRanking struct {
	Completed int     `json:"completed"`
	Dropped   int     `json:"dropped"`
	TotalDays float64 `json:"totalDays"`
	Chapters  int     `json:"chapters"`
	Volumes   int     `json:"volumes"`
}

func DBRankingToSchema(fromDb malusers.Ranking) Ranking {
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
