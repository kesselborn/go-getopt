package main

import (
	"fmt"
	"os"
	getopt "github.com/kesselborn/go-getopt"
)

func main() {
	sco := getopt.SubCommandOptions{
		"*": {
			{"foo|f", "some arg", getopt.Optional, ""},
			{"command", "command to execute", getopt.IsSubcommand, ""}},
		"getenv": {
			{"bar|b", "some arg", getopt.Optional, ""},
			{"name", "app's name", getopt.IsArg | getopt.Required, ""},
			{"key", "environment variable's name", getopt.IsArg | getopt.Required, ""}},
		"register": {
			{"deploytype|t", "deploy type (one of mount, bazapta, lxc)", getopt.NoLongOpt | getopt.Optional | getopt.ExampleIsDefault, "lxc"},
			{"name|n", "app's name", getopt.IsArg | getopt.Required, ""}},
	}

	scope, options, arguments, passThrough, e := sco.ParseCommandLine()

	if e != nil {
		exit_code := 0
		description := "this is a small sample application for getopt demonstration"

		switch {
		case e.ErrorCode == getopt.WantsUsage:
			fmt.Print(sco.Usage(scope))
		case e.ErrorCode == getopt.WantsHelp:
			fmt.Print(sco.Help(description, scope))
		default:
			fmt.Println("**** Error: ", e.Message, "\n", sco.Help(description, scope))
			exit_code = e.ErrorCode
		}
		os.Exit(exit_code)
	}

  fmt.Printf("scope:\n%s", scope)
	fmt.Printf("options:\n%#v", options)
	fmt.Printf("arguments: %#v\n", arguments)
	fmt.Printf("passThrough: %#v\n", passThrough)
}
