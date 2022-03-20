/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"twilio_copy_studio_flow/internal/config"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name:     "source_account_sid",
		Prompt:   &survey.Input{Message: "Source Account SID"},
		Validate: survey.Required,
	},
	{
		Name:     "source_api_key",
		Prompt:   &survey.Input{Message: "Source API Key"},
		Validate: survey.Required,
	},
	{
		Name:     "source_api_secret",
		Prompt:   &survey.Input{Message: "Source API Secret"},
		Validate: survey.Required,
	},
	{
		Name:     "source_workspace",
		Prompt:   &survey.Input{Message: "Source Workspace"},
		Validate: survey.Required,
	},
	{
		Name:     "target_account_sid",
		Prompt:   &survey.Input{Message: "Target Account SID"},
		Validate: survey.Required,
	},
	{
		Name:   "target_api_secret",
		Prompt: &survey.Input{Message: "Target API Secret"},
	},
	{
		Name:   "target_api_key",
		Prompt: &survey.Input{Message: "Target API Key"},
	},
	{
		Name:   "target_workspace",
		Prompt: &survey.Input{Message: "Target Workspace"},
	},
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Setup credentials for Twilio",
	Long: `Setup the credentials to use on Twilio.
Has two sets of credentials Source and Target
If the Target credentials is not set then the source credendials will be used as Target too.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.Config{}
		err := survey.Ask(qs, &conf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = config.CreateConfigFile()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		path, err := config.GetDefaultConfigPath()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = config.UpdateConfigFile(&conf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Config file created at: ", path)
	},
}

var catConfig = &cobra.Command{
	Use:   "cat",
	Short: "View the config file",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := config.ReadConfigFile()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(s)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(catConfig)
}
