package utfreporter

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.Flags().StringVarP(&flagFile, "file", "f", "", "Input text file")
	cmdRoot.Flags().BoolVarP(&flagDetailed, "verbose", "v", false, "Display additional detailed information")
}

const (
	version = "0.1.0"
)

var (
	flagFile     string
	flagDetailed bool
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

func runRoot(cmd *cobra.Command, args []string) (err error) {
	var info os.FileInfo

	var input *os.File

	// this could be cleaned up a bit
	if flagFile == "" || flagFile == "-" {
		input = os.Stdin
		info, err = input.Stat()
		if err != nil {
			return
		}
	} else {
		flagFile = path.Join(".", flagFile)
		info, err = os.Stat(flagFile)
		if err != nil {
			if err == os.ErrNotExist {
				fmt.Printf("The path '%s' does not exist.\n", flagFile)
				return
			}
			return
		}
		input, err = os.Open(flagFile)
		if err != nil {
			return
		}
	}

	// Check stdin for piped input
	if flagFile == "" {
		if info.Mode()&os.ModeNamedPipe == 0 {
			err = ErrInvalidInputFile
			return
		}
	} else {
		if !info.Mode().IsRegular() {
			err = ErrInvalidInputFile
			return
		}
	}

	if flagDetailed {
		err = processDetailed(input)
	} else {
		err = processBasic(input)
	}
	if err != nil {
		cmd.SilenceUsage = true
		return
	}

	return
}
