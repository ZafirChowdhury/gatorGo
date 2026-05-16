package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonBytes, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	if err := json.Unmarshal(jsonBytes, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func write(cfg *Config) error {
	fileData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	configPath, err := getConfigFilePath()

	if err := os.WriteFile(configPath, fileData, 0644); err != nil {
		return err
	}

	return nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	if err := write(c); err != nil {
		return err
	}

	return nil
}
