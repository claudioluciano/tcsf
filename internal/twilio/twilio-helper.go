package twilio

import (
	"encoding/json"
	"fmt"
	"regexp"

	openapiStudio "github.com/twilio/twilio-go/rest/studio/v2"
)

func GetWorkFlowSIDFromFlow(flow *openapiStudio.StudioV2Flow) (string, error) {
	b, err := json.Marshal(flow.Definition)
	if err != nil {
		fmt.Println("error", err)
		return "", err
	}

	r, _ := regexp.Compile(`"workflow":"([a-zA-Z0-9]+)"`)
	ffs := r.FindStringSubmatch(string(b))
	if len(ffs) < 1 {
		fmt.Println("error flow without workflow, try put a workflow on the send to flex widget: source account")
		return "", err
	}

	return ffs[1], nil
}

func ReplaceWorkFlowSIDFronFlowDefinition(flow *openapiStudio.StudioV2Flow, sid string) (map[string]interface{}, error) {
	b, err := json.Marshal(flow.Definition)
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

func GetTaskQueueIDFromWorkflowConfiguration(configuration string) (map[string]string, error) {
	workFlowCongfiguration, err := WorkFlowConfigurationStringToStruct(configuration)
	if err != nil {
		return nil, err
	}

	var taskQueueID map[string]string
	for _, filters := range workFlowCongfiguration.TaskRouting.Filters {
		for _, targets := range filters.Targets {
			if targets.Queue != nil {
				taskQueueID[*targets.Queue] = filters.FilterFriendlyName
			}
		}
	}

	if workFlowCongfiguration.TaskRouting.DefaultFilter != nil {
		taskQueueID[workFlowCongfiguration.TaskRouting.DefaultFilter.Queue] = "default"
	}

	return taskQueueID, nil
}

func ReplaceTaskQueueSidOnConfiguration(configuration string, sourceTqTargetTq map[string]string) (string, error) {
	workFlowCongfiguration, err := WorkFlowConfigurationStringToStruct(configuration)
	if err != nil {
		return "", err
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
		return "", err
	}

	return string(b), nil
}
