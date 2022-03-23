package twiliosdk

import (
	sdktwilio "github.com/twilio/twilio-go"

	openapiStudio "github.com/twilio/twilio-go/rest/studio/v2"
	openapiTaskrouter "github.com/twilio/twilio-go/rest/taskrouter/v1"
)

type twilioapi struct {
	client *sdktwilio.RestClient
}

type TwilioApi interface {
	ListWorkspace(params *openapiTaskrouter.ListWorkspaceParams) ([]openapiTaskrouter.TaskrouterV1Workspace, error)

	ListFlow(params *openapiStudio.ListFlowParams) ([]openapiStudio.StudioV2Flow, error)
	FetchFlow(sid string) (*openapiStudio.StudioV2Flow, error)
	CreateFlow(params *openapiStudio.CreateFlowParams) (*openapiStudio.StudioV2Flow, error)
	UpdateFlow(sid string, params *openapiStudio.UpdateFlowParams) (*openapiStudio.StudioV2Flow, error)

	ListWorkflow(WorkspaceSid string, params *openapiTaskrouter.ListWorkflowParams) ([]openapiTaskrouter.TaskrouterV1Workflow, error)
	FetchWorkflow(WorkspaceSid string, Sid string) (*openapiTaskrouter.TaskrouterV1Workflow, error)
	CreateWorkflow(WorkspaceSid string, params *openapiTaskrouter.CreateWorkflowParams) (*openapiTaskrouter.TaskrouterV1Workflow, error)
	UpdateWorkflow(WorkspaceSid string, Sid string, params *openapiTaskrouter.UpdateWorkflowParams) (*openapiTaskrouter.TaskrouterV1Workflow, error)

	FetchTaskQueue(WorkspaceSid string, Sid string) (*openapiTaskrouter.TaskrouterV1TaskQueue, error)
	ListTaskQueue(WorkspaceSid string, params *openapiTaskrouter.ListTaskQueueParams) ([]openapiTaskrouter.TaskrouterV1TaskQueue, error)
	CreateTaskQueue(WorkspaceSid string, params *openapiTaskrouter.CreateTaskQueueParams) (*openapiTaskrouter.TaskrouterV1TaskQueue, error)
	UpdateTaskQueue(WorkspaceSid string, Sid string, params *openapiTaskrouter.UpdateTaskQueueParams) (*openapiTaskrouter.TaskrouterV1TaskQueue, error)
}

func New(username, password string) TwilioApi {
	client := sdktwilio.NewRestClientWithParams(sdktwilio.RestClientParams{
		Username: username,
		Password: password,
	})

	return &twilioapi{
		client: client,
	}
}

func (t *twilioapi) ListWorkspace(params *openapiTaskrouter.ListWorkspaceParams) ([]openapiTaskrouter.TaskrouterV1Workspace, error) {
	return t.client.TaskrouterV1.ListWorkspace(params)
}

func (t *twilioapi) ListFlow(params *openapiStudio.ListFlowParams) ([]openapiStudio.StudioV2Flow, error) {
	return t.client.StudioV2.ListFlow(params)
}

func (t *twilioapi) FetchFlow(sid string) (*openapiStudio.StudioV2Flow, error) {
	return t.client.StudioV2.FetchFlow(sid)
}

func (t *twilioapi) CreateFlow(params *openapiStudio.CreateFlowParams) (*openapiStudio.StudioV2Flow, error) {
	return t.client.StudioV2.CreateFlow(params)
}

func (t *twilioapi) UpdateFlow(sid string, params *openapiStudio.UpdateFlowParams) (*openapiStudio.StudioV2Flow, error) {
	return t.client.StudioV2.UpdateFlow(sid, params)
}

func (t *twilioapi) ListWorkflow(WorkspaceSid string, params *openapiTaskrouter.ListWorkflowParams) ([]openapiTaskrouter.TaskrouterV1Workflow, error) {
	return t.client.TaskrouterV1.ListWorkflow(WorkspaceSid, params)
}

func (t *twilioapi) FetchWorkflow(WorkspaceSid string, Sid string) (*openapiTaskrouter.TaskrouterV1Workflow, error) {
	return t.client.TaskrouterV1.FetchWorkflow(WorkspaceSid, Sid)
}

func (t *twilioapi) CreateWorkflow(WorkspaceSid string, params *openapiTaskrouter.CreateWorkflowParams) (*openapiTaskrouter.TaskrouterV1Workflow, error) {
	return t.client.TaskrouterV1.CreateWorkflow(WorkspaceSid, params)
}

func (t *twilioapi) UpdateWorkflow(WorkspaceSid string, Sid string, params *openapiTaskrouter.UpdateWorkflowParams) (*openapiTaskrouter.TaskrouterV1Workflow, error) {
	return t.client.TaskrouterV1.UpdateWorkflow(WorkspaceSid, Sid, params)
}

func (t *twilioapi) FetchTaskQueue(WorkspaceSid string, Sid string) (*openapiTaskrouter.TaskrouterV1TaskQueue, error) {
	return t.client.TaskrouterV1.FetchTaskQueue(WorkspaceSid, Sid)
}

func (t *twilioapi) ListTaskQueue(WorkspaceSid string, params *openapiTaskrouter.ListTaskQueueParams) ([]openapiTaskrouter.TaskrouterV1TaskQueue, error) {
	return t.client.TaskrouterV1.ListTaskQueue(WorkspaceSid, params)
}

func (t *twilioapi) CreateTaskQueue(WorkspaceSid string, params *openapiTaskrouter.CreateTaskQueueParams) (*openapiTaskrouter.TaskrouterV1TaskQueue, error) {
	return t.client.TaskrouterV1.CreateTaskQueue(WorkspaceSid, params)
}

func (t *twilioapi) UpdateTaskQueue(WorkspaceSid string, Sid string, params *openapiTaskrouter.UpdateTaskQueueParams) (*openapiTaskrouter.TaskrouterV1TaskQueue, error) {
	return t.client.TaskrouterV1.UpdateTaskQueue(WorkspaceSid, Sid, params)
}
