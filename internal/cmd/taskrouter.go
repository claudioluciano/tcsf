package cmd

import (
	"fmt"

	"github.com/claudioluciano/tcsf/internal/pkg/config"
	"github.com/claudioluciano/tcsf/internal/pkg/twilio"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RunListWorkspace lists all workspaces
func RunListWorkspace(cmd *cobra.Command, args []string) error {
	target := viper.GetBool("target")
	twClient := twilio.New(target)

	ws, err := twClient.ListWorkspace()
	if err != nil {
		return err
	}

	for _, v := range ws {
		fmt.Println(`
SID: `, *v.Sid, `
FriendlyName: `, *v.FriendlyName, `
URL: `, *v.URL, ``)
	}

	return nil
}

// RunListWorkflow lists all Workflows
func RunListWorkflow(cmd *cobra.Command, args []string) error {
	var (
		target    = viper.GetBool("target")
		twClient  = twilio.New(target)
		name      = viper.GetString("name")
		workflows []*twilio.Workflow
		cfg       = config.GetConfigFromViper(target)
	)

	var err error
	if name != "" {
		workflows, err = twClient.ListWorkFlowByFriendlyName(cfg.SourceWorkspace, name)
	} else {
		workflows, err = twClient.ListWorkflow(cfg.SourceWorkspace)
	}
	if err != nil {
		return err
	}

	for _, v := range workflows {
		fmt.Println(`
SID: `, *v.Sid, `
FriendlyName: `, *v.FriendlyName, `
URL: `, *v.URL, ``)
	}

	return nil
}
