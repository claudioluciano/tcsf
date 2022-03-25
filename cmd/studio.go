package cmd

import (
	cmd2 "github.com/claudioluciano/tcsf/internal/cmd"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// FlowStatus represents the status of all flow will be copied
var (
	FlowStatus string = "draft"
)

// studioCmd represents the flow command
var studioCmd = &cobra.Command{
	Use:   "studio",
	Short: "Handles flow tasks.",
	Long: `Flow (flow) helps manipulating an flow inside twilio. It
	provides options for copy from a source and a few more.
	`,
	Aliases: []string{"s"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}

		return nil
	},
}

var flowCmd = &cobra.Command{
	Use:     "flow",
	Short:   "Handle studio flow tasks",
	Aliases: []string{"f"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}

		return nil
	},
}

var listflowCmd = &cobra.Command{
	Use:     "list",
	Short:   "List flows",
	Aliases: []string{"ls"},
	RunE:    cmd2.RunListFlow,
}

var copyFlowCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy a flow",
	RunE:  cmd2.RunCopyFlow,
}

func init() {
	rootCmd.AddCommand(studioCmd)
	studioCmd.AddCommand(flowCmd)

	flowCmd.AddCommand(listflowCmd)
	flowCmd.AddCommand(copyFlowCmd)

	listflowCmd.Flags().String("name", "", "Name of the flow to search")
	_ = viper.BindPFlag("name", listflowCmd.Flags().Lookup("name"))

	copyFlowCmd.Flags().String("sid", "", "SID of the flow to copy")
	copyFlowCmd.MarkFlagRequired("sid")
	_ = viper.BindPFlag("sid", copyFlowCmd.Flags().Lookup("sid"))
}
