package cmd

import (
	"fmt"

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
FriendlyName: `, *v.FriendlyName, ``)
	}

	return nil
}
