package cmd

import (
	"fmt"

	"github.com/claudioluciano/tcsf/internal/pkg/twilio"

	"github.com/spf13/cobra"
)

func RunListWorkspace(cmd *cobra.Command, args []string) error {
	twClient := twilio.New()

	ws, err := twClient.ListWorkspace()
	if err != nil {
		return err
	}

	for _, v := range ws {
		fmt.Println(`
SID: `, *v.Sid, `
FriendlyName: `, *v.FriendlyName, ``)
	}

	return nil
}
