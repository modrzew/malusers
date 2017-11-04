package core

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// AnimeStats is structure for holding anime statistics
type AnimeStats struct {
	gorm.Model `json:"-"`
	Username   string  `json:"-"`
	InProgress int     `json:"inProgress"`
	Completed  int     `json:"completed"`
	OnHold     int     `json:"onHold"`
	Dropped    int     `json:"dropped"`
	Planned    int     `json:"planned"`
	Rewatched  int     `json:"rewatched"`
	Days       float64 `json:"totalDays"`
	MeanScore  float64 `json:"meanScore"`
	Episodes   int     `json:"totalEpisodes"`
}

// MangaStats is structure for holding manga statistics
type MangaStats struct {
	gorm.Model `json:"-"`
	Username   string  `json:"-"`
	InProgress int     `json:"inProgress"`
	Completed  int     `json:"completed"`
	OnHold     int     `json:"onHold"`
	Dropped    int     `json:"dropped"`
	Planned    int     `json:"planned"`
	Rewatched  int     `json:"rewatched"`
	Days       float64 `json:"totalDays"`
	MeanScore  float64 `json:"meanScore"`
	Chapters   int     `json:"totalChapters"`
	Volumes    int     `json:"totalVolumes"`
}

// Relation - `from` user having `to` as friend
type Relation struct {
	User1   User
	User1ID uint
	User2   User
	User2ID uint
}

// NewRelation returns Relation where users are alphabetized
func NewRelation(user1 *User, user2 *User) *Relation {
	if user1.Username <= user2.Username {
		return &Relation{
			User1ID: user1.ID,
			User2ID: user2.ID,
		}
	}
	return &Relation{
		User1ID: user2.ID,
		User2ID: user1.ID,
	}
}

// BasicInfo holds info about user without any database info
type BasicInfo struct {
	Birthday pq.NullTime
	Gender   string
}

// User holds info about user and whether they were fetched
type User struct {
	gorm.Model
	Username    string
	DisplayName string
	Birthday    pq.NullTime `gorm:"type:date"`
	Gender      string
	Fetched     bool `gorm:"index"`
	Fetching    bool `gorm:"index"`
	AnimeStats  AnimeStats
	MangaStats  MangaStats
	Ranking     Ranking
}

// TemporaryRanking holds info about user's ranking temporarily when table is recreated
type TemporaryRanking struct {
	gorm.Model
	Username       string
	CompletedAnime int
	CompletedManga int
	DroppedAnime   int
	DroppedManga   int
	TotalDaysAnime float64
	TotalDaysManga float64
	EpisodesAnime  int
	ChaptersManga  int
	VolumesManga   int
}

// Ranking holds info about user's ranking
type Ranking struct {
	TemporaryRanking
}

// GlobalStats holds info about ranking for all users grouped by birth year and gender
type GlobalStats struct {
	gorm.Model
	Users             int
	BirthYear         int
	Gender            string
	AnimeCompletedSum int
	AnimeCompletedAvg int
	AnimeDroppedSum   int
	AnimeDroppedAvg   int
	AnimeDaysSum      int
	AnimeDaysAvg      int
	MangaCompletedSum int
	MangaCompletedAvg int
	MangaDroppedSum   int
	MangaDroppedAvg   int
	MangaDaysSum      int
	MangaDaysAvg      int
}
