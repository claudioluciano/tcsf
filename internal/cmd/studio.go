package cmd

import (
	"fmt"

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

// RunListFlow lists all flows
func RunListFlow(cmd *cobra.Command, args []string) (err error) {
	target := viper.GetBool("target")
	var (
		twClient = twilio.New(target)
		name     = viper.GetString("name")
		flows    []*twilio.Flow
	)

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

// RunCopyFlow copies a flow
func RunCopyFlow(cmd *cobra.Command, args []string) error {
	// Here we get the flag target to know if we need to invert the credentials
	target := viper.GetBool("target")
	// If target is true, we will use the credentials from the target as the source otherwise we will use the credentials from the source

	var (
		sflowSid           = viper.GetString("sid")
		sourceTwilioClient = twilio.New(target)
		targetTwilioClient = twilio.New(!target)
		cfg                = config.GetConfigFromViper(target)
	)

	// Get the source flow
	sourceStudioFlow, err := sourceTwilioClient.FetchFlow(sflowSid)
	if err != nil {
		if !twilio.NotFound(err) {
			fmt.Println("Not found flow with sid: ", sflowSid)
		}

		return err
	}

	// Get the workflowsid from the source flow
	sourceWorkflowSid, err := sourceStudioFlow.WorkflowSid()
	if err != nil {
		m := "source"
		if target {
			m = "target"
		}

		fmt.Println(fmt.Sprintf("Flow without workflow, try put a workflow on the send_to_flex widget: %s account", m))
		return err
	}

	// Get the source workflow
	sourceWorkflow, err := sourceTwilioClient.FetchWorkflow(cfg.SourceWorkspace, sourceWorkflowSid)
	if err != nil {
		if !twilio.NotFound(err) {
			fmt.Println("Not found workflow with sid: ", sourceWorkflowSid)
		}

		return err
	}

	// Get the source workflow's tasks queue
	sourceTaskQueuesIds, err := sourceWorkflow.GetTaskQueuesSidFromConfiguration()
	if err != nil {
		return err
	}

	// Here we map the source task queues to the target task queues
	// if the target task queue does not exist, we create it otherwise we will verify if it's necessary to update the target task queue
	// and if it's necessary, we update it

	sourceTqTargetTq := make(map[string]string)
	for _, v := range sourceTaskQueuesIds {
		if _, ok := sourceTqTargetTq[v]; ok {
			continue
		}

		// Get the source task queue
		sourceTq, err := sourceTwilioClient.FetchTaskQueue(cfg.SourceWorkspace, v)
		if err != nil {
			if !twilio.NotFound(err) {
				fmt.Println("Not found task queue with sid: ", v)
			}
			return err
		}

		// Get the target task queue
		targetTq, err := targetTwilioClient.FetchTaskQueueByFriendlyName(cfg.TargetWorkspace, *sourceTq.FriendlyName)
		if err != nil {
			// If the target task queue does not exist, we create it
			if !twilio.NotFound(err) {
				fmt.Println("Not found task queue with friendly name: ", *sourceTq.FriendlyName)
				return err
			}

			// Create the target task queue
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

		// Verify if the target task queue needs to be updated
		if !targetTq.Equal(sourceTq) {
			// Update the target task queue
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

		// Map the source task queue to the target task queue
		sourceTqTargetTq[v] = *targetTq.Sid
	}

	// Replace the source task queues sids with the target task queues sids
	sourceWorkflow.ReplaceTaskQueueSidOnConfiguration(sourceTqTargetTq)
	if err != nil {
		return err
	}

	// Get the target workflow
	targetWorkflow, err := targetTwilioClient.FetchWorkflowByFriendlyName(cfg.TargetWorkspace, *sourceWorkflow.FriendlyName)
	if err != nil {
		// If the target workflow does not exist, we create it
		if !twilio.NotFound(err) {
			fmt.Println("Not found workflow with friendly name: ", *sourceWorkflow.FriendlyName)
			return err
		}

		// Create the target workflow
		targetWorkflow, err = targetTwilioClient.CreateWorkflow(cfg.TargetWorkspace, &openapiTaskrouter.CreateWorkflowParams{
			FriendlyName:  sourceWorkflow.FriendlyName,
			Configuration: sourceWorkflow.Configuration,
		})
		if err != nil {
			return err
		}
	}

	// Verify if the target workflow needs to be updated
	if targetWorkflow.Configuration != sourceWorkflow.Configuration {
		// Update the target workflow
		targetWorkflow, err = targetTwilioClient.UpdateWorkflow(cfg.TargetWorkspace, *targetWorkflow.Sid, &openapiTaskrouter.UpdateWorkflowParams{
			Configuration: sourceWorkflow.Configuration,
		})
		if err != nil {
			return err
		}
	}

	// Set the target workflow sid on the source flow
	def, err := sourceStudioFlow.SetWorkflowSid(*targetWorkflow.Sid)
	if err != nil {
		return err
	}

	// Fetch the target flow
	targetStudioFlow, err := targetTwilioClient.FetchFlowByFriendlyName(*sourceStudioFlow.FriendlyName)
	if err != nil {
		// If the target flow does not exist, we create it
		if !twilio.NotFound(err) {
			fmt.Println("Not found flow with friendly name: ", *sourceStudioFlow.FriendlyName)
			return err
		}

		// Create the target flow
		targetStudioFlow, err = targetTwilioClient.CreateFlow(&openapiStudio.CreateFlowParams{
			FriendlyName:  sourceStudioFlow.FriendlyName,
			CommitMessage: sourceStudioFlow.CommitMessage,
			Status:        &FlowStatus,
			Definition:    &def,
		})
		if err != nil {
			return err
		}

		fmt.Println(`
Flow Created with success:
SID: `, *targetStudioFlow.Sid, `
FriendlyName: `, *targetStudioFlow.FriendlyName, `
URL: `, *targetStudioFlow.URL, ``)

		return nil
	}

	// Update the target flow
	targetStudioFlow, err = targetTwilioClient.UpdateFlow(*targetStudioFlow.Sid, &openapiStudio.UpdateFlowParams{
		FriendlyName:  sourceStudioFlow.FriendlyName,
		CommitMessage: sourceStudioFlow.CommitMessage,
		Status:        &FlowStatus,
		Definition:    &def,
	})

	fmt.Println(`
Flow Updated with success:
SID: `, *targetStudioFlow.Sid, `
FriendlyName: `, *targetStudioFlow.FriendlyName, `
URL: `, *targetStudioFlow.URL, ``)

	return nil
}
