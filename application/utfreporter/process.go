package utfreporter

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/chzyer/readline/runes"
	"github.com/storskegg/utf-reporter/runic"
)

func processBasic(r io.Reader) error {
	reader := bufio.NewReader(r)
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
			)
		}
	}

	return nil
}

func processDetailed(r io.Reader) error {
	reader := bufio.NewReader(r)
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
