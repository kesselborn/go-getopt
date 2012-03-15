// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

type SubCommandOptions map[string]Options

func (sco SubCommandOptions) flattenToOptions(subCommand string) (options Options, err *GetOptError) {
	genericOptions := sco["*"]

	if subCommandOptions, present := sco[subCommand]; present == true {

		for _, option := range genericOptions {
			options = append(options, option)
		}
		for _, option := range subCommandOptions {
			options = append(options, option)
		}
	} else {
		err = &GetOptError{UnknownSubcommand, "Unknown command: " + subCommand}
	}

	return
}

func (sco SubCommandOptions) findSubcommand() (subCommand string, err *GetOptError) {
	options := sco["*"]

	_, arguments, _, _ := options.ParseCommandLine()

	if len(arguments) < 1 {
		err = &GetOptError{NoSubcommand, "Couldn't find sub command"}
	} else {
		subCommand = arguments[0]
	}

	return
}

func (sco SubCommandOptions) ParseCommandLine() (options map[string]OptionValue, arguments []string, passThrough []string, err *GetOptError) {
	var subCommand string

	if subCommand, err = sco.findSubcommand(); err == nil {
		var flattenedOptions Options
		if flattenedOptions, err = sco.flattenToOptions(subCommand); err == nil {
			options, arguments, passThrough, err = flattenedOptions.ParseCommandLine()
		}
	}

	return
}
