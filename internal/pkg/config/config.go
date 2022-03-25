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

func GetConfigFromViper(forceTarget ...bool) *Config {
	target := false
	if len(forceTarget) > 0 {
		target = forceTarget[0]
	}

	sourceAPIKey := viper.GetString("source_api_key")
	sourceAPISecret := viper.GetString("source_api_secret")
	sourceWorkspace := viper.GetString("source_workspace")
	targetAPIKey := viper.GetString("target_api_key")
	targetAPISecret := viper.GetString("target_api_secret")
	targetWorkspace := viper.GetString("target_workspace")

	if targetAPIKey == "" || targetAPISecret == "" || targetWorkspace == "" {
		targetAPIKey = sourceAPIKey
		targetAPISecret = sourceAPISecret
		targetWorkspace = sourceWorkspace
	}

	if target {
		tempAPIKey := targetAPIKey
		tempAPISecret := targetAPISecret
		tempWorkspace := targetWorkspace

		targetAPIKey = sourceAPIKey
		targetAPISecret = sourceAPISecret
		targetWorkspace = sourceWorkspace

		sourceAPIKey = tempAPIKey
		sourceAPISecret = tempAPISecret
		sourceWorkspace = tempWorkspace
	}

	return &Config{
		SourceAPIKey:    sourceAPIKey,
		SourceAPISecret: sourceAPISecret,
		SourceWorkspace: sourceWorkspace,
		TargetAPIKey:    targetAPIKey,
		TargetAPISecret: targetAPISecret,
		TargetWorkspace: targetWorkspace,
	}
}

func GetRightCredentialsFromConfig(forceTarget ...bool) (string, string) {
	target := false
	if len(forceTarget) > 0 {
		target = forceTarget[0]
	}

	config := GetConfigFromViper(target)

	apiKey := config.SourceAPIKey
	apiSecret := config.SourceAPISecret
	return apiKey, apiSecret
}
