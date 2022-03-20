/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"twilio_copy_studio_flow/internal/twilio"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// studioCmd represents the flow command
var studioCmd = &cobra.Command{
	Use:   "studio",
	Short: "Handles flow tasks.",
	Long: `Flow (flow) helps manipulating an flow inside twilio. It
	provides options for copy from a source and a few more.
	`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println("service:", err)
			return
		}
	},
}

var flowCmd = &cobra.Command{
	Use:     "flow",
	Short:   "List flows",
	Aliases: []string{"f"},
	Run: func(cmd *cobra.Command, args []string) {
		twClient := twilio.New()

		sf, err := twClient.GetStudioFlows()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range sf {
			fmt.Println(`
SID: `, *v.Sid, `
FriendlyName: `, *v.FriendlyName, ``)
		}
	},
}

var copyFlowCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy a flow",
	Run: func(cmd *cobra.Command, args []string) {
		key := viper.GetString("sid")
		if key == "" {
			fmt.Println("'sid' flag is mandatory")
			return
		}

		twClient := twilio.New()

		sf, err := twClient.GetStudioFlows()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range sf {
			fmt.Println(`
SID: `, *v.Sid, `
FriendlyName: `, *v.FriendlyName, ``)
		}
	},
}

func init() {
	rootCmd.AddCommand(studioCmd)
	studioCmd.AddCommand(flowCmd)
	flowCmd.AddCommand(copyFlowCmd)

	copyFlowCmd.Flags().String("sid", "", "SID of the flow to copy")
	_ = viper.BindPFlag("sid", copyFlowCmd.Flags().Lookup("sid"))

	addTargetFlag(studioCmd)
	addTargetFlag(flowCmd)
}
