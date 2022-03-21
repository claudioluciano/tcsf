/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"
	"twilio_copy_studio_flow/internal/config"
	"twilio_copy_studio_flow/internal/twilio"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	openapiStudio "github.com/twilio/twilio-go/rest/studio/v2"
	openapiTaskrouter "github.com/twilio/twilio-go/rest/taskrouter/v1"
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

		sourceStudioFlows, err := twClient.GetStudioFlows()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range sourceStudioFlows {
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
		sflowSid := viper.GetString("sid")
		if sflowSid == "" {
			fmt.Println("'sid' flag is mandatory")
			return
		}

		var (
			sourceTwilioClient *twilio.Twilio
			targetTwilioClient *twilio.Twilio
			cfg                *config.Config
		)
		sourceTwilioClient = twilio.New()
		targetTwilioClient = twilio.New(true)

		cfg = config.GetConfigFromViper()

		sourceStudioFlow, err := sourceTwilioClient.GetStudioFlow(sflowSid)
		if err != nil {
			fmt.Println(err)
			return
		}

		sourceWorkflowSid, err := twilio.GetWorkFlowSIDFromFlow(sourceStudioFlow)
		if err != nil {
			fmt.Println(err)
			return
		}

		sourceWorkflow, err := sourceTwilioClient.GetWorkflow(cfg.SourceWorkspace, sourceWorkflowSid)
		if err != nil {
			fmt.Println(err)
			return
		}

		sourceTaskQueuesIds, err := twilio.GetTaskQueueIDFromWorkflowConfiguration(*sourceWorkflow.Configuration)
		if err != nil {
			fmt.Println(err)
			return
		}

		sourceTqTargetTq := make(map[string]string)
		for k, v := range sourceTaskQueuesIds {
			if v != "default" {
				tq, err := targetTwilioClient.GetTaskQueueByFriendlyName(cfg.SourceWorkspace, v)
				if err == nil {
					sourceTqTargetTq[k] = *tq.Sid
					continue
				}

				if !strings.Contains(err.Error(), "not found") {
					fmt.Println(err)
					return
				}
			}

			sourceTq, err := sourceTwilioClient.GetTaskQueue(cfg.SourceWorkspace, k)
			if err != nil {
				fmt.Println(err)
				return
			}

			targetTq, err := targetTwilioClient.CreateTaskQueue(cfg.TargetWorkspace, &openapiTaskrouter.CreateTaskQueueParams{
				FriendlyName:  sourceTq.FriendlyName,
				TargetWorkers: sourceTq.TargetWorkers,
				TaskOrder:     sourceTq.TaskOrder,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
			sourceTqTargetTq[k] = *targetTq.Sid
		}

		targetWorkflowConfiguration, err := twilio.ReplaceTaskQueueSidOnConfiguration(*sourceWorkflow.Configuration, sourceTqTargetTq)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Get target workflow
		targetWorkflow, err := targetTwilioClient.GetWorkflowByFriendlyName(cfg.TargetWorkspace, *sourceStudioFlow.FriendlyName)
		if err != nil {
			if !strings.Contains(err.Error(), "could not retrieve payload from response") {
				fmt.Println(err)
				return
			}
		}

		if targetWorkflow == nil {
			targetWorkflow, err = targetTwilioClient.CreateWorkflow(cfg.TargetWorkspace, &openapiTaskrouter.CreateWorkflowParams{
				FriendlyName:                  sourceWorkflow.FriendlyName,
				AssignmentCallbackUrl:         sourceWorkflow.AssignmentCallbackUrl,
				Configuration:                 &targetWorkflowConfiguration,
				FallbackAssignmentCallbackUrl: sourceWorkflow.FallbackAssignmentCallbackUrl,
				TaskReservationTimeout:        sourceWorkflow.TaskReservationTimeout,
			})
		} else {
			targetWorkflow, err = targetTwilioClient.UpdateWorkflow(cfg.TargetWorkspace, *targetWorkflow.Sid, &openapiTaskrouter.UpdateWorkflowParams{
				FriendlyName:                  sourceWorkflow.FriendlyName,
				AssignmentCallbackUrl:         sourceWorkflow.AssignmentCallbackUrl,
				Configuration:                 &targetWorkflowConfiguration,
				FallbackAssignmentCallbackUrl: sourceWorkflow.FallbackAssignmentCallbackUrl,
				TaskReservationTimeout:        sourceWorkflow.TaskReservationTimeout,
			})
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		targetStudioFlow, err := targetTwilioClient.GetStudioFlowswByFriendlyName(*sourceStudioFlow.FriendlyName)
		if err != nil {
			fmt.Println(err)
			return
		}

		def, err := twilio.ReplaceWorkFlowSIDFronFlowDefinition(sourceStudioFlow, *targetWorkflow.Sid)
		if err != nil {
			fmt.Println(err)
			return
		}

		if targetStudioFlow != nil {
			targetStudioFlow, err = targetTwilioClient.UpdateStudioFlow(*targetStudioFlow.Sid, &openapiStudio.UpdateFlowParams{
				FriendlyName:  sourceStudioFlow.FriendlyName,
				CommitMessage: sourceStudioFlow.CommitMessage,
				Status:        sourceStudioFlow.Status,
				Definition:    &def,
			})
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		targetStudioFlow, err = targetTwilioClient.CreateStudioFlow(&openapiStudio.CreateFlowParams{
			FriendlyName:  sourceStudioFlow.FriendlyName,
			CommitMessage: sourceStudioFlow.CommitMessage,
			Status:        sourceStudioFlow.Status,
			Definition:    &def,
		})
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
