// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
//"os"
//"path/filepath"
//"sort"
//"strings"
//"fmt"
)

type Scopes map[string]SubCommandOptions

type SubSubCommandOptions struct {
	Global Options
	Scopes Scopes
}

func (ssco SubSubCommandOptions) flattenToSubcommandOptions(scope string) (sco SubCommandOptions, err *GetOptError) {
	globalCommand := ssco.Global
	var present bool

	if sco, present = ssco.Scopes[scope]; present == true {
		//print("\n======= " + scope + "========\n"); print(strings.Replace(fmt.Sprintf("%#v", scope.SubCommands),  "getopt", "\n", -1)); print("\n================")
		sco.Global.Definitions = append(globalCommand.Definitions, sco.Global.Definitions...)
	} else {
		err = &GetOptError{UnknownSubcommand, "Unknown scope: " + scope}
	}

	return
}

func (ssco SubSubCommandOptions) flattenToOptions(scope string, subCommand string) (options Options, err *GetOptError) {
	if sco, err := ssco.flattenToSubcommandOptions(scope); err == nil {
		options, err = sco.flattenToOptions(subCommand)
	}

	return
}

//func (sco SubCommandOptions) findSubcommand() (subCommand string, err *GetOptError) {
//	options := sco["*"]
//	subCommand = "*"
//
//	_, arguments, _, _ := options.ParseCommandLine()
//
//	if len(arguments) < 1 {
//		err = &GetOptError{NoSubcommand, "Couldn't find sub command"}
//	} else {
//		subCommand = arguments[0]
//	}
//
//	return
//}
//
//func (sco SubCommandOptions) ParseCommandLine() (subCommand string, options map[string]OptionValue, arguments []string, passThrough []string, err *GetOptError) {
//
//	if subCommand, err = sco.findSubcommand(); err == nil {
//		var flattenedOptions Options
//		if flattenedOptions, err = sco.flattenToOptions(subCommand); err == nil {
//			options, arguments, passThrough, err = flattenedOptions.ParseCommandLine()
//			arguments = arguments[1:]
//		}
//	}
//
//	return
//}
//
//func (sco SubCommandOptions) Usage(scope string) (output string) {
//	return sco.UsageCustomArg0(scope, filepath.Base(os.Args[0]))
//}
//
//func (sco SubCommandOptions) UsageCustomArg0(scope string, arg0 string) (output string) {
//	subCommand, err := sco.findSubcommand()
//
//	if err != nil {
//		subCommand = "*"
//	} else {
//		arg0 = arg0 + " " + scope
//	}
//
//	if flattenedOptions, present := sco[subCommand]; present {
//		output = flattenedOptions.UsageCustomArg0(arg0)
//	}
//
//	return
//}
//
//func (sco SubCommandOptions) Help(description string, scope string) (output string) {
//	return sco.HelpCustomArg0(description, scope, filepath.Base(os.Args[0]))
//}
//
//func (sco SubCommandOptions) HelpCustomArg0(description string, scope string, arg0 string) (output string) {
//	subCommand, err := sco.findSubcommand()
//
//	if err != nil {
//		subCommand = "*"
//	} else {
//		arg0 = arg0 + " " + scope
//	}
//
//	if flattenedOptions, present := sco[subCommand]; present {
//		output = flattenedOptions.HelpCustomArg0(description, arg0)
//	}
//
//	if subCommand == "*" {
//		output = output + "Available commands:\n"
//
//		keys := make([]string, len(sco))
//		i := 0
//
//		for k := range sco {
//			keys[i] = k
//			i = i + 1
//		}
//		sort.Strings(keys)
//
//		for _, key := range keys {
//			if key != "*" {
//				output = output + "    " + key + "\n"
//			}
//		}
//		output = output + "\n"
//	}
//
//	return
//}
