package cmd

import (
	"fmt"
	"twilio_copy_studio_flow/internal/pkg/twilio"

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
