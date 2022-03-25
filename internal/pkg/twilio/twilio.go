package twilio

import (
	"strings"

	"github.com/claudioluciano/tcsf/internal/pkg/config"
	twiliosdk "github.com/claudioluciano/tcsf/internal/pkg/twilio/twilio-sdk"

	openapiStudio "github.com/twilio/twilio-go/rest/studio/v2"
	openapiTaskrouter "github.com/twilio/twilio-go/rest/taskrouter/v1"
)

type Twilio struct {
	api twiliosdk.TwilioApi
}

type Options struct {
	APIKey    string
	APISecret string
}

func New(forceTarget ...bool) *Twilio {
	apiKey, apiSecret := config.GetRightCredentialsFromConfig(forceTarget...)

	return NewWithOptions(&Options{
		APIKey:    apiKey,
		APISecret: apiSecret,
	})
}

func NewWithOptions(options *Options) *Twilio {
	client := twiliosdk.New(options.APIKey, options.APISecret)

	return &Twilio{
		api: client,
	}
}

func (t *Twilio) ListWorkspace() ([]*Workspace, error) {
	ws, err := t.api.ListWorkspace(nil)
	if err != nil {
		return nil, err
	}

	wss := []*Workspace{}
	for _, v := range ws {
		wss = append(wss, &Workspace{
			Sid:          v.Sid,
			FriendlyName: v.FriendlyName,
			URL:          v.Url,
		})
	}

	return wss, nil
}

func (t *Twilio) ListFlow() ([]*Flow, error) {
	sf, err := t.api.ListFlow(nil)
	if err != nil {
		return nil, err
	}

	sff := []*Flow{}
	for _, v := range sf {
		sff = append(sff, &Flow{
			Sid:           v.Sid,
			FriendlyName:  v.FriendlyName,
			CommitMessage: v.CommitMessage,
			Definition:    v.Definition,
			URL:           v.Url,
		})
	}

	return sff, nil
}

func (t *Twilio) FetchFlow(sid string) (*Flow, error) {
	sf, err := t.api.FetchFlow(sid)
	if err != nil {
		return nil, err
	}

	return &Flow{
		Sid:           sf.Sid,
		FriendlyName:  sf.FriendlyName,
		CommitMessage: sf.CommitMessage,
		Definition:    sf.Definition,
		URL:           sf.Url,
	}, nil
}

func (t *Twilio) FetchFlowByFriendlyName(friendlyName string) (*Flow, error) {
	flows, err := t.ListFlow()
	if err != nil {
		return nil, err
	}

	for _, v := range flows {
		if *v.FriendlyName == friendlyName {
			return v, nil
		}
	}

	return nil, newError(HTTPCodeNotFound, "Flow not found")
}

func (t *Twilio) ListFlowByFriendlyName(friendlyName string) ([]*Flow, error) {
	flows, err := t.ListFlow()
	if err != nil {
		return nil, err
	}

	ff := []*Flow{}
	for _, v := range flows {
		if strings.Contains(*v.FriendlyName, friendlyName) {
			ff = append(ff, v)
		}
	}

	return ff, nil
}

func (t *Twilio) CreateFlow(params *openapiStudio.CreateFlowParams) (*Flow, error) {
	sf, err := t.api.CreateFlow(params)
	if err != nil {
		return nil, err
	}

	return &Flow{
		Sid:           sf.Sid,
		FriendlyName:  sf.FriendlyName,
		CommitMessage: sf.CommitMessage,
		Definition:    sf.Definition,
		URL:           sf.Url,
	}, nil
}

func (t *Twilio) UpdateFlow(sid string, params *openapiStudio.UpdateFlowParams) (*Flow, error) {
	sf, err := t.api.UpdateFlow(sid, params)
	if err != nil {
		return nil, err
	}

	return &Flow{
		Sid:           sf.Sid,
		FriendlyName:  sf.FriendlyName,
		CommitMessage: sf.CommitMessage,
		Definition:    sf.Definition,
		URL:           sf.Url,
	}, nil
}

func (t *Twilio) FetchWorkflow(WorkspaceSid, Sid string) (*Workflow, error) {
	wf, err := t.api.FetchWorkflow(WorkspaceSid, Sid)
	if err != nil {
		return nil, err
	}

	return &Workflow{
		Sid:           wf.Sid,
		FriendlyName:  wf.FriendlyName,
		URL:           wf.Url,
		Configuration: wf.Configuration,
	}, nil
}

func (t *Twilio) ListWorkflow(WorkspaceSid string) ([]Workflow, error) {
	wf, err := t.api.ListWorkflow(WorkspaceSid, nil)
	if err != nil {
		return nil, err
	}

	wff := []Workflow{}
	for _, v := range wf {
		wff = append(wff, Workflow{
			Sid:           v.Sid,
			FriendlyName:  v.FriendlyName,
			URL:           v.Url,
			Configuration: v.Configuration,
		})
	}

	return wff, nil
}

func (t *Twilio) FetchWorkflowByFriendlyName(WorkspaceSid, friendlyName string) (*Workflow, error) {
	wf, err := t.api.ListWorkflow(WorkspaceSid, &openapiTaskrouter.ListWorkflowParams{
		FriendlyName: &friendlyName,
	})
	if err != nil {
		return nil, err
	}

	if len(wf) <= 0 {
		return nil, newError(HTTPCodeNotFound, "Workflow not found")
	}

	return &Workflow{
		Sid:           wf[0].Sid,
		FriendlyName:  wf[0].FriendlyName,
		URL:           wf[0].Url,
		Configuration: wf[0].Configuration,
	}, nil
}

func (t *Twilio) CreateWorkflow(WorkspaceSid string, params *openapiTaskrouter.CreateWorkflowParams) (*Workflow, error) {
	wf, err := t.api.CreateWorkflow(WorkspaceSid, params)
	if err != nil {
		return nil, err
	}

	return &Workflow{
		Sid:           wf.Sid,
		FriendlyName:  wf.FriendlyName,
		URL:           wf.Url,
		Configuration: wf.Configuration,
	}, nil
}

func (t *Twilio) UpdateWorkflow(WorkspaceSid string, Sid string, params *openapiTaskrouter.UpdateWorkflowParams) (*Workflow, error) {
	wf, err := t.api.UpdateWorkflow(WorkspaceSid, Sid, params)
	if err != nil {
		return nil, err
	}

	return &Workflow{
		Sid:           wf.Sid,
		FriendlyName:  wf.FriendlyName,
		URL:           wf.Url,
		Configuration: wf.Configuration,
	}, nil
}

func (t *Twilio) FetchTaskQueue(WorkspaceSid, Sid string) (*TaskQueue, error) {
	tq, err := t.api.FetchTaskQueue(WorkspaceSid, Sid)
	if err != nil {
		return nil, err
	}

	return &TaskQueue{
		Sid:                tq.Sid,
		FriendlyName:       tq.FriendlyName,
		URL:                tq.Url,
		MaxReservedWorkers: tq.MaxReservedWorkers,
		TaskOrder:          tq.TaskOrder,
		TargetWorkers:      tq.TargetWorkers,
	}, nil
}

func (t *Twilio) FetchTaskQueueByFriendlyName(WorkspaceSid, friendlyName string) (*TaskQueue, error) {
	tq, err := t.api.ListTaskQueue(WorkspaceSid, &openapiTaskrouter.ListTaskQueueParams{
		FriendlyName: &friendlyName,
	})
	if err != nil {
		return nil, err
	}

	if len(tq) <= 0 {
		return nil, newError(HTTPCodeNotFound, "TaskQueue not found")
	}

	return &TaskQueue{
		Sid:                tq[0].Sid,
		FriendlyName:       tq[0].FriendlyName,
		URL:                tq[0].Url,
		MaxReservedWorkers: tq[0].MaxReservedWorkers,
		TaskOrder:          tq[0].TaskOrder,
		TargetWorkers:      tq[0].TargetWorkers,
	}, nil
}

func (t *Twilio) CreateTaskQueue(WorkspaceSid string, params *openapiTaskrouter.CreateTaskQueueParams) (*TaskQueue, error) {
	tq, err := t.api.CreateTaskQueue(WorkspaceSid, params)
	if err != nil {
		return nil, err
	}

	return &TaskQueue{
		Sid:                tq.Sid,
		FriendlyName:       tq.FriendlyName,
		URL:                tq.Url,
		MaxReservedWorkers: tq.MaxReservedWorkers,
		TaskOrder:          tq.TaskOrder,
		TargetWorkers:      tq.TargetWorkers,
	}, nil
}

func (t *Twilio) UpdateTaskQueue(WorkspaceSid string, Sid string, params *openapiTaskrouter.UpdateTaskQueueParams) (*TaskQueue, error) {
	tq, err := t.api.UpdateTaskQueue(WorkspaceSid, Sid, params)
	if err != nil {
		return nil, err
	}

	return &TaskQueue{
		Sid:                tq.Sid,
		FriendlyName:       tq.FriendlyName,
		URL:                tq.Url,
		MaxReservedWorkers: tq.MaxReservedWorkers,
		TaskOrder:          tq.TaskOrder,
		TargetWorkers:      tq.TargetWorkers,
	}, nil
}
