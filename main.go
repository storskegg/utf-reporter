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
    "github.com/storskegg/utf-reporter/runic"
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

    flag.StringVar(&flagFile, "f", "", "Input text file")
    flag.Parse()
    if flag.NArg() > 0 {
        if !(flag.NArg() == 1 && flag.Arg(0) == "-") {
            printUsage()
            return
        }
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
    //captured := make(runic.SpecialRunesLines)

    lineNum := 0
    colNum := 0

    var sr runic.SpecialRunes
    var rr runic.Runic

    for {
        input, _, err := reader.ReadLine()
        if err != nil && err == io.EOF {
            break
        }

        if err != nil {
            panic(err)
            return
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

    if lineNum == 0 {
        fmt.Println("No non-standard characters found.")
        return
    }

    fmt.Println()
}
