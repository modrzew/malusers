package main

import (
	"github.com/jinzhu/gorm"
)

// Stats is base structure for either anime/manga stats; see AnimeStats or MangaStats
type stats struct {
}

// AnimeStats is structure for holding anime statistics
type AnimeStats struct {
	gorm.Model
	Username   string
	InProgress int
	Completed  int
	OnHold     int
	Dropped    int
	Planned    int
	Rewatched  int
	Days       float64
	MeanScore  float64
	Episodes   int
}

// MangaStats is structure for holding manga statistics
type MangaStats struct {
	gorm.Model
	Username   string
	InProgress int
	Completed  int
	OnHold     int
	Dropped    int
	Planned    int
	Rewatched  int
	Days       float64
	MeanScore  float64
	Chapters   int
	Volumes    int
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

// User holds info about user and whether they were fetched
type User struct {
	gorm.Model
	Username string
	Fetched  bool
	Fetching bool
}
