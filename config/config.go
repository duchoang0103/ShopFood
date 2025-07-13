package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	MYSQL_STRING string `json:"MYSQL_STRING"`
}

var AppConfig Config

func LoadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Cannot open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("Cannot decode config JSON: %v", err)
	}
}
