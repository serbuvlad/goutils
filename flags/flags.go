package flags

import (
	"os"
	"unicode/utf8"
)

var boolFlags = make(map[rune]*bool)
var stringFlags = make(map[rune]**string)
var funcFlags = make(map[rune]func(rune, int) int)

func isFlag(r rune) bool {
	_, ok1 := boolFlags[r]
	_, ok2 := stringFlags[r]
	_, ok3 := funcFlags[r]

	if ok1 || ok2 || ok3 {
		return true
	}

	return false
}

var Args []string

func Bool(fp *bool, flag rune) {
	boolFlags[flag] = fp
}

func Func(f func(rune, int) int, flag rune) {
	funcFlags[flag] = f
}

func Parse(usage func()) {
	Args = make([]string, 0, len(os.Args))

	if len(os.Args) == 1 {
		return
	}

	var i int
	for i = 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if len(arg) == 0 {
			break
		}

		if arg == "-" {
			if !isFlag('-') {
				break
			}

			i += parseFlag('-', i, usage)
			continue
		}

		if arg == "--" || arg[0] != '-' {
			break
		}

		if utf8.RuneCountInString(arg) == 2 {
			flag, _ := utf8.DecodeRune([]byte(arg[1:]))
			i += parseFlag(flag, i, usage)
			continue
		}

		for _, flag := range arg[1:] {
			fp, ok := boolFlags[flag]
			if !ok {
				usage()
			}
			*fp = true
		}
	}

	Args = os.Args[i:]
}

func parseFlag(flag rune, i int, usage func()) int {
	fp, ok := boolFlags[flag]
	if ok {
		*fp = true
		return 0
	}

	sp, ok := stringFlags[flag]
	if ok {
		if len(os.Args) == i + 1 {
			usage()
		}
		a := os.Args[i + 1]
		*sp = &a
		return 1
	}

	f, ok := funcFlags[flag]
	if ok {
		return f(flag, i)
	}

	usage()

	const whoEvenCaresUsageDoesNotReturn = 69
	return whoEvenCaresUsageDoesNotReturn
}

func String(fp **string, flag rune) {
	stringFlags[flag] = fp
}
