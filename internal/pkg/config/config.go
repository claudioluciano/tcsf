package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const configFileName = ".tcst-cli.json"

type Config struct {
	SourceAPIKey    string `json:"source_api_key,omitempty"`
	SourceAPISecret string `json:"source_api_secret,omitempty"`
	SourceWorkspace string `json:"source_workspace,omitempty"`
	TargetAPIKey    string `json:"target_api_key,omitempty"`
	TargetAPISecret string `json:"target_api_secret,omitempty"`
	TargetWorkspace string `json:"target_workspace,omitempty"`
}

func GetDefaultConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, configFileName), nil
}

func ConfigExist() bool {
	cPath, err := GetDefaultConfigPath()
	if err != nil {
		return false
	}

	_, err = os.Stat(cPath)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func ReadConfigAsString() (string, error) {
	cPath, err := GetDefaultConfigPath()
	if err != nil {
		return "", err
	}

	b, err := os.ReadFile(cPath)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func ReadConfig(c *Config) error {
	cPath, err := GetDefaultConfigPath()
	if err != nil {
		return err
	}

	b, err := os.ReadFile(cPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, c)
}

func CreateConfigFile() error {
	cPath, err := GetDefaultConfigPath()
	if err != nil {
		return err
	}

	f, err := os.Create(cPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return err
}

func UpdateConfigFile(config *Config) error {
	cPath, err := GetDefaultConfigPath()
	if err != nil {
		return err
	}

	b, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(cPath, b, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}

func GetConfigFromViper() *Config {
	sak := viper.GetString("source_api_key")
	sas := viper.GetString("source_api_secret")
	sws := viper.GetString("source_workspace")

	tak := viper.GetString("target_api_key")
	tas := viper.GetString("target_api_secret")
	tws := viper.GetString("target_workspace")

	if tas == "" {
		tas = sas
	}

	if tak == "" {
		tak = sak
	}

	if tws == "" {
		tws = sws
	}

	return &Config{
		SourceAPIKey:    sak,
		SourceAPISecret: sas,
		SourceWorkspace: sws,
		TargetAPIKey:    tak,
		TargetAPISecret: tas,
		TargetWorkspace: tws,
	}
}

func GetRightCredentialsFromConfig(forceTarget ...bool) (string, string) {
	target := false
	if len(forceTarget) > 0 {
		target = forceTarget[0]
	}

	config := GetConfigFromViper()

	apiKey := config.SourceAPIKey
	apiSecret := config.SourceAPISecret
	if target {
		apiKey = config.TargetAPIKey
		apiSecret = config.TargetAPISecret
	}

	return apiKey, apiSecret
}
