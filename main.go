package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/chzyer/readline/runes"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	rune2 "github.com/storskegg/utf-reporter/rune"
)

func printUsage() {
	fmt.Println("Examples:")
	fmt.Println("  cat somefile.csv | utf-reporter")
	fmt.Println("  utf-reporter -f <path to text file>")
	fmt.Println()
	fmt.Println("Usage of utf-reporter:")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("  %s\t%s\n", f.Name, f.Usage)
	})
}

func main() {
	var err error
	var info os.FileInfo

	// Our Flags
	var flagFile string
	var flagNoColor bool

	flag.StringVar(&flagFile, "f", "", "Input text file")
	flag.BoolVar(&flagNoColor, "no-color", false, "Disable color output")
	flag.Parse()
	if flag.NArg() > 0 {
		if !(flag.NArg() == 1 && flag.Arg(0) == "-") {
			printUsage()
			return
		}
	}

	if flagNoColor {
		color.NoColor = true
	}

	var input *os.File

	// this could be cleaned up a bit
	if flagFile == "" {
		input = os.Stdin
		info, err = input.Stat()
		if err != nil {
			panic(err)
		}
	} else {
		flagFile = path.Join(".", flagFile)
		info, err = os.Stat(flagFile)
		if err != nil {
			if err == os.ErrNotExist {
				fmt.Printf("The path '%s' does not exist.\n", flagFile)
				return
			}
			panic(err)
		}
		input, err = os.Open(flagFile)
		if err != nil {
			panic(err)
		}
	}

	// Check stdin for piped input
	if flagFile == "" {
		if info.Mode()&os.ModeNamedPipe == 0 {
			printUsage()
			return
		}
	} else {
		if !info.Mode().IsRegular() {
			printUsage()
			return
		}
	}

	// Capture piped input, capturing runes
	reader := bufio.NewReader(input)
	//captured := make(rune2.SpecialRunesLines)

	lineNum := 0
	colNum := 0

	var sr rune2.SpecialRunes
	var rr rune2.Rune

	// Set up our console formatters
	titleFmt := color.New(color.FgHiWhite, color.Bold).PrintfFunc()
	headerFmt := color.New(color.FgHiBlue, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	// Instantiate a reusable table to be reassigned each lines iteration
	var tbl table.Table

	for {
		// Create a table for each line
		tbl = table.New("Column", "Rune", "Hex", "Type", "Width", "Reference")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(3)
		tbl.WithWidthFunc(func(s string) int {
			return runes.WidthAll([]rune(s))
		})

		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		sr = rune2.ProcessLine(string(input))
		if sr == nil {
			continue
		}
		for _, colNum = range sr.SortedColumns() {
			rr, _ = sr.Get(colNum)
			tbl.AddRow(
				colNum,
				fmt.Sprintf("'%c'", rr),
				"0x"+rr.CharCodeWithPadding(),
				rr.RuneType(),
				runes.Width(rr.Rune()),
				fmt.Sprintf("https://www.compart.com/en/unicode/U+%s", strings.ToUpper(rr.CharCodeWithPadding())),
			)
		}

		// Print our title and table
		titleFmt("Line %d Table\n", lineNum+1)
		tbl.Print()
		fmt.Println()
		//captured[lineNum] = sr
		lineNum++
	}

	if lineNum == 0 {
		fmt.Println("No non-standard characters found.")
	}
}
