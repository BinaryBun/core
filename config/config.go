package config

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"os"
)

const configPath = "./config/config.json"

type Config struct {
	IsOnion string `json:"onion"`
}

func LoadConfig() (c *Config) {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return nil
	}

	err = json.NewDecoder(jsonFile).Decode(&c)
	if err != nil {
		return nil
	}

	err = validator.New().Struct(c)
	if err != nil {
		return nil
	}
	return
}
