/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"twilio_copy_studio_flow/internal/twilio"

	"github.com/spf13/cobra"
)

// taskrouterCmd represents the taskrouter command
var taskrouterCmd = &cobra.Command{
	Use:     "taskrouter",
	Short:   "Handle taskrouter task's",
	Long:    `Handle taskrouter routine tasks`,
	Aliases: []string{"t"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println("taskrouter:", err)
			return
		}
	},
}

var workspaceCmd = &cobra.Command{
	Use:     "workspace",
	Short:   "List workspaces",
	Aliases: []string{"w"},
	Run: func(cmd *cobra.Command, args []string) {
		twClient := twilio.New()

		ws, err := twClient.GetWorkspaces()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range ws {
			fmt.Println(`
SID: `, *v.Sid, `
FriendlyName: `, *v.FriendlyName, ``)
		}
	},
}

func init() {
	rootCmd.AddCommand(taskrouterCmd)
	taskrouterCmd.AddCommand(workspaceCmd)

	addTargetFlag(taskrouterCmd)
}
