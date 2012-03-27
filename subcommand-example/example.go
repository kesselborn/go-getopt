package main

import (
	"fmt"
	getopt "github.com/kesselborn/go-getopt"
	"os"
)

func main() {
	sco := getopt.SubCommandOptions{
		getopt.Options{
			"global description",
			getopt.Definitions{
				{"foo|f", "some arg", getopt.Optional, ""},
				{"command", "command to execute", getopt.IsSubCommand, ""},
			},
		},
		getopt.SubCommands{
			"getenv": {
				"getenv description",
				getopt.Definitions{
					{"bar|b", "some arg", getopt.Optional, ""},
					{"name", "app's name", getopt.IsArg | getopt.Required, ""},
					{"key", "environment variable's name", getopt.IsArg | getopt.Required, ""},
				},
			},
			"register": {
				"register description",
				getopt.Definitions{
					{"deploytype|t", "deploy type (one of mount, bazapta, lxc)", getopt.NoLongOpt | getopt.Optional | getopt.ExampleIsDefault, "lxc"},
					{"name|n", "app's name", getopt.IsArg | getopt.Required, ""},
				},
			},
		},
	}

	scope, options, arguments, passThrough, e := sco.ParseCommandLine()

	help, wantsHelp := options["help"]

	if e != nil || wantsHelp {
		exit_code := 0

		switch {
		case wantsHelp && help.String == "usage":
			fmt.Print(sco.Usage())
		case wantsHelp && help.String == "help":
			fmt.Print(sco.Help())
		default:
			fmt.Println("**** Error: ", e.Error(), "\n", sco.Help())
			exit_code = e.ErrorCode
		}
		os.Exit(exit_code)
	}

	fmt.Printf("scope:\n%s", scope)
	fmt.Printf("options:\n%#v", options)
	fmt.Printf("arguments: %#v\n", arguments)
	fmt.Printf("passThrough: %#v\n", passThrough)
}
