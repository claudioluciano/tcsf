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

func ReplaceWorkFlowSID(flow *openapiStudio.StudioV2Flow, sid string) (string, error) {
	b, err := json.Marshal(flow.Definition)
	if err != nil {
		fmt.Println("error", err)
		return "", err
	}

	r, _ := regexp.Compile(`"workflow":"([a-zA-Z0-9]+)"`)
	s := r.ReplaceAllString(string(b), `"workflow":"`+sid+`"`)

	return s, nil
}
