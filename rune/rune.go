package rune2

import (
	"sort"
	"strconv"
	"strings"
)

const (
	ASCII = "ASCII"
	UTF   = "UTF"
)

type Rune rune

func (r Rune) CharCode() int {
	return int(r)
}
func (r Rune) CharCodeWithPadding() string {
	s := strconv.FormatInt(int64(r), 16)
	pad := 4 - len(s)
	if pad < 0 {
		pad = 0
	}
	return strings.Repeat("0", pad) + s
}
func (r Rune) IsNormalCharacter() bool {
	if r > 31 && r < 128 {
		return true
	}
	if r == 10 || r == 13 {
		return true
	}
	return false
}
func (r Rune) RuneType() string {
	if r < 256 {
		return ASCII
	}
	return UTF
}

type SpecialRunes map[int]Rune

func (sr SpecialRunes) SortedColumns() []int {
	cols := make([]int, len(sr), len(sr))
	j := 0
	for c, _ := range sr {
		cols[j] = c
		j++
	}
	sort.Ints(cols)
	return cols
}

func ProcessLine(line string) SpecialRunes {
	var rr Rune
	sr := make(SpecialRunes)

	for idx, r := range line {
		rr = Rune(r)
		if !rr.IsNormalCharacter() {
			sr[idx] = rr
		}
	}

	if len(sr) == 0 {
		return nil
	}
	return sr
}

type SpecialRunesLines map[int]SpecialRunes

func (srl SpecialRunesLines) SortedColumns() []int {
	lines := make([]int, len(srl), len(srl))
	j := 0
	for c, _ := range srl {
		lines[j] = c
		j++
	}
	sort.Ints(lines)
	return lines
}
