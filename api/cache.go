package api

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers/core"
	"github.com/modrzew/malusers/data"
)

type statsKey struct {
	kind   string
	filter string
}

// Cache is a main cache container
type Cache struct {
	db         *gorm.DB
	totalCount int
	users      map[string]*UserStats
	stats      map[statsKey]*data.Stats
}

// GetCache creates new Cache instance with initial values
func GetCache(db *gorm.DB) *Cache {
	return &Cache{
		db:         db,
		totalCount: -1,
		users:      make(map[string]*UserStats),
		stats:      make(map[statsKey]*data.Stats),
	}
}

// GetCount returns number of all users in the database
func (c *Cache) GetCount() int {
	if c.totalCount == -1 {
		c.db.Model(&core.User{}).Count(&c.totalCount)
	}
	return c.totalCount
}

// GetUser returns data about single user. It will reuse cached data if possible.
func (c *Cache) GetUser(username string) (*UserStats, error) {
	user := core.User{}
	if user, ok := c.users[username]; ok {
		return user, nil
	}
	c.db.Where("username = ?", username).First(&user)
	if c.db.NewRecord(user) {
		return nil, errors.New("not found")
	}
	c.db.Where("user_id = ?", user.ID).Find(&user.AnimeStats)
	c.db.Where("user_id = ?", user.ID).Find(&user.MangaStats)
	c.db.Where("user_id = ?", user.ID).Find(&user.Ranking)
	age := time.Since(user.Birthday.Time).Hours() / 24 / 365
	stats := &UserStats{
		Username:        user.DisplayName,
		Age:             int(age),
		LastUpdate:      user.UpdatedAt.UTC(),
		AnimeStats:      user.AnimeStats,
		MangaStats:      user.MangaStats,
		Ranking:         DBRankingToSchema(user.Ranking),
		TotalUsersCount: c.GetCount(),
	}
	c.users[username] = stats
	return stats, nil
}

// GetStats returns data about global statustics. It will reuse cached data if possible.
func (c *Cache) GetStats(kind string, filter string) data.Stats {
	key := statsKey{kind: kind, filter: filter}
	if stats, ok := c.stats[key]; ok {
		return *stats
	}
	stats := data.GetGlobalStats(c.db, kind, filter)
	c.stats[key] = &stats
	return stats
}
