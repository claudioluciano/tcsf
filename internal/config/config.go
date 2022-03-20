package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const configFileName = ".tcst-cli.json"

type Config struct {
	SourceAPIKey    string `survey:"source_api_key" json:"source_api_key,omitempty"`
	SourceAPISecret string `survey:"source_api_secret" json:"source_api_secret,omitempty"`
	SourceWorkspace string `survey:"source_workspace" json:"source_workspace,omitempty"`
	TargetAPIKey    string `survey:"target_api_key" json:"target_api_key,omitempty"`
	TargetAPISecret string `survey:"target_api_secret" json:"target_api_secret,omitempty"`
	TargetWorkspace string `survey:"target_workspace" json:"target_workspace,omitempty"`
}

func GetDefaultConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, configFileName), nil
}

func ReadConfigFile() (string, error) {
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
