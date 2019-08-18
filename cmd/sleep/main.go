package main

import (
	"fmt"
	"os"
	"time"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: sleep time")
	os.Exit(-1)
}

var d time.Duration

func init() {
	if len(os.Args) != 2 {
		usage()
	}

	var err error
	d, err = time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "sleep: could not parse time: ", err)
		usage()
	}
}

func main() {
	time.Sleep(d)
}
