package twilio

import "encoding/json"

type WorkFlowCongfiguration struct {
	TaskRouting TaskRouting `json:"task_routing,omitempty"`
}

type TaskRouting struct {
	Filters       []*Filters     `json:"filters,omitempty"`
	DefaultFilter *DefaultFilter `json:"default_filter,omitempty"`
}

type Filters struct {
	FilterFriendlyName string     `json:"filter_friendly_name,omitempty"`
	Expression         *string    `json:"expression,omitempty"`
	Targets            []*Targets `json:"targets,omitempty"`
}

type DefaultFilter struct {
	Queue string `json:"queue,omitempty"`
}

type Targets struct {
	Queue      *string `json:"queue,omitempty"`
	Priority   *int    `json:"priority,omitempty"`
	Expression *string `json:"expression,omitempty"`
	SkipIf     *string `json:"skip_if,omitempty"`
}

func WorkFlowConfigurationStringToStruct(configuration string) (*WorkFlowCongfiguration, error) {
	var workFlowCongfiguration WorkFlowCongfiguration
	if err := json.Unmarshal([]byte(configuration), &workFlowCongfiguration); err != nil {
		return nil, err
	}

	return &workFlowCongfiguration, nil
}