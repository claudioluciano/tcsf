package twilio

import (
	"twilio_copy_studio_flow/internal/config"

	"github.com/spf13/viper"
	"github.com/twilio/twilio-go"

	openapiStudio "github.com/twilio/twilio-go/rest/studio/v2"
	openapiTaskrouter "github.com/twilio/twilio-go/rest/taskrouter/v1"
)

type Twilio struct {
	client *twilio.RestClient
}

type Options struct {
	APIKey    string
	APISecret string
}

func New(targetOptions ...bool) *Twilio {
	target := viper.GetBool("target")
	if len(targetOptions) > 0 {
		target = targetOptions[0]
	}

	config := config.GetConfigFromViper()

	apiKey := config.SourceAPIKey
	apiSecret := config.SourceAPISecret
	if target {
		apiKey = config.TargetAPIKey
		apiSecret = config.TargetAPISecret
	}

	return NewWithOptions(&Options{
		APIKey:    apiKey,
		APISecret: apiSecret,
	})
}

func NewWithOptions(opt *Options) *Twilio {
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: opt.APIKey,
		Password: opt.APISecret,
	})

	return &Twilio{
		client: client,
	}
}

func (t *Twilio) GetWorkspaces() ([]openapiTaskrouter.TaskrouterV1Workspace, error) {
	ws, err := t.client.TaskrouterV1.ListWorkspace(nil)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (t *Twilio) GetStudioFlows() ([]openapiStudio.StudioV2Flow, error) {
	sf, err := t.client.StudioV2.ListFlow(nil)
	if err != nil {
		return nil, err
	}

	return sf, nil
}

func (t *Twilio) GetStudioFlow(sid string) (*openapiStudio.StudioV2Flow, error) {
	sf, err := t.client.StudioV2.FetchFlow(sid)
	if err != nil {
		return nil, err
	}

	return sf, nil
}

func (t *Twilio) GetStudioFlowswByFriendlyName(friendlyName string) (*openapiStudio.StudioV2Flow, error) {
	flows, err := t.GetStudioFlows()
	if err != nil {
		return nil, err
	}

	for _, v := range flows {
		if *v.FriendlyName == friendlyName {
			return &v, nil
		}
	}

	return nil, nil
}

func (t *Twilio) CreateStudioFlow(params *openapiStudio.CreateFlowParams) (*openapiStudio.StudioV2Flow, error) {
	return t.client.StudioV2.CreateFlow(params)
}

func (t *Twilio) UpdateStudioFlow(flowSid string, params *openapiStudio.UpdateFlowParams) (*openapiStudio.StudioV2Flow, error) {
	return t.client.StudioV2.UpdateFlow(flowSid, params)
}

func (t *Twilio) GetWorkflow(workspaceSID, workFlowSid string) (*openapiTaskrouter.TaskrouterV1Workflow, error) {
	wf, err := t.client.TaskrouterV1.FetchWorkflow(workspaceSID, workFlowSid)
	if err != nil {
		return nil, err
	}

	return wf, nil
}

func (t *Twilio) GetWorkflows(workspaceSID, workFlowSid string) ([]openapiTaskrouter.TaskrouterV1Workflow, error) {
	wf, err := t.client.TaskrouterV1.ListWorkflow(workspaceSID, nil)
	if err != nil {
		return nil, err
	}

	return wf, nil
}

func (t *Twilio) GetWorkflowByFriendlyName(workspaceSID, friendlyName string) (*openapiTaskrouter.TaskrouterV1Workflow, error) {
	wf, err := t.client.TaskrouterV1.ListWorkflow(workspaceSID, &openapiTaskrouter.ListWorkflowParams{
		FriendlyName: &friendlyName,
	})
	if err != nil {
		return nil, err
	}

	return &wf[0], nil
}

func (t *Twilio) CreateWorkflow(workspaceSID string, params *openapiTaskrouter.CreateWorkflowParams) (*openapiTaskrouter.TaskrouterV1Workflow, error) {
	tq, err := t.client.TaskrouterV1.CreateWorkflow(workspaceSID, params)
	if err != nil {
		return nil, err
	}

	return tq, nil
}

func (t *Twilio) UpdateWorkflow(workspaceSID string, workflowSID string, params *openapiTaskrouter.UpdateWorkflowParams) (*openapiTaskrouter.TaskrouterV1Workflow, error) {
	tq, err := t.client.TaskrouterV1.UpdateWorkflow(workspaceSID, workflowSID, params)
	if err != nil {
		return nil, err
	}

	return tq, nil
}

func (t *Twilio) GetTaskQueue(workspaceSID, taskQueueSID string) (*openapiTaskrouter.TaskrouterV1TaskQueue, error) {
	tq, err := t.client.TaskrouterV1.FetchTaskQueue(workspaceSID, taskQueueSID)
	if err != nil {
		return nil, err
	}

	return tq, nil
}

func (t *Twilio) GetTaskQueueByFriendlyName(workspaceSID, friendlyName string) (*openapiTaskrouter.TaskrouterV1TaskQueue, error) {
	tq, err := t.client.TaskrouterV1.ListTaskQueue(workspaceSID, &openapiTaskrouter.ListTaskQueueParams{
		FriendlyName: &friendlyName,
	})
	if err != nil {
		return nil, err
	}

	return &tq[0], nil
}

func (t *Twilio) CreateTaskQueue(workspaceSID string, params *openapiTaskrouter.CreateTaskQueueParams) (*openapiTaskrouter.TaskrouterV1TaskQueue, error) {
	tq, err := t.client.TaskrouterV1.CreateTaskQueue(workspaceSID, params)
	if err != nil {
		return nil, err
	}

	return tq, nil
}
