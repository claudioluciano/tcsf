package cmd

import (
	"bytes"
	"testing"
)

func TestWorkspaceCmd(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"tr", "workspace"})
	rootCmd.Execute()
}
