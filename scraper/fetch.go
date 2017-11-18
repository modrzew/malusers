package scraper

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"github.com/modrzew/malusers/core"
)

// Client is a wrapper over http.Client
type Client struct {
	client *http.Client
}

var logger = log.New(os.Stdout, "[fetch] ", 0)

// Get sends GET request and converts response into goquery.Document
func (c *Client) Get(url string) *goquery.Document {
	request, err := http.NewRequest("GET", url, nil)
	// logger.Printf("GET %s\n", url)
	if err != nil {
		panic(err)
	}
	response, err := c.client.Do(request)
	if err != nil {
		panic(err)
	}
	// logger.Printf("Got response for %s\n", url)
	document, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		panic(err)
	}
	return document
}

func getClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

func getStats(username string) (*core.BasicInfo, *core.AnimeStats, *core.MangaStats) {
	client := getClient()
	url := fmt.Sprintf("https://myanimelist.net/profile/%s", username)
	document := client.Get(url)
	return ExtractBasicInfo(document.Find("ul.user-status")),
		ExtractAnimeStats(document.Find("div.stats.anime")),
		ExtractMangaStats(document.Find("div.stats.manga"))
}

func getFriends(channel chan []string, username string, offset int) {
	client := getClient()
	url := fmt.Sprintf("https://myanimelist.net/profile/%s/friends?offset=%d", username, offset)
	document := client.Get(url)
	block := document.Find("div.majorPad")
	names := ExtractFriendNames(block)
	channel <- names
	if block.Find("div.friendBlock").Length() >= 100 {
		getFriends(channel, username, offset+100)
	} else {
		close(channel)
	}
}

// GetUser obtains stats for single user and their friends
func GetUser(username string, db *gorm.DB, finished chan bool, relationsChannel chan []core.Relation) {
	user := GetOrCreateUser(username, db)
	if user.Fetched {
		return
	}
	removeFromToFetch(user.Username)
	user.Fetched = false
	user.Fetching = true
	db.Save(&user)
	basicInfo, animeStats, mangaStats := getStats(username)
	user.Birthday = basicInfo.Birthday
	user.Gender = basicInfo.Gender
	user.Location = basicInfo.Location
	animeStats.UserID = user.ID
	db.Save(animeStats)
	mangaStats.UserID = user.ID
	db.Save(mangaStats)

	friendsChannel := make(chan []string, 1)
	go getFriends(friendsChannel, username, 0)
	var relations []core.Relation
	for friendsPage := range friendsChannel {
		for i := range friendsPage {
			friendName := friendsPage[i]
			friend := GetOrCreateUser(friendName, db)
			relations = append(relations, core.NewRelation(user, friend))
		}
	}
	relationsChannel <- relations
	user.Fetched = true
	user.Fetching = false
	db.Save(&user)
	finished <- true
}
