package malusers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
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

func getStats(username string) (*BasicInfo, *AnimeStats, *MangaStats) {
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
func GetUser(username string, db *gorm.DB, finished chan bool) {
	username = strings.ToLower(username)
	user := getOrCreateUser(username, db)
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
	animeStats.Username = username
	db.Create(animeStats)
	mangaStats.Username = username
	db.Create(mangaStats)

	friendsChannel := make(chan []string, 1)
	go getFriends(friendsChannel, username, 0)
	for friendsPage := range friendsChannel {
		for i := range friendsPage {
			friendName := friendsPage[i]
			friend := getOrCreateUser(friendName, db)
			relation := NewRelation(user, friend)
			db.Where(relation).FirstOrCreate(relation)
		}
	}
	user.Fetched = true
	user.Fetching = false
	db.Save(&user)
	finished <- true
}
