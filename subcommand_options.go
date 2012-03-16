// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
	"os"
	"path/filepath"
	"sort"
)

type SubCommands map[string]Options

type SubCommandOptions struct {
	Global      Options
	SubCommands SubCommands
}

func (sco SubCommandOptions) flattenToOptions(subCommand string) (options Options, err *GetOptError) {
	genericOptions := sco.Global

	if subCommandOptions, present := sco.SubCommands[subCommand]; present == true {

		if subCommand != "*" {
			for _, option := range genericOptions.definitions {
				options.definitions = append(options.definitions, option)
			}
		}

		for _, option := range subCommandOptions.definitions {
			options.definitions = append(options.definitions, option)
			options.description = subCommandOptions.description
		}
	} else {
		err = &GetOptError{UnknownSubcommand, "Unknown command: " + subCommand}
	}

	return
}

func (sco SubCommandOptions) findSubcommand() (subCommand string, err *GetOptError) {
	options := sco.Global
	subCommand = "*"

	_, arguments, _, _ := options.ParseCommandLine()

	if len(arguments) < 1 {
		err = &GetOptError{NoSubcommand, "Couldn't find sub command"}
	} else {
		subCommand = arguments[0]
	}

	return
}

func (sco SubCommandOptions) ParseCommandLine() (subCommand string, options map[string]OptionValue, arguments []string, passThrough []string, err *GetOptError) {

	if subCommand, err = sco.findSubcommand(); err == nil {
		var flattenedOptions Options
		if flattenedOptions, err = sco.flattenToOptions(subCommand); err == nil {
			options, arguments, passThrough, err = flattenedOptions.ParseCommandLine()
			arguments = arguments[1:]
		}
	}

	return
}

func (sco SubCommandOptions) Usage(scope string) (output string) {
	return sco.UsageCustomArg0(scope, filepath.Base(os.Args[0]))
}

func (sco SubCommandOptions) UsageCustomArg0(scope string, arg0 string) (output string) {
	subCommand, err := sco.findSubcommand()
	flattenedOptions, foundSubCommand := sco.SubCommands[subCommand]

	if err != nil || !foundSubCommand {
		output = sco.Global.UsageCustomArg0(arg0)
	} else {
		output = flattenedOptions.UsageCustomArg0(arg0 + " " + scope)
	}

	return
}

func (sco SubCommandOptions) Help(description string, scope string) (output string) {
	return sco.HelpCustomArg0(description, scope, filepath.Base(os.Args[0]))
}

func (sco SubCommandOptions) HelpCustomArg0(description string, scope string, arg0 string) (output string) {
	subCommand, err := sco.findSubcommand()
	flattenedOptions, foundSubCommand := sco.SubCommands[subCommand]

	if err != nil || !foundSubCommand {
		subCommand = "*"
		flattenedOptions = sco.Global
	} else {
		arg0 = arg0 + " " + scope
	}

	output = flattenedOptions.HelpCustomArg0(arg0)

	if subCommand == "*" {
		output = output + "Available commands:\n"

		keys := make([]string, len(sco.SubCommands))
		i := 0

		for k := range sco.SubCommands {
			keys[i] = k
			i = i + 1
		}
		sort.Strings(keys)

		for _, key := range keys {
			output = output + "    " + key + "\n"
		}
		output = output + "\n"
	}

	return
}
