package main

import (
	"fmt"
	getopt "github.com/kesselborn/go-getopt"
	"os"
)

func main() {
	ssco := getopt.SubSubCommandOptions{
		getopt.Options{
			"global description",
			getopt.Definitions{
				{"config|c", "config file", getopt.IsConfigFile | getopt.ExampleIsDefault, "/etc/visor.conf"},
				{"server|s", "doozer server", getopt.Required, ""},
				{"scope", "scope", getopt.IsSubCommand, ""},
			},
		},
		getopt.Scopes{
			"app": {
				getopt.Options{
					"app description",
					getopt.Definitions{
						{"foo|f", "a param", getopt.Optional, ""},
						{"command", "command to execute", getopt.IsSubCommand, ""},
					},
				},
				getopt.SubCommands{
					"getenv": {
						"app getenv description",
						getopt.Definitions{
							{"key", "environment variable's name", getopt.IsArg | getopt.Required, ""},
						},
					},
					"setenv": {
						"app setenv description",
						getopt.Definitions{
							{"persist|p", "persist this", getopt.Optional | getopt.Flag | getopt.NoLongOpt, ""},
							{"alias|a", "alias name", getopt.Optional, ""},
							{"name", "app name", getopt.IsArg | getopt.Required, ""},
							{"key", "env key", getopt.IsArg | getopt.Required, ""},
							{"value", "env value", getopt.IsArg | getopt.Optional, ""},
						},
					},
				},
			},
			"revision": {
				getopt.Options{
					"app revision description",
					getopt.Definitions{
						{"rev|r", "revision", getopt.IsArg | getopt.Required, ""},
						{"command", "command to execute", getopt.IsSubCommand, ""},
					},
				},
				getopt.SubCommands{
					"list": {
						"list revisions",
						getopt.Definitions{
							{"all|a", "long list output", getopt.Flag, ""},
						},
					},
				},
			},
		},
	}

	scope, subCommand, options, arguments, passThrough, e := ssco.ParseCommandLine()

	help, wantsHelp := options["help"]

	if e != nil || wantsHelp {
		exit_code := 0

		switch {
		case wantsHelp && help.String == "usage":
			fmt.Print(ssco.Usage())
		case wantsHelp && help.String == "help":
			fmt.Print(ssco.Help())
		default:
			fmt.Println("**** Error: ", e.Error(), "\n", ssco.Help())
			exit_code = e.ErrorCode
		}
		os.Exit(exit_code)
	}

	fmt.Printf("scope:\n%s", scope)
	fmt.Printf("subCommand:\n%s", subCommand)
	fmt.Printf("options:\n%#v", options)
	fmt.Printf("arguments: %#v\n", arguments)
	fmt.Printf("passThrough: %#v\n", passThrough)
}
