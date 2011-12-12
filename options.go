// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
	"fmt"
	"os"
)

type Options []Option

func (options Options) FindOption(optionString string) (option Option, found bool) {
	for _, cur := range options {
		if cur.ShortOpt() == optionString || cur.LongOpt() == optionString {
			option = cur
			found = true
			break
		}
	}

	return option, found
}

func (options Options) IsOptional(optionName string) (isRequired bool) {
	if option, found := options.FindOption(optionName); found && option.Flags&Optional != 0 {
		isRequired = true
	}

	return isRequired
}

func (options Options) IsRequired(optionName string) (isRequired bool) {
	if option, found := options.FindOption(optionName); found && option.Flags&Required != 0 {
		isRequired = true
	}

	return isRequired
}

func (options Options) IsFlag(optionName string) (isFlag bool) {
	if option, found := options.FindOption(optionName); found && option.Flags&Flag != 0 {
		isFlag = true
	}

	return isFlag
}

func (options Options) RequiredOptions() (requiredOptions []string) {

	for _, cur := range options {
		if cur.Flags&Required != 0 {
			requiredOptions = append(requiredOptions, cur.LongOpt())
		}
	}

	return requiredOptions
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

func (options Options) Usage() (output string) {
	programName := basename(os.Args[0])
	output = "\n\n    Usage: " + programName

	for _, option := range options {
		output = output + " " + option.Usage()
	}

	output = output + "\n\n"

	return
}

func (options Options) Help(description string) (output string) {
	output = options.Usage()
	if description != "" {
		output = output + description + "\n\n"
	}

	longOptTextLength := 0

	for _, option := range options {
		if length := len(option.LongOptString()); length > longOptTextLength {
			longOptTextLength = length
		}
	}

	longOptTextLength = longOptTextLength + 2

	var argumentsString string
	var optionsString string
	var passThroughString string

	usageOpt, helpOpt := options.usageHelpOptionNames()

	for _, option := range options {
		switch {
		case option.Flags&IsPassThrough > 0:
			passThroughString = passThroughString + option.HelpText(longOptTextLength) + "\n"
		case option.Flags&IsArg > 0:
			argumentsString = argumentsString + option.HelpText(longOptTextLength) + "\n"
		case option.LongOpt() != helpOpt:
			optionsString = optionsString + option.HelpText(longOptTextLength) + "\n"
		}
	}

	if optionsString != "" {
		helpHelp := fmt.Sprintf("usage (-%s) / detailed help text (--%s)", usageOpt, helpOpt)

		if option, found := options.FindOption(helpOpt); found {
			helpHelp = option.Description
		}

		usageHelpOption := Option{fmt.Sprintf("%s|%s", helpOpt, usageOpt),
			helpHelp,
			Usage | Help | Flag, ""}
		optionsString = optionsString + usageHelpOption.HelpText(longOptTextLength) + "\n"
		output = output + "Options:\n" + optionsString + "\n"
	}

	if argumentsString != "" {
		output = output + "Arguments:\n" + argumentsString + "\n"
	}

	if passThroughString != "" {
		output = output + "Pass through arguments:\n" + passThroughString + "\n"
	}

	return
}
