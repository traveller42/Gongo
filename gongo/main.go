package main

import (
	"fmt"
	"github.com/skybrian/Gongo"
	"io"
	"os"
	"strconv"
)

func UsageError() {
	fmt.Fprintf(os.Stderr, "Usage: %v [--nosuperko] [sampleCount]\n\n", os.Args[0])
	os.Exit(1)
}

func main() {
	var conf gongo.Config
	result := gongo.SetInhibitSuperKo(false)
	if len(os.Args) == 1 {
		conf.SampleCount = 1000
	} else if len(os.Args) == 2 {
		if os.Args[1] == "--nosuperko" {
			result = gongo.SetInhibitSuperKo(true)
		} else {
			val, err := strconv.Atoi(os.Args[1])
			if err != nil {
				UsageError()
			}
			conf.SampleCount = val
		}
	} else if len(os.Args) == 3 {
		if os.Args[1] == "--nosuperko" {
			result = gongo.SetInhibitSuperKo(true)
			val, err := strconv.Atoi(os.Args[2])
			if err != nil {
				UsageError()
			}
			conf.SampleCount = val
		} else {
			UsageError()
		}
	} else {
		UsageError()
	}
	bot, err := gongo.NewConfiguredMultiRobot(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error: %v", err)
	}
	err = gongo.Run(bot, os.Stdin, os.Stdout)
	if err == io.EOF {
		fmt.Fprintln(os.Stderr, "got EOF")
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error: %v", err)
	}
	if result {
		fmt.Fprintf(os.Stderr, "Using Simple Ko\n")
	} else {
		fmt.Fprintf(os.Stderr, "Using Super Ko\n")
	}
}
