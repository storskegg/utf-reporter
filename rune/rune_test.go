package rune2

import (
    "reflect"
    "testing"
)

func TestProcessLine(t *testing.T) {
    type args struct {
        line string
    }
    tests := []struct {
        name string
        args args
        want SpecialRunes
    }{
        {
            name: "no special characters",
            args: args{line: "some boring text"},
            want: nil,
        },
        {
            name: "Señor",
            args: args{line: "Señor"},
            want: specialRunes{
                2: localRune('ñ'),
            },
        },
        {
            name: "4th semicolon is greek question mark",
            args: args{line: ";;;;;;;"},
            want: specialRunes{
                3: localRune(';'),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ProcessLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ProcessLine() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRune_CharCode(t *testing.T) {
    tests := []struct {
        name string
        r    Rune
        want int
    }{
        {
            name: "ñ, as in Señor",
            r:    localRune('ñ'),
            want: 241,
        },
        {
            name: "fancy forward double quotes",
            r:    localRune('“'),
            want: 8220,
        },
        {
            name: "greek question mark",
            r:    localRune(';'),
            want: 894,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.r.CharCode(); got != tt.want {
                t.Errorf("CharCode() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRune_CharCodeWithPadding(t *testing.T) {
    tests := []struct {
        name string
        r    Rune
        want string
    }{
        {
            name: "ñ, as in Señor",
            r:    localRune('ñ'),
            want: "00f1",
        },
        {
            name: "fancy forward double quotes",
            r:    localRune('“'),
            want: "201c",
        },
        {
            name: "greek question mark",
            r:    localRune(';'),
            want: "037e",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.r.CharCodeWithPadding(); got != tt.want {
                t.Errorf("CharCodeWithPadding() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRune_IsNormalCharacter(t *testing.T) {
    tests := []struct {
        name string
        r    Rune
        want bool
    }{
        {
            name: "Latin Capital S",
            r:    localRune('S'),
            want: true,
        },
        {
            name: "Left Double Quotation Mark",
            r:    localRune('“'),
            want: false,
        },
        {
            name: "ASCII Equal Sign",
            r:    localRune('='),
            want: true,
        },
        {
            name: "High Surrogates",
            r:    localRune('𝐜'),
            want: false,
        },
        {
            name: "ASCII standard space",
            r:    localRune(' '),
            want: true,
        },
        {
            name: "Zero Width Space (ZWSP)",
            r:    localRune('​'),
            want: false,
        },
        {
            name: "ASCII newline",
            r:    localRune('\n'),
            want: true,
        },
        {
            name: "Cyrillic Small Letter Ie",
            r:    localRune('е'),
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.r.IsNormalCharacter(); got != tt.want {
                t.Errorf("IsNormalCharacter() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSpecialRunes_SortedColumns(t *testing.T) {
    tests := []struct {
        name string
        s    SpecialRunes
        want []int
    }{
        {
            name: "rename me",
            s: &specialRunes{
                2:  localRune('ñ'),
                20: localRune('“'),
                13: localRune(';'),
            },
            want: []int{2, 13, 20},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.s.SortedColumns(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SortedColumns() = %v, want %v", got, tt.want)
            }
        })
    }
}
