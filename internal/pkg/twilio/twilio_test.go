package twilio

import (
	"os"
	"testing"
)

var client *Twilio

func TestMain(m *testing.M) {
	os.Exit(func() int {
		return m.Run()
	}())
}
