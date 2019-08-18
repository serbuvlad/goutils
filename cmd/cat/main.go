package main

import (
	"io"
	"os"
	"fmt"
)

const bufSize = 4096

func perror(err ...interface{}) {
	fmt.Fprint(os.Stderr, "cat: ")
	fmt.Fprintln(os.Stderr, err...)
}

func main() {
	if len(os.Args) == 1 {
		cat(os.Stdin, "<stdin>")
		return
	}

	for _, fname := range os.Args[1:] {
		f, err := os.Open(fname)
		if err != nil {
			perror("could not open file ", fname, ": ", err)
			os.Exit(1)
		}

		cat(f, fname)
		f.Close()
	}
}

func cat(in io.Reader, fname string) {
	b := make([]byte, bufSize)

	var errRead error
	for errRead != io.EOF {
		var n int
		n, errRead = in.Read(b)
		if errRead != nil && errRead != io.EOF {
			perror("error reading file ", fname, ": ", errRead)
			os.Exit(2)
		}

		_, errWrite := os.Stdout.Write(b[:n])
		if errWrite != nil {
			perror("error writting: ", errWrite)
			os.Exit(3)
		}
	}
}

