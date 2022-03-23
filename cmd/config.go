/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	cmd2 "github.com/claudioluciano/tcsf/internal/cmd"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Handle the configuration",
	Aliases: []string{"c"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}

		return nil
	},
}

var initConfigCmd = &cobra.Command{
	Use:   "init",
	Short: "Handle the initial configuration",
	Long: `Handle the initial configuration of your credentials to use on Twilio.
Has two sets of credentials Source and Target
If the Target credentials is not set then the source credendials will be used as Target too.`,
	Aliases: []string{"i"},
	RunE:    cmd2.RunInitConfig,
}

var catConfigCmd = &cobra.Command{
	Use:   "cat",
	Short: "View the config file",
	RunE:  cmd2.RunCatConfig,
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(initConfigCmd)
	configCmd.AddCommand(catConfigCmd)
}
