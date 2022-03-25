package url

import (
	"fmt"
	"strings"
)

const (
	// base console url
	base = "https://console.twilio.com"
)

// BuildWorkspaceConsoleURL builds the console url for a Workspace
func BuildWorkspaceConsoleURL(sid string) string {
	fr := fmt.Sprintf("frameUrl=console/taskrouter/workspaces/%s/overview", sid)
	// change / to %2F
	fr = strings.Replace(fr, "/", "%2F", -1)
	return fmt.Sprintf("%s/service/taskrouter/%s/taskrouter-workspace-overview?%s", base, sid, fr)
}

// BuildWorkflowConsoleURL builds the console url for a Workflow
func BuildWorkflowConsoleURL(workspaceSid, sid string) string {
	fr := fmt.Sprintf("frameUrl=console/taskrouter/workspaces/%s/workflows/%s", workspaceSid, sid)
	// change / to %2F
	fr = strings.Replace(fr, "/", "%2F", -1)
	return fmt.Sprintf("%s/service/taskrouter/%s/taskrouter-workspace-workflows?%s", base, workspaceSid, fr)
}

// BuildTaskQueueConsoleURL builds the console url for a Taskqueue
func BuildTaskQueueConsoleURL(workspaceSid, sid string) string {
	fr := fmt.Sprintf("frameUrl=console/taskrouter/workspaces/%s/taskqueues/%s", workspaceSid, sid)
	// change / to %2F
	fr = strings.Replace(fr, "/", "%2F", -1)
	return fmt.Sprintf("%s/service/taskrouter/%s/taskrouter-workspace-taskqueues?%s", base, workspaceSid, fr)
}

// BuildFlowConsoleURL builds the console url for a Flow
func BuildFlowConsoleURL(sid string) string {
	fr := fmt.Sprintf("frameUrl=console/studio/flows/%s/Flows", sid)
	// change / to %2F
	fr = strings.Replace(fr, "/", "%2F", -1)
	return fmt.Sprintf("%s/service/studio/%s/studio-flow-instance-canvas?%s", base, sid, fr)
}
