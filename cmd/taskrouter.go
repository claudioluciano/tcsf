package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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

var workflowCmd = &cobra.Command{
	Use:     "workflow",
	Short:   "Handle taskrouter workflow tasks",
	Aliases: []string{"wf"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}

		return nil
	},
}

var listWorkflowCmd = &cobra.Command{
	Use:     "list",
	Short:   "List workflow",
	Aliases: []string{"ls"},
	RunE:    cmd2.RunListWorkflow,
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
	taskrouterCmd.AddCommand(workflowCmd)
	workflowCmd.AddCommand(listWorkflowCmd)

	listWorkflowCmd.Flags().String("name", "", "Name of the flow to search")
	_ = viper.BindPFlag("name", listWorkflowCmd.Flags().Lookup("name"))

	taskrouterCmd.AddCommand(workspaceCmd)
	workspaceCmd.AddCommand(listWorkspaceCmd)
}
