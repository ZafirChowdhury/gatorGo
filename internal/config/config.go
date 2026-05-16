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
