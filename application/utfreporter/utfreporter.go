package utfreporter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/chzyer/readline/runes"
	"github.com/spf13/cobra"
	"github.com/storskegg/utf-reporter/runic"
)

func init() {
	cmdRoot.Flags().StringVarP(&flagFile, "file", "f", "", "Input text file")
}

const (
	version = "0.1.0"
)

var flagFile string

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

	// Capture piped input, capturing runes
	reader := bufio.NewReader(input)

	lineNum := 0
	colNum := 0

	var sr runic.SpecialRunes
	var rr runic.Runic

	for {
		// TODO: Prefer reader.ReadLine() over scanner.Scan()?
		input, _, rlErr := reader.ReadLine()
		if rlErr != nil && rlErr == io.EOF {
			break
		}

		if rlErr != nil {
			cmd.SilenceUsage = true
			return rlErr
		}

		lineNum++

		if len(input) == 0 {
			continue
		}

		sr = runic.ProcessLine(string(input))
		if sr == nil {
			continue
		}

		for _, colNum = range sr.SortedColumns() {
			rr, _ = sr.Get(colNum)
			fmt.Println(
				lineNum, "\t",
				colNum, "\t",
				fmt.Sprintf("'%c'", rr), "\t",
				"0x"+rr.CharCodeWithPadding(), "\t",
				rr.RuneType(), "\t",
				runes.Width(rr.Rune()), "\t",
				fmt.Sprintf("https://www.compart.com/en/unicode/U+%s", strings.ToUpper(rr.CharCodeWithPadding())),
			)
		}
	}

	return nil
}
