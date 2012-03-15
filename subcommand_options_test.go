// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
	"os"
	"testing"
)

func equalOptionsArray(array1 Options, array2 Options) (equal bool) {
	if len(array1) == len(array2) {
		for i := 0; i < len(array1); i++ {
			if array1[i] != array2[i] {
				break
			}
		}
		equal = true
	}

	return
}

func TestSubcommandOptionsConverter(t *testing.T) {
	sco := SubCommandOptions{
		"*": {
			{"command", "command to execute", IsSubcommand, ""}},
		"getenv": {
			{"name", "app's name", IsArg | Required, ""},
			{"key", "environment variable's name", IsArg | Required, ""}},
		"register": {
			{"name|n", "app's name", IsArg | Required, ""},
			{"deploytype|t", "deploy type (one of mount, bazapta, lxc)", Optional | ExampleIsDefault, "lxc"}},
	}

	expectedGetenvOptions := Options{
		{"command", "command to execute", IsSubcommand, ""},
		{"name", "app's name", IsArg | Required, ""},
		{"key", "environment variable's name", IsArg | Required, ""},
	}

	expectedRegisterOptions := Options{
		{"command", "command to execute", IsSubcommand, ""},
		{"name|n", "app's name", IsArg | Required, ""},
		{"deploytype|t", "deploy type (one of mount, bazapta, lxc)", Optional | ExampleIsDefault, "lxc"},
	}

	if _, err := sco.flattenToOptions("getenv"); err != nil {
		t.Errorf("conversion SubCommandOptions -> Options failed (getenv); \nGot the following error: %s", err.Message)
	}

	if options, _ := sco.flattenToOptions("getenv"); equalOptionsArray(options, expectedGetenvOptions) == false {
		t.Errorf("conversion SubCommandOptions -> Options failed (getenv); \nGot\n\t#%#v#\nExpected:\n\t#%#v#\n", options, expectedGetenvOptions)
	}

	if _, err := sco.flattenToOptions("register"); err != nil {
		t.Errorf("conversion SubCommandOptions -> Options failed (register); \nGot the following error: %s", err.Message)
	}

	if options, _ := sco.flattenToOptions("register"); equalOptionsArray(options, expectedRegisterOptions) == false {
		t.Errorf("conversion SubCommandOptions -> Options failed (register); \nGot\n\t#%#v#\nExpected:\n\t#%#v#\n", options, expectedGetenvOptions)
	}

	if _, err := sco.flattenToOptions("nonexistantsubcommand"); err.ErrorCode != UnknownSubcommand {
		t.Errorf("non existant sub command didn't throw error")
	}

}

func TestSubcommandOptionsSubCommandFinder(t *testing.T) {
	sco := SubCommandOptions{
		"*": {
			{"command", "command to execute", IsSubcommand, ""},
			{"foo|f", "some arg", Optional, ""}},
		"getenv": {
			{"name", "app's name", IsArg | Required, ""},
			{"key", "environment variable's name", IsArg | Required, ""}},
	}

	os.Args = []string{"prog", "getenv"}
	if command, _ := sco.findSubcommand(); command != "getenv" {
		t.Errorf("did not correctly find subcommand getenv")
	}

	os.Args = []string{"prog", "-f", "bar", "getenv", "name", "key"}
	if command, _ := sco.findSubcommand(); command != "getenv" {
		t.Errorf("did not correctly find subcommand getenv")
	}

	os.Args = []string{"prog"}
	if _, err := sco.findSubcommand(); err == nil || err.ErrorCode != NoSubcommand {
		t.Errorf("did not throw error on unknown subcommand")
	}
}

func TestSubcommandOptionsParser(t *testing.T) {
	sco := SubCommandOptions{
		"*": {
			{"command", "command to execute", IsSubcommand, ""},
			{"foo|f", "some arg", Optional, ""}},
		"getenv": {
			{"bar|b", "some arg", Optional, ""},
			{"name", "app's name", IsArg | Required, ""},
			{"key", "environment variable's name", IsArg | Required, ""}},
	}

	os.Args = []string{"prog", "-fbar", "getenv", "--bar=foo", "foo", "bar"}
	options, arguments, _, _ := sco.ParseCommandLine()

	if options["foo"].String != "bar" {
		t.Errorf("SubCommandOptions parsing: failed to correctly parse option: Expected: bar, Got: " + options["foo"].String)
	}

	if options["bar"].String != "foo" {
		t.Errorf("SubCommandOptions parsing: failed to correctly parse option: Expected:  foo, Got: " + options["foo"].String)
	}

	if arguments[0] != "getenv" {
		t.Errorf("SubCommandOptions parsing: failed to correctly parse sub command: Expected: getenv, Got: " + arguments[0])
	}

	if arguments[1] != "foo" {
		t.Errorf("SubCommandOptions parsing: failed to correctly parse arg1: Expected: foo, Got: " + arguments[1])
	}

	if arguments[2] != "bar" {
		t.Errorf("SubCommandOptions parsing: failed to correctly parse arg2: Expected: bar, Got: " + arguments[2])
	}
}
