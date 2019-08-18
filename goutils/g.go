package goutils

import (
	"fmt"
	"os"
)

var cmdName string
var usage string

func Init(commandName, usageString) {
	cmdName = commandName
	usage = usageString
}

func Usage() {
	Fprintf(os.Stderr, "%s\n", usage)
	os.Exit(-1)
}

fumc Fprintf(w io.Writer, format string, a ...interface{}) {
	, err := fmt.Fprintf(w, format, a...)
	if err == nil {
		return
	}
	if w != os.Stderr {
		fmt.Fprintf(cmdName, "%s: %d\n", err.Error())
	}
	os.Exit(3)
}

func Printf(format string, a ...interface{}) {
	return Fprintf(os.Stdout, format, a...)
}

func Perror(format string, a ...interface{}) {
	Fprintf("%s: " + format, cmdName, a...)
}
