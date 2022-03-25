package twilio

// Workspace struct for TaskrouterV1Workspace
type Workspace struct {
	// The string that you assigned to describe the Workspace resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Workspace resource
	URL string `json:"url,omitempty"`
}

// Flow struct for StudioV2Flow
type Flow struct {
	// Description of change made in the revision
	CommitMessage *string `json:"commit_message,omitempty"`
	// JSON representation of flow definition
	Definition *map[string]interface{} `json:"definition,omitempty"`
	// The string that you assigned to describe the Flow
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the resource
	URL string `json:"url,omitempty"`
}

// Workflow struct for TaskrouterV1Workflow
type Workflow struct {
	// A JSON string that contains the Workflow's configuration
	Configuration *string `json:"configuration,omitempty"`
	// The string that you assigned to describe the Workflow resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Workflow resource
	URL string `json:"url,omitempty"`
}

// WorkflowCongfiguration struct for Configuration of a Workflow
type WorkflowCongfiguration struct {
	TaskRouting WorkflowTaskRouting `json:"task_routing,omitempty"`
}

// WorkflowTaskRouting struct for TaskRouting of a Workflow
type WorkflowTaskRouting struct {
	Filters       []*WorkflowFilters     `json:"filters,omitempty"`
	DefaultFilter *WorkflowDefaultFilter `json:"default_filter,omitempty"`
}

// WorkflowFilters struct for Filters of a Workflow
type WorkflowFilters struct {
	FilterFriendlyName string             `json:"filter_friendly_name,omitempty"`
	Expression         *string            `json:"expression,omitempty"`
	Targets            []*WorkflowTargets `json:"targets,omitempty"`
}

// WorkflowDefaultFilter struct for DefaultFilter of a Workflow
type WorkflowDefaultFilter struct {
	Queue string `json:"queue,omitempty"`
}

// WorkflowTargets struct for Targets of a Workflow
type WorkflowTargets struct {
	Queue      *string `json:"queue,omitempty"`
	Priority   *int    `json:"priority,omitempty"`
	Expression *string `json:"expression,omitempty"`
	SkipIf     *string `json:"skip_if,omitempty"`
}

// TaskQueue struct for TaskrouterV1TaskQueue
type TaskQueue struct {
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The maximum number of Workers to reserve
	MaxReservedWorkers *int `json:"max_reserved_workers,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// How Tasks will be assigned to Workers
	TaskOrder *string `json:"task_order,omitempty"`
	// The absolute URL of the TaskQueue resource
	URL string `json:"url,omitempty"`
	// A string describing the Worker selection criteria for any Tasks that enter the TaskQueue
	TargetWorkers *string `json:"target_workers,omitempty"`
}
