package config

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

const ConfigFile = ".mh.json"

type Config struct {
	RootURL string `json:"root_url"`
}

func IsValidURL(url string) bool {
	re := regexp.MustCompile(`^https?://[a-zA-Z0-9.-]+(:[0-9]+)?(/.*)?$`)
	return re.MatchString(url)
}

func LoadConfig() (Config, error) {
	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		return Config{}, fmt.Errorf("configuration file not found, please run 'init' first")
	}

	var conf Config
	if err := json.Unmarshal(data, &conf); err != nil {
		return Config{}, fmt.Errorf("invalid configuration file")
	}

	return conf, nil
}

func SaveConfig(conf Config) error {
	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(ConfigFile, data, 0644)
}
