package rune2

import (
	"sort"
	"strconv"
	"strings"
)

type RuneType int

func (rt RuneType) String() string {
	return runeTypeLabels[rt]
}

const (
	ASCII RuneType = iota
	UTF
)

const (
	RuneTypeLabelASCII = "ASCII"
	RuneTypeLabelUTF   = "UTF"
)

var runeTypeLabels = map[RuneType]string{
	ASCII: RuneTypeLabelASCII,
	UTF:   RuneTypeLabelUTF,
}

type Rune interface {
	Rune() rune
	CharCode() int
	CharCodeWithPadding() string
	IsNormalCharacter() bool
	RuneType() RuneType
}

func NewRune(r rune) Rune {
	return localRune(r)
}

type localRune rune

func (r localRune) Rune() rune {
	return rune(r)
}

func (r localRune) CharCode() int {
	return int(r)
}

func (r localRune) CharCodeWithPadding() string {
	s := strconv.FormatInt(int64(r.Rune()), 16)
	pad := 4 - len(s)
	if pad < 0 {
		pad = 0
	}
	return strings.Repeat("0", pad) + s
}

func (r localRune) IsNormalCharacter() bool {
	if r > 31 && r < 128 {
		return true
	}
	if r == 10 || r == 13 {
		return true
	}
	return false
}

func (r localRune) RuneType() RuneType {
	if r < 256 {
		return ASCII
	}
	return UTF
}

type SpecialRunes interface {
	SortedColumns() []int
	Get(idx int) (Rune, bool)
	Set(idx int, r Rune)
	Len() int
}

type specialRunes map[int]Rune

func NewSpecialRunes() SpecialRunes {
	return make(specialRunes)
}

func (sr specialRunes) SortedColumns() []int {
	cols := make([]int, len(sr), len(sr))
	j := 0
	for c, _ := range sr {
		cols[j] = c
		j++
	}
	sort.Ints(cols)
	return cols
}

func (sr specialRunes) Get(idx int) (Rune, bool) {
	r, ok := sr[idx]
	return r, ok
}

func (sr specialRunes) Set(idx int, r Rune) {
	sr[idx] = r
}

func (sr specialRunes) Len() int {
	return len(sr)
}

func ProcessLine(line string) SpecialRunes {
	var rr Rune
	sr := NewSpecialRunes()

	for idx, r := range line {
		rr = NewRune(r)
		if !rr.IsNormalCharacter() {
			sr.Set(idx, rr)
		}
	}

	if sr.Len() == 0 {
		return nil
	}
	return sr
}

type SpecialRunesLines interface {
	SortedColumns() []int
}

func NewSpecialRunesLines() SpecialRunesLines {
	return make(specialRunesLines)
}

type specialRunesLines map[int]SpecialRunes

func (srl specialRunesLines) SortedColumns() []int {
	lines := make([]int, len(srl), len(srl))
	j := 0
	for c, _ := range srl {
		lines[j] = c
		j++
	}
	sort.Ints(lines)
	return lines
}
