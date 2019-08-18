package main

import (
	"fmt"
	"os"
	"sort"

	"goutils/flags"
)

func perror(err ...interface{}) {
	fmt.Fprint(os.Stderr, "ls: ")
	fmt.Fprintln(os.Stderr, err...)
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: ls [] name ...")
	os.Exit(-1)
}

var dflag bool
var lflag bool
var nflag bool
var tflag bool
var rflag bool
var cfflag bool

func init() {
	flags.Bool(&dflag, 'd')
	flags.Bool(&lflag, 'l')
	flags.Bool(&nflag, 'n')
	flags.Bool(&tflag, 't')
	flags.Bool(&rflag, 'r')
	flags.Bool(&cfflag, 'F')

	flags.Parse(usage)
}

type entry struct {
	info os.FileInfo
	name string
}

type entries []entry

func (e entries) Len() int {
	return len(e)
}

func (e entries) Less(i, j int) bool {
	if tflag {
		return e[i].info.ModTime().Before(e[j].info.ModTime())
	}
	return e[i].info.Name() < e[j].info.Name()
}

func (e entries) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func main() {
	if len(flags.Args) == 0 {
		ls(".", "")
	}
}

func ls(fname, prefix string) {
	f, err := os.Open(fname)
	if err != nil {
		perror("could not open file ", fname, ": ", err)
		os.Exit(1)
	}

	var entries []entry
	if dflag {
		entries = []entry{ entry{stat(f, fname), fname}}
	} else {
		entriesList, err := f.Readdir(0)
		if err != nil {
			perror("could not read directory ", fname, "; ", err)
			os.Exit(3)
		}

		entries = make([]entry, len(entriesList))
		for i, e := range entriesList {
			entries[i] = entry{ e, prefix + e.Name() }
		}
	}

	sorte(entries)
	le(entries)
}


func stat(f *os.File, fname string) os.FileInfo {
	fi, err := f.Stat()
	if err != nil {
		perror("could not stat file ", fname, ": ", err)
		os.Exit(2)
	}
	return fi
}

func sorte(entries []entry) {
	if nflag {
		return
	}

	var i sort.Interface = entries(entries)
	if rflag {
		i = sort.Reverse(i)
	}
	sort.Sort(i)
}

func le(entries []entry) {
	for _, e := range entries {
		l(e)
	}
}

func l(e entry) {
	if lflag {
		pmode(e.info.Mode())
		fmt.Print(" ")
		fmt.Print(e.info.Size())
		fmt.Print(" ")
		fmt.Print(e.info.ModTime())
		fmt.Print(" ")
	}
	pname(e)
	fmt.Println("")
}

func pmode(m os.FileMode) {
	fmt.Print(m.String())
}

func pname(e entry) {
	fmt.Print(e.name)
	if cfflag {
		if e.info.IsDir() {
			fmt.Print("/")
		}
	}
}
