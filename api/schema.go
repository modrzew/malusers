package api

import (
	"time"

	"github.com/modrzew/malusers"
)

// UserStats contains response structure for single user
type UserStats struct {
	Username   string              `json:"username"`
	Age        int                 `json:"age"`
	LastUpdate time.Time           `json:"lastUpdate"`
	AnimeStats malusers.AnimeStats `json:"animeStats"`
	MangaStats malusers.MangaStats `json:"mangaStats"`
}
