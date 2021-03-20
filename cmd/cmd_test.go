package cmd

import (
	"os"
	"testing"
)

// this test doesn't do anything
func TestCmd(t *testing.T) {
	Start(append(os.Args, "start"))
}
