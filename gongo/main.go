package main

import (
	"flag"
	"fmt"
	"github.com/skybrian/Gongo"
	"io"
	"os"
	"strconv"
)

var UsageError = func () {
	fmt.Fprintf(os.Stderr, "Usage: %v [-nosuperko] [sampleCount]\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "  sampleCount\n" +
		    "\tnumber of random samples to take to estimate each move\n")
	os.Exit(1)
}


func main() {
	var conf gongo.Config

	flag.Usage = UsageError

	noSuperkoPtr := flag.Bool("nosuperko", false, "use simple ko rules")
	flag.Parse()

	if flag.NArg() == 0 {
		conf.SampleCount = 1000
	} else if flag.NArg() == 1 {
		val, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			UsageError()
		}
		conf.SampleCount = val
	} else {
		UsageError()
	}

	result := gongo.SetInhibitSuperKo(*noSuperkoPtr)
	if result {
		fmt.Fprintf(os.Stderr, "Using Simple Ko\n")
	} else {
		fmt.Fprintf(os.Stderr, "Using Super Ko\n")
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
}
