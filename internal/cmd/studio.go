package cmd

import (
	"fmt"
	"strings"

	"github.com/claudioluciano/tcsf/internal/pkg/config"
	"github.com/claudioluciano/tcsf/internal/pkg/twilio"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	openapiStudio "github.com/twilio/twilio-go/rest/studio/v2"
	openapiTaskrouter "github.com/twilio/twilio-go/rest/taskrouter/v1"
)

// FlowStatus represents the status of all flow will be copied
var (
	FlowStatus string = "draft"
)

func RunListFlow(cmd *cobra.Command, args []string) (err error) {
	name := viper.GetString("name")

	twClient := twilio.New()

	var flows []*twilio.Flow

	if name != "" {
		flows, err = twClient.ListFlowByFriendlyName(name)
	} else {
		flows, err = twClient.ListFlow()
	}

	if err != nil {
		return err
	}

	for _, v := range flows {
		fmt.Println(`
SID: `, *v.Sid, `
FriendlyName: `, *v.FriendlyName, `
URL: `, *v.URL, ``)
	}

	return nil
}

func RunCopyFlow(cmd *cobra.Command, args []string) error {
	var (
		sflowSid           = viper.GetString("sid")
		sourceTwilioClient *twilio.Twilio
		targetTwilioClient *twilio.Twilio
		cfg                *config.Config
	)
	sourceTwilioClient = twilio.New()
	targetTwilioClient = twilio.New(true)

	cfg = config.GetConfigFromViper()

	sourceStudioFlow, err := sourceTwilioClient.FetchFlow(sflowSid)
	if err != nil {
		return err
	}

	sourceWorkflowSid, err := sourceStudioFlow.WorkflowSid()
	if err != nil {
		return err
	}

	sourceWorkflow, err := sourceTwilioClient.FetchWorkflow(cfg.SourceWorkspace, sourceWorkflowSid)
	if err != nil {
		return err
	}

	sourceTaskQueuesIds, err := sourceWorkflow.GetTaskQueuesSidFromConfiguration()
	if err != nil {
		return err
	}

	sourceTqTargetTq := make(map[string]string)
	for _, v := range sourceTaskQueuesIds {
		if _, ok := sourceTqTargetTq[v]; ok {
			continue
		}

		sourceTq, err := sourceTwilioClient.FetchTaskQueue(cfg.SourceWorkspace, v)
		if err != nil {
			return err
		}

		targetTq, err := targetTwilioClient.FetchTaskQueueByFriendlyName(cfg.TargetWorkspace, *sourceTq.FriendlyName)
		if err != nil {
			if !strings.Contains(err.Error(), "not found") {
				return err
			}

			targetTq, err = targetTwilioClient.CreateTaskQueue(cfg.TargetWorkspace, &openapiTaskrouter.CreateTaskQueueParams{
				FriendlyName:       sourceTq.FriendlyName,
				TaskOrder:          sourceTq.TaskOrder,
				TargetWorkers:      sourceTq.TargetWorkers,
				MaxReservedWorkers: sourceTq.MaxReservedWorkers,
			})
			if err != nil {
				return err
			}
		}

		if !targetTq.Equal(sourceTq) {
			targetTq, err = targetTwilioClient.UpdateTaskQueue(cfg.TargetWorkspace, *targetTq.Sid, &openapiTaskrouter.UpdateTaskQueueParams{
				TargetWorkers:      sourceTq.TargetWorkers,
				MaxReservedWorkers: sourceTq.MaxReservedWorkers,
				TaskOrder:          sourceTq.TaskOrder,
				FriendlyName:       sourceTq.FriendlyName,
			})
			if err != nil {
				return err
			}
		}

		sourceTqTargetTq[v] = *targetTq.Sid
	}

	sourceWorkflow.ReplaceTaskQueueSidOnConfiguration(sourceTqTargetTq)
	if err != nil {
		return err
	}

	// Get target workflow
	targetWorkflow, err := targetTwilioClient.FetchWorkflowByFriendlyName(cfg.TargetWorkspace, *sourceWorkflow.FriendlyName)
	if err != nil {
		if !strings.Contains(err.Error(), "could not retrieve payload from response") {
			return err
		}

		targetWorkflow, err = targetTwilioClient.CreateWorkflow(cfg.TargetWorkspace, &openapiTaskrouter.CreateWorkflowParams{
			FriendlyName:  sourceWorkflow.FriendlyName,
			Configuration: sourceWorkflow.Configuration,
		})
		if err != nil {
			return err
		}
	}

	if targetWorkflow.Configuration != sourceWorkflow.Configuration {
		targetWorkflow, err = targetTwilioClient.UpdateWorkflow(cfg.TargetWorkspace, *targetWorkflow.Sid, &openapiTaskrouter.UpdateWorkflowParams{
			Configuration: sourceWorkflow.Configuration,
		})
		if err != nil {
			return err
		}
	}

	def, err := sourceStudioFlow.SetWorkflowSid(*targetWorkflow.Sid)
	if err != nil {
		return err
	}

	targetStudioFlow, err := targetTwilioClient.FetchFlowByFriendlyName(*sourceStudioFlow.FriendlyName)
	if err != nil {
		if !strings.Contains(err.Error(), "could not retrieve payload from response") {
			return err
		}
		targetStudioFlow, err = targetTwilioClient.CreateFlow(&openapiStudio.CreateFlowParams{
			FriendlyName:  sourceStudioFlow.FriendlyName,
			CommitMessage: sourceStudioFlow.CommitMessage,
			Status:        &FlowStatus,
			Definition:    &def,
		})
		if err != nil {
			return err
		}

		return nil
	}

	targetStudioFlow, err = targetTwilioClient.UpdateFlow(*targetStudioFlow.Sid, &openapiStudio.UpdateFlowParams{
		FriendlyName:  sourceStudioFlow.FriendlyName,
		CommitMessage: sourceStudioFlow.CommitMessage,
		Status:        &FlowStatus,
		Definition:    &def,
	})

	return nil
}
