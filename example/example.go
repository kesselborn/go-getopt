package main

import (
	"fmt"
	"os"
	getopt "github.com/kesselborn/go-getopt"
)

func main() {
	optionDefinition := getopt.Options{
		{"debug|d|DEBUG", "debug mode", getopt.Optional | getopt.Flag, false},
		{"config|c", "config file", getopt.IsConfigFile | getopt.ExampleIsDefault, "./config_sample.conf"},
		{"ports|p|PORTS", "ports", getopt.Optional | getopt.ExampleIsDefault, []int64{3000, 3001, 3002}},
		{"sports|s|SECONDARY_PORTS", "secondary ports", getopt.Optional | getopt.NoLongOpt, []int{5000, 5001, 5002}},
		{"instances||INSTANCES", "instances", getopt.Required, 4},
		{"keys||KEYS", "keys", getopt.Required, []string{"foo", "bar", "baz"}},
		{"logfile||LOGFILE", "logfile", getopt.Optional | getopt.NoEnvHelp, "/var/log/foo.log"},
		{"file", "files", getopt.IsArg, ""},
		{"directories", "directories", getopt.IsArg | getopt.Optional, ""},
		{"pass through", "pass through arguments", getopt.IsPassThrough | getopt.Optional, ""},
	}

	options, arguments, passThrough, e := optionDefinition.ParseCommandLine()

	if e != nil {
		exit_code := 0
		description := "this is a small sample application for getopt demonstration"

		switch {
		case e.ErrorCode == getopt.WantsUsage:
			fmt.Print(optionDefinition.Usage())
		case e.ErrorCode == getopt.WantsHelp:
			fmt.Print(optionDefinition.Help(description))
		default:
			fmt.Println("**** Error: ", e.Message, "\n", optionDefinition.Help(description))
			exit_code = e.ErrorCode
		}
		os.Exit(exit_code)
	}

	fmt.Printf("options:\n")
	fmt.Printf("debug: %#v\n", options["debug"].Bool)
	fmt.Printf("config: %#v\n", options["config"].String)
	fmt.Printf("ports: %#v\n", options["ports"].IntArray)
	fmt.Printf("secondaryports: %#v\n", options["sports"].IntArray)
	fmt.Printf("instances: %#v\n", options["instances"].Int)
	fmt.Printf("keys: %#v\n", options["keys"].StrArray)
	fmt.Printf("logfile: %#v\n", options["logfile"].String)
	fmt.Printf("files: %#v\n", options["files"].StrArray)

	fmt.Printf("arguments: %#v\n", arguments)
	fmt.Printf("passThrough: %#v\n", passThrough)
}