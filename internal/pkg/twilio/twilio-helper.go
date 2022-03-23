package twilio

import (
	"encoding/json"
	"regexp"
)

// WorkflowSid get the workflow sid from the flow definition
func (f *Flow) WorkflowSid() (string, error) {
	b, err := json.Marshal(f.Definition)
	if err != nil {
		return "", err
	}

	r, _ := regexp.Compile(`"workflow":"([a-zA-Z0-9]+)"`)
	ffs := r.FindStringSubmatch(string(b))
	if len(ffs) < 1 {
		return "", err
	}

	return ffs[1], nil
}

// SetWorkflowSid set the workflow sid on the flow definition
func (f *Flow) SetWorkflowSid(sid string) (map[string]interface{}, error) {
	b, err := json.Marshal(f.Definition)
	if err != nil {
		return nil, err
	}

	r, _ := regexp.Compile(`"workflow":"([a-zA-Z0-9]+)"`)
	s := r.ReplaceAllString(string(b), `"workflow":"`+sid+`"`)

	var def map[string]interface{}
	err = json.Unmarshal([]byte(s), &def)
	if err != nil {
		return nil, err
	}
	return def, nil
}

// ConfigurationAsStruct convert the configuration string to struct
func (w *Workflow) ConfigurationAsStruct() (*WorkflowCongfiguration, error) {
	var workFlowCongfiguration WorkflowCongfiguration
	if err := json.Unmarshal([]byte(*w.Configuration), &workFlowCongfiguration); err != nil {
		return nil, err
	}

	return &workFlowCongfiguration, nil
}

func (w *Workflow) GetTaskQueuesSidFromConfiguration() ([]string, error) {
	workFlowCongfiguration, err := w.ConfigurationAsStruct()
	if err != nil {
		return nil, err
	}

	taskQueueID := []string{}
	for _, filters := range workFlowCongfiguration.TaskRouting.Filters {
		for _, targets := range filters.Targets {
			if targets.Queue != nil {
				taskQueueID = append(taskQueueID, *targets.Queue)
			}
		}
	}

	if workFlowCongfiguration.TaskRouting.DefaultFilter != nil {
		taskQueueID = append(taskQueueID, workFlowCongfiguration.TaskRouting.DefaultFilter.Queue)
	}

	return taskQueueID, nil
}

// ReplaceTaskQueueSidOnConfiguration replace the task queues sid on the workflow configuration
func (w *Workflow) ReplaceTaskQueueSidOnConfiguration(sourceTqTargetTq map[string]string) error {
	workFlowCongfiguration, err := w.ConfigurationAsStruct()
	if err != nil {
		return err
	}

	for _, filters := range workFlowCongfiguration.TaskRouting.Filters {
		for _, targets := range filters.Targets {
			if targets.Queue != nil {
				if _, ok := sourceTqTargetTq[*targets.Queue]; ok {
					sid := sourceTqTargetTq[*targets.Queue]
					targets.Queue = &sid
				}
			}
		}
	}

	if workFlowCongfiguration.TaskRouting.DefaultFilter != nil {
		if _, ok := sourceTqTargetTq[workFlowCongfiguration.TaskRouting.DefaultFilter.Queue]; ok {
			workFlowCongfiguration.TaskRouting.DefaultFilter.Queue = sourceTqTargetTq[workFlowCongfiguration.TaskRouting.DefaultFilter.Queue]
		}
	}

	b, err := json.Marshal(workFlowCongfiguration)
	if err != nil {
		return err
	}

	s := string(b)
	w.Configuration = &s

	return nil
}

// Equal compare two TaskQueue
func (t *TaskQueue) Equal(tq *TaskQueue) bool {
	return t.TaskOrder == tq.TaskOrder && t.MaxReservedWorkers == tq.MaxReservedWorkers && t.TargetWorkers == tq.TargetWorkers
}
