// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
	"strings"
	"os"
)

const InvalidOption = 1
const MissingValue = 2
const InvalidValue = 3
const MissingOption = 4
const OptionValueError = 5
const ConsistencyError = 6
const UsageOrHelp = 7

const OPTIONS_SEPARATOR = "--"

type GetOptError struct {
	errorCode int
	message   string
}

func mapifyEnviron(environment []string) (envArray map[string]string) {
	envArray = make(map[string]string)

	for _, cur := range environment {
		envVar := strings.Split(cur, "=")
		if len(envVar) > 1 {
			envArray[envVar[0]] = envVar[1]
		}
	}

	return
}

func (optionsDefinition Options) setOverwrites(options map[string]OptionValue, overwrites []string) (err *GetOptError) {
	overwritesMap := mapifyEnviron(overwrites)
	acceptedEnvVars := make(map[string]Option)

	for _, opt := range optionsDefinition {
		if value := opt.EnvVar(); value != "" {
			acceptedEnvVars[value] = opt
		}
	}

	for key, acceptedEnvVar := range acceptedEnvVars {
		if value := overwritesMap[key]; value != "" {
			options[acceptedEnvVar.LongOpt()], err = assignValue(acceptedEnvVar.DefaultValue, value)
			if err != nil {
				break
			}
		}
	}

	return
}

func checkOptionsDefinitionConsistency(optionsDefinition Options) (err *GetOptError) {

	for _, option := range optionsDefinition {
		switch {
		case option.Flags&Optional > 0 && option.Flags&Required > 0:
			err = &GetOptError{ConsistencyError, "an option can not be Required and Optional"}
		case option.Flags&Flag > 0 && option.Flags&ExampleIsDefault > 0:
			err = &GetOptError{ConsistencyError, "an option can not be a Flag and have ExampleIsDefault"}
		case option.Flags&Required > 0 && option.Flags&ExampleIsDefault > 0:
			err = &GetOptError{ConsistencyError, "an option can not be Required and have ExampleIsDefault"}
		case option.Flags&Required > 0 && option.Flags&IsArg > 0:
			err = &GetOptError{ConsistencyError, "an option can not be Required and be an argument (IsArg)"}
		case option.Flags&NoLongOpt > 0 && !option.HasShortOpt() && option.Flags&IsArg == 0:
			err = &GetOptError{ConsistencyError, "an option must have either NoLongOpt or a ShortOption"}
		case option.Flags&Flag > 0 && option.Flags&IsArg > 0:
			err = &GetOptError{ConsistencyError, "an option can not be a Flag and be an argument (IsArg)"}
		}
	}

	return
}

func (optionsDefinition Options) usageHelpOptionNames() (shortOpt string, longOpt string) {
	shortOpt = "h"
	longOpt = "help"

	for _, option := range optionsDefinition {
		if option.Flags&Usage > 0 {
			shortOpt = option.ShortOpt()
		}
		if option.Flags&Help > 0 {
			longOpt = option.LongOpt()
		}
	}

	return
}

// copied from the os package ... why isn't this exposed :(
func basename(name string) string {
	i := len(name) - 1
	// Remove trailing slashes
	for ; i > 0 && name[i] == '/'; i-- {
		name = name[:i]
	}
	// Remove leading directory name
	for i--; i >= 0; i-- {
		if name[i] == '/' {
			name = name[i+1:]
			break
		}
	}

	return name
}

// todo: method signature sucks
func (optionsDefinition Options) checkForHelpOrUsage(args []string, usageString string, helpString string, description string) (err *GetOptError) {
	for _, arg := range args {
		switch {
		case arg == usageString:
			err = &GetOptError{UsageOrHelp, optionsDefinition.Usage(basename(os.Args[0]))}
		case arg == helpString:
			err = &GetOptError{UsageOrHelp, optionsDefinition.Help(basename(os.Args[0]), description)}
		}
	}

	return
}

func (optionsDefinition Options) parse(args []string,
defaults []string,
description string,
flags int) (options map[string]OptionValue,
arguments []string,
passThrough []string,
err *GetOptError) {

	if err = checkOptionsDefinitionConsistency(optionsDefinition); err == nil {
		options = make(map[string]OptionValue)
		arguments = make([]string, 0)

		for _, option := range optionsDefinition {
			switch {
			case option.Flags&Flag != 0: // all flags are false by default
				options[option.Key()], err = assignValue(false, "false")
			case option.Flags&ExampleIsDefault != 0: // set default
				options[option.Key()], err = assign(option.DefaultValue)
			}
		}

		// set overwrites
		usageString, helpString := optionsDefinition.usageHelpOptionNames()
		usageString = "-" + usageString
		helpString = "--" + helpString
		err = optionsDefinition.checkForHelpOrUsage(args, usageString, helpString, description)

		if err == nil {
			err = optionsDefinition.setOverwrites(options, defaults)

			for i := 0; i < len(args) && err == nil; i++ {

				var opt, val string
				var found bool

				token := args[i]

				if argumentsEnd(token) {
					passThrough = args[i:]
					break
				}

				if isValue(token) {
					arguments = append(arguments, token)
					continue
				}

				opt, val, found = parseShortOpt(token)

				if found {
					buffer := token

					for found && optionsDefinition.IsFlag(opt) && len(buffer) > 2 {
						// concatenated options ... continue parsing
						currentOption, _ := optionsDefinition.FindOption(opt)
						key := currentOption.Key()

						options[key], err = assignValue(currentOption.DefaultValue, "true")

						// make it look as if we have a normal option with a '-' prefix
						buffer = "-" + buffer[2:]
						opt, val, found = parseShortOpt(buffer)
					}

				} else {
					opt, val, found = parseLongOpt(token)
				}

				currentOption, found := optionsDefinition.FindOption(opt)
				key := currentOption.Key()

				if !found {
					err = &GetOptError{InvalidOption, "invalid option '" + token + "'"}
					break
				}

				if optionsDefinition.IsFlag(opt) {
					options[key], err = assignValue(true, "true")
				} else {
					if val == "" {
						if len(args) > i+1 && isValue(args[i+1]) {
							i = i + 1
							val = args[i]
						} else {
							err = &GetOptError{MissingValue, "Option '" + token + "' needs a value"}
							break
						}
					}

					if !isValue(val) {
						err = &GetOptError{InvalidValue, "Option '" + token + "' got invalid value: '" + val + "'"}
						break
					}

					options[key], err = assignValue(currentOption.DefaultValue, val)
				}

			}
		}

		if err == nil {
			for _, requiredOption := range optionsDefinition.RequiredOptions() {
				if options[requiredOption].set == false {
					err = &GetOptError{MissingOption, "Option '" + requiredOption + "' is missing"}
					break
				}
			}
		}
	}

	return
}
