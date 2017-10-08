package main

import (
	"encoding/json"
	"os"
)

// Configuration holds database settings
type Configuration struct {
	Host          string
	Port          int
	Database      string
	Username      string
	Password      string
	SSLMode       string
	MaxConcurrent int
}

var configuration = Configuration{}
var read = false

// ReadConfig returns config from JSON file
func ReadConfig() *Configuration {
	if read {
		return &configuration
	}
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	read = true
	return &configuration
}
