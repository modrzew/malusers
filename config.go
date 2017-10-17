package malusers

import (
	"encoding/json"
	"os"
	"path"
)

type dbConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	SslMode  string
}

type apiConfig struct {
	Host string
	Port int
}

type scraperConfig struct {
	MaxConcurrent int
}

// Configuration holds database settings
type Configuration struct {
	Database dbConfig
	Scraper  scraperConfig
	API      apiConfig
}

var configuration = Configuration{}
var read = false

// ReadConfig returns config from JSON file
func ReadConfig() *Configuration {
	if read {
		return &configuration
	}
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path.Join(dir, "config.json"))
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	read = true
	return &configuration
}
