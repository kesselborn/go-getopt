// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
	"os"
	"strings"
	"testing"
)

func equalOptions(options1 Options, options2 Options) (equal bool) {
	if len(options1.definitions) == len(options2.definitions) && options1.description == options2.description {
		for i := 0; i < len(options1.definitions); i++ {
			if options1.definitions[i] != options2.definitions[i] {
				goto loopend
			}
		}
		equal = true
	}
loopend:

	return
}

func TestSubcommandOptionsConverter(t *testing.T) {
	sco := SubCommandOptions{
		Options{
			"global description",
			Definitions{
				{"command", "command to execute", IsSubcommand, ""},
			},
		},
		SubCommands{
			"getenv": {
				"getenv description",
				Definitions{
					{"name", "app's name", IsArg | Required, ""},
					{"key", "environment variable's name", IsArg | Required, ""},
				},
			},
			"register": {
				"register description",
				Definitions{
					{"name|n", "app's name", IsArg | Required, ""},
					{"deploytype|t", "deploy type (one of mount, bazapta, lxc)", Optional | ExampleIsDefault, "lxc"},
				},
			},
		},
	}

	expectedGetenvOptions := Options{
		"getenv description",
		Definitions{{"command", "command to execute", IsSubcommand, ""},
			{"name", "app's name", IsArg | Required, ""},
			{"key", "environment variable's name", IsArg | Required, ""}},
	}

	expectedRegisterOptions := Options{
		"register description",
		Definitions{{"command", "command to execute", IsSubcommand, ""},
			{"name|n", "app's name", IsArg | Required, ""},
			{"deploytype|t", "deploy type (one of mount, bazapta, lxc)", Optional | ExampleIsDefault, "lxc"}},
	}

	if _, err := sco.flattenToOptions("getenv"); err != nil {
		t.Errorf("conversion SubCommandOptions -> Options failed (getenv); \nGot the following error: %s", err.Message)
	}

	if options, _ := sco.flattenToOptions("getenv"); equalOptions(options, expectedGetenvOptions) == false {
		t.Errorf("conversion SubCommandOptions -> Options failed (getenv); \nGot\n\t#%#v#\nExpected:\n\t#%#v#\n", options, expectedGetenvOptions)
	}

	if _, err := sco.flattenToOptions("register"); err != nil {
		t.Errorf("conversion SubCommandOptions -> Options failed (register); \nGot the following error: %s", err.Message)
	}

	if options, _ := sco.flattenToOptions("register"); equalOptions(options, expectedRegisterOptions) == false {
		t.Errorf("conversion SubCommandOptions -> Options failed (register); \nGot\n\t#%#v#\nExpected:\n\t#%#v#\n", options, expectedGetenvOptions)
	}

	if _, err := sco.flattenToOptions("nonexistantsubcommand"); err.ErrorCode != UnknownSubcommand {
		t.Errorf("non existant sub command didn't throw error")
	}

}

func TestSubcommandOptionsSubCommandFinder(t *testing.T) {
	sco := SubCommandOptions{
		Options{
			"global description",
			Definitions{
				{"command", "command to execute", IsSubcommand, ""},
				{"foo|f", "some arg", Optional, ""},
			},
		},
		SubCommands{
			"getenv": {
				"getenv description",
				Definitions{
					{"name", "app's name", IsArg | Required, ""},
					{"key", "environment variable's name", IsArg | Required, ""},
				},
			},
		},
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
		Options{
			"global description",
			Definitions{
				{"command", "command to execute", IsSubcommand, ""},
				{"foo|f", "some arg", Optional, ""},
			},
		},
		SubCommands{
			"getenv": {
				"getenv description",
				Definitions{
					{"bar|b", "some arg", Optional, ""},
					{"name", "app's name", IsArg | Required, ""},
					{"key", "environment variable's name", IsArg | Required, ""},
				},
			},
		},
	}

	os.Args = []string{"prog", "-fbar", "getenv", "--bar=foo", "foo", "bar"}
	scope, options, arguments, _, _ := sco.ParseCommandLine()

	if scope != "getenv" {
		t.Errorf("SubCommandOptions parsing: failed to correctly parse scope: Expected: getenv, Got: " + scope)
	}

	if options["foo"].String != "bar" {
		t.Errorf("SubCommandOptions parsing: failed to correctly parse option: Expected: bar, Got: " + options["foo"].String)
	}

	if options["bar"].String != "foo" {
		t.Errorf("SubCommandOptions parsing: failed to correctly parse option: Expected:  foo, Got: " + options["foo"].String)
	}

	if scope != "getenv" {
		t.Errorf("SubCommandOptions parsing: failed to correctly parse sub command: Expected: getenv, Got: " + scope)
	}

	if arguments[0] != "foo" {
		t.Errorf("SubCommandOptions parsing: failed to correctly parse arg1: Expected: foo, Got: " + arguments[0])
	}

	if arguments[1] != "bar" {
		t.Errorf("SubCommandOptions parsing: failed to correctly parse arg2: Expected: bar, Got: " + arguments[1])
	}
}

func TestErrorMessageForMissingArgs(t *testing.T) {
	sco := SubCommandOptions{
		Options{
			"global description",
			Definitions{
				{"foo|f", "some arg", Optional, ""},
				{"command", "command to execute", IsSubcommand, ""},
			},
		},
		SubCommands{
			"getenv": {
				"getenv description",
				Definitions{
					{"bar|b", "some arg", Optional, ""},
					{"name", "app's name", IsArg | Required, ""},
					{"key", "environment variable's name", IsArg | Required, ""},
				},
			},
		},
	}

	os.Args = []string{"prog", "getenv"}
	_, _, _, _, err := sco.ParseCommandLine()

	if err == nil {
		t.Errorf("missing arg did not raise error")
	}

	if expected := "Missing required argument <name>"; err.Message != expected {
		t.Errorf("Error handling for missing arguments is messed up:\n\tGot     : " + err.Message + "\n\tExpected: " + expected)
	}

}

func TestSubCommandHelp(t *testing.T) {
	sco := SubCommandOptions{
		Options{
			"global description",
			Definitions{
				{"foo|f", "some arg", Optional, ""},
				{"command", "command to execute", IsSubcommand, ""},
			},
		},
		SubCommands{
			"getenv": {
				"getenv description",
				Definitions{
					{"bar|b", "some arg", Optional, ""},
					{"name", "app's name", IsArg | Required, ""},
					{"key", "environment variable's name", IsArg | Required, ""},
				},
			},
			"register": {
				"register description",
				Definitions{
					{"deploytype|t", "deploy type (one of mount, bazapta, lxc)", NoLongOpt | Optional | ExampleIsDefault, "lxc"},
					{"name|n", "app's name", IsArg | Required, ""},
				},
			},
		},
	}

	os.Args = []string{"prog"}
	expectedHelp := `Usage: prog [-f <foo>] <command>

global description

Options:
    -f, --foo=<foo>           some arg
    -h, --help                usage (-h) / detailed help text (--help)

Available commands:
    getenv                    getenv description
    register                  register description

`
	expectedUsage := `Usage: prog [-f <foo>] <command>

`

	if got := sco.Help("*"); got != expectedHelp {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
	}

	if got := sco.Usage("*"); got != expectedUsage {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
	}

	os.Args = []string{"prog", "register"}
	expectedHelp = `Usage: prog register [-t <deploytype>] <name>

register description

Options:
    -t <deploytype>     deploy type (one of mount, bazapta, lxc) (default: lxc)
    -h, --help          usage (-h) / detailed help text (--help)

Arguments:
    <name>              app's name

`

	expectedUsage = `Usage: prog register [-t <deploytype>] <name>

`

	if got := sco.Help("register"); got != expectedHelp {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
	}

	if got := sco.Usage("register"); got != expectedUsage {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
	}

}
