/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	cmd2 "github.com/claudioluciano/tcsf/internal/cmd"
)

// taskrouterCmd represents the taskrouter command
var taskrouterCmd = &cobra.Command{
	Use:     "taskrouter",
	Short:   "Handle taskrouter tasks",
	Aliases: []string{"t"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}

		return nil
	},
}

var workspaceCmd = &cobra.Command{
	Use:     "workspace",
	Short:   "Handle taskrouter workspace tasks",
	Aliases: []string{"w"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}

		return nil
	},
}

var listWorkspaceCmd = &cobra.Command{
	Use:     "list",
	Short:   "List workspaces",
	Aliases: []string{"ls"},
	RunE:    cmd2.RunListWorkspace,
}

func init() {
	rootCmd.AddCommand(taskrouterCmd)
	taskrouterCmd.AddCommand(workspaceCmd)
	workspaceCmd.AddCommand(listWorkspaceCmd)

	addTargetFlag(listWorkspaceCmd)
}
