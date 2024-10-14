package utfreporter

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	version = "0.1.0"
)

var cmdRoot = &cobra.Command{
	Use:     "utf-reporter",
	Short:   "Find and report unicode and non-standard ascii characters in a text file or stream.",
	Args:    nil,
	Version: version,
	RunE:    runRoot,
}

func Execute() error {
	return cmdRoot.Execute()
}

func runRoot(cmd *cobra.Command, args []string) error {
	return errors.New("not implemented yet")
}
