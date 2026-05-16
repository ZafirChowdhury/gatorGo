package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	configPath := homeDir + configFileName

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

func (c *Config) SetUser(username string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := homeDir + configFileName
	c.CurrentUserName = username

	fileData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	if err := os.WriteFile(configPath, fileData, 0644); err != nil {
		return err
	}

	return nil
}
