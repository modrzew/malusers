package api

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers"
)

type Cache struct {
	db         *gorm.DB
	totalCount int
	users      map[string]*UserStats
}

func GetCache(db *gorm.DB) *Cache {
	return &Cache{
		db:         db,
		totalCount: -1,
		users:      make(map[string]*UserStats),
	}
}

func (c *Cache) GetCount() int {
	if c.totalCount == -1 {
		c.db.Model(&malusers.User{}).Count(&c.totalCount)
	}
	return c.totalCount
}

func (c *Cache) GetUser(username string) (*UserStats, error) {
	user := malusers.User{}
	if user, ok := c.users[username]; ok {
		return user, nil
	}
	c.db.Where("username = ?", username).First(&user)
	if c.db.NewRecord(user) {
		return nil, errors.New("not found")
	}
	c.db.Where("username = ?", username).Find(&user.AnimeStats)
	c.db.Where("username = ?", username).Find(&user.MangaStats)
	c.db.Where("username = ?", username).Find(&user.Ranking)
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
