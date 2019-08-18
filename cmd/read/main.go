package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"goutils/flags"
)

const bufSize = 4096

func perror(err ...interface{}) {
	fmt.Fprint(os.Stderr, "read: ")
	fmt.Fprintln(os.Stderr, err...)
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: read [-m] [-n nline] [file ...]")
}

var nline int = 1
var mflag bool

func init() {
	var nflag *string
	flags.String(&nflag, 'n')
	flags.Bool(&mflag, 'm')

	flags.Parse(usage)

	if nflag != nil {
		var err error
		nline, err = strconv.Atoi(*nflag)

		if err != nil {
			perror("could not parse -n: ", err)
			usage()
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		read(os.Stdin, "<stdin>")
		return
	}

	for _, fname := range flags.Args[:] {
		f, err := os.Open(fname)
		if err != nil {
			perror("could not open file ", fname, ": ", err)
			os.Exit(1)
		}
		done := read(f, fname)
		f.Close()
		if done {
			return
		}
	}
}

func read(in io.Reader, fname string) (done bool) {
	b := make([]byte, 0, bufSize)

	var errRead error
	for errRead != io.EOF {
		var n int
		n, errRead = readLine(in, b)
		if errRead != nil && errRead != io.EOF {
			perror("error reading file ", fname, ": ", errRead)
			os.Exit(2)
		}

		_, errWrite := os.Stdout.Write(b[:n])
		if errWrite != nil {
			perror("error writting: ", errWrite)
			os.Exit(3)
		}

		nline--
		if nline == 0 && !mflag {
			return true
		}
	}
	return false
}

var a = make([]byte, 1)
func readLine(in io.Reader, b []byte) (n int, err error) {
	b = b[:0]
	n = 0
	for len(b) == 0  || b[n - 1] != '\n' {
		_, err = in.Read(a)
		b = append(b, a[0])
		n++

		if err != nil {
			return
		}
	}
	return
}

