package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	rune2 "github.com/storskegg/utf-reporter/rune"
)

func printUsage() {
	fmt.Println("utf-reporter is intended to be used with piped input. File arguments are not yet supported.")
	fmt.Println("Example Usage: cat somefile.csv | utf-reporter")

	//flag.VisitAll(func(f *flag.Flag) {
	//	fmt.Printf("  %s\t%s\n", f.Name, f.Usage)
	//})
}

func main() {
	// Our Flags
	//var posix, all bool
	//
	//flag.BoolVar(&posix, "posix", false, "Use POSIX regex.  Defaults to PCRE")
	//flag.BoolVar(&all, "all", false, "Return all matches. Defaults to the first match")
	//
	//flag.Parse()
	//
	//if flag.NArg() < 1 {
	//	printUsage()
	//	return
	//}

	// Check stdin for piped input
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		printUsage()
		return
	}

	// Capture piped input, capturing runes
	reader := bufio.NewReader(os.Stdin)
	captured := make(rune2.SpecialRunesLines)

	lineNum := 0
	colNum := 0

	var sr rune2.SpecialRunes
	var rr rune2.Rune

	for {
		lineNum++
		input, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}

		sr := rune2.ProcessLine(string(input))
		if sr == nil {
			continue
		}
		captured[lineNum] = sr
	}

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	if len(captured) == 0 {
		fmt.Println("No non-standard characters found.")
	} else {
		for _, lineNum := range captured.SortedColumns() {
			fmt.Printf("===[ Line %d ]==========================================\n", lineNum)
			tbl := table.New("Column", "Rune", "Hex", "Type")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			sr = captured[lineNum]
			for _, colNum = range sr.SortedColumns() {
				rr = sr[colNum]
				tbl.AddRow(colNum, fmt.Sprintf("'%c'", rr), "0x"+rr.CharCodeWithPadding(), rr.RuneType())
				//fmt.Printf(" Found: 0x%s ('%c') at column %d of type %s\n", rr.CharCodeWithPadding(), rr, colNum, rr.RuneType())
			}
			tbl.Print()
			fmt.Println()
		}
	}
}
