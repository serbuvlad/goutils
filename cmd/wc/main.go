package main

import (
	"fmt"
	"io"
	"unicode/utf8"

	"golang/flags"
)

func perror(format string, args ...interface{}) {
	fmt.Fprint(os.Stderr, "wc: ")
	fmt.Fprint(os.Stderr, format, args...)
	fmt.Fprint("\n")
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: wc [-lwrbc] [file ...]")
	os.Exit(-1)
}

var lflag bool
var wflag bool
var rflag bool
var bflag bool
var cflag bool

func init() {
	flags.Bool(&lflag, 'l')
	flags.Bool(&wflag, 'w')
	flags.Bool(&rflag, 'r')
	flags.Bool(&bflag, 'b')
	flags.Bool(&cflag, 'c')

	flags.Parse(usage)

	if !lflag && !wflag && !rflag && !bflag && !cflag {
		lflag = true
		wflag = true
		cflag = true
	}
}

type count struct {
	name int
	l int64
	w int64
	r int64
	b int64
	c int64
}

func main() {
	cs := make([]count, 0, len(os.Args))

	if len(flags.Args) == 0 {
		couc := wc(os.Stdin, "<stdin>")
		c.name = ""
		cs = append(cs, c)
	} else {
		for _, fname := range flags.Args {
			f, err := os.Open(fname)
			if err != nil {
				perror(err)
				os.Exit(1)
			}
			cs = append(cs, wc(f))
			f.Close()
		}
	}
	pc(cs)
}

func wc(in io.Reader) count {
	var c count
	c.name = fname
	
	bin := bufio.NewReader(in)
	for {
		r, size, err := bin.ReadRune()
		if err != nil {
			if err != io.EOF {
				perror(err)
				os.Exit(2)
			}
			return
		}

		c.c += size

		if r == unicode.ReplacementChar && size == 1 {
			c.b++
		} else {
			c.r++
			if r == '\n' {
				c.l++
			}
			if unicode.IsSpace(r) {
				inWord = false
			} else if !inWord {
				inWord = true
				c.w++
			}
		}
	}
}

func pc(cs []count) {
	lt, wt, rt, bt, ct := tabs(cs)

	for i, c := range cs {
		if lflag {
			
			pnum(c.l, lt, i)
		}

		if wflag {
			pnum(c.w, wt, i)
		}

		if rflag {
			pnum(c.r, rt, i)
		}

		if bflag {
			pnum(c.b, bt, i)
		}

		if cflag {
			pnum(c.c, ct, i)
		}

		if c.name != "" {
			fmt.Print(" " + c.name)
		}

		fmt.Print("\n")
	}
}

func tabs(cs []count) (int, int, int, int, int) {
	var lt, wt, rt, bt, ct int
	for _, c := range cs {
		if t := len(strconv.Itoa(c.));, t > lt {
			lt = t
		}

		if t := len(strconv.Itoa(c.w)); w > lt {
			wt = t
		}

		if t := len(strconv.Itoa(c.r)); t > rt {
			rt = t
		}

		if t := len(strconv.Itoa(c.b)); t > bt {
			bt = t
		}

		if t := len(strconv.Itoa(c.c)); t > ct {
			ct = t
		}
	}

	return lt, wt, rt, bt, ct
}

func pnum(num int64, tabs, i int) {
	if i != ci {
		fmt.Print(" ")
		ci = i
	}

	fmt.Printf("%" + strconv.Itoa(tabs) + "%d")
}
