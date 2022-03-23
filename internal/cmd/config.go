package cmd

import (
	"fmt"
	"twilio_copy_studio_flow/internal/pkg/config"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name:   "source_api_key",
		Prompt: &survey.Input{Message: "Source API Key"},
	},
	{
		Name:   "source_api_secret",
		Prompt: &survey.Input{Message: "Source API Secret"},
	},
	{
		Name:   "source_workspace",
		Prompt: &survey.Input{Message: "Source Workspace"},
	},
	{
		Name:   "target_api_key",
		Prompt: &survey.Input{Message: "Target API Key"},
	},
	{
		Name:   "target_api_secret",
		Prompt: &survey.Input{Message: "Target API Secret"},
	},
	{
		Name:   "target_workspace",
		Prompt: &survey.Input{Message: "Target Workspace"},
	},
}

type surveyConfig struct {
	SourceAPIKey    string `survey:"source_api_key"`
	SourceAPISecret string `survey:"source_api_secret"`
	SourceWorkspace string `survey:"source_workspace"`
	TargetAPIKey    string `survey:"target_api_key"`
	TargetAPISecret string `survey:"target_api_secret"`
	TargetWorkspace string `survey:"target_workspace"`
}

func RunInitConfig(cmd *cobra.Command, args []string) error {
	sConf := surveyConfig{}
	err := survey.Ask(qs, &sConf)
	if err != nil {
		return err
	}

	nonExist := false

	if !config.ConfigExist() {
		err = config.CreateConfigFile()
		if err != nil {
			return err
		}

		nonExist = true
	}

	conf := config.Config{}
	config.ReadConfig(&conf)

	if sConf.SourceAPIKey != "" {
		conf.SourceAPIKey = sConf.SourceAPIKey
	}
	if sConf.SourceAPISecret != "" {
		conf.SourceAPISecret = sConf.SourceAPISecret
	}
	if sConf.SourceWorkspace != "" {
		conf.SourceWorkspace = sConf.SourceWorkspace
	}
	if sConf.TargetAPIKey != "" {
		conf.TargetAPIKey = sConf.TargetAPIKey
	}
	if sConf.TargetAPISecret != "" {
		conf.TargetAPISecret = sConf.TargetAPISecret
	}
	if sConf.TargetWorkspace != "" {
		conf.TargetWorkspace = sConf.TargetWorkspace
	}

	err = config.UpdateConfigFile(&conf)
	if err != nil {
		return err
	}

	path, err := config.GetDefaultConfigPath()
	if err != nil {
		return err
	}

	if nonExist {
		fmt.Printf("Created config file at %s\n", path)
	} else {
		fmt.Printf("Updated config file at %s\n", path)
	}

	return nil
}

func RunCatConfig(cmd *cobra.Command, args []string) error {
	conf, err := config.ReadConfigAsString()
	if err != nil {
		fmt.Println(`config file not found. Run "twilio-copy-studio-flow config init" to create one.`)

		return err
	}

	fmt.Println(conf)

	return nil
}
