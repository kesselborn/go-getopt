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

func testingSubSubDefinitions() (ssco SubSubCommandOptions) {
	ssco = SubSubCommandOptions{
		Options{
			"global description",
			Definitions{
				{"config|c", "config file", IsConfigFile | ExampleIsDefault, "/etc/visor.conf"},
				{"server|s", "doozer server", Optional, ""},
				{"scope", "scope", IsSubCommand, ""},
			},
		},
		Scopes{
			"app": {
				Options{
					"app description",
					Definitions{
						{"foo|f", "a param", Optional, ""},
						{"command", "command to execute", IsSubCommand, ""},
					},
				},
				SubCommands{
					"getenv": {
						"app getenv description",
						Definitions{
							{"key", "environment variable's name", IsArg | Required, ""},
						},
					},
					"setenv": {
						"app setenv description",
						Definitions{
							{"persist|p", "persist this", Optional | Flag | NoLongOpt, ""},
							{"alias|a", "alias name", Optional, ""},
							{"name", "app name", IsArg | Required, ""},
							{"key", "env key", IsArg | Required, ""},
							{"value", "env value", IsArg | Optional, ""},
						},
					},
				},
			},
			"revision": {
				Options{
					"app revision description",
					Definitions{
						{"rev|r", "revision", IsArg | Required, ""},
						{"command", "command to execute", IsSubCommand, ""},
					},
				},
				SubCommands{
					"list": {
						"list revisions",
						Definitions{
							{"all|a", "long list output", Flag, ""},
						},
					},
				},
			},
		},
	}

	return
}
func TestSubSubCommandOptionsConverter(t *testing.T) {
	ssco := testingSubSubDefinitions()

	expectedAppSubOptions := SubCommandOptions{
		Options{
			"app description",
			Definitions{
				{"config|c", "config file", IsConfigFile | ExampleIsDefault, "/etc/visor.conf"},
				{"server|s", "doozer server", Optional, ""},
				{"scope", "scope", IsSubCommand, ""},
				{"foo|f", "a param", Optional, ""},
				{"command", "command to execute", IsSubCommand, ""},
			},
		},
		SubCommands{
			"getenv": {
				"app getenv description",
				Definitions{
					{"key", "environment variable's name", IsArg | Required, ""},
				},
			},
			"setenv": {
				"app setenv description",
				Definitions{
					{"persist|p", "persist this", Optional | Flag | NoLongOpt, ""},
					{"alias|a", "alias name", Optional, ""},
					{"name", "app name", IsArg | Required, ""},
					{"key", "env key", IsArg | Required, ""},
					{"value", "env value", IsArg | Optional, ""},
				},
			},
		},
	}

	expectedRevisionOptions := SubCommandOptions{
		Options{
			"app revision description",
			Definitions{
				{"config|c", "config file", IsConfigFile | ExampleIsDefault, "/etc/visor.conf"},
				{"server|s", "doozer server", Optional, ""},
				{"scope", "scope", IsSubCommand, ""},
				{"rev|r", "revision", IsArg | Required, ""},
				{"command", "command to execute", IsSubCommand, ""},
			},
		},
		SubCommands{
			"list": {
				"list revisions",
				Definitions{
					{"all|a", "long list output", Flag, ""},
				},
			},
		},
	}

	if _, err := ssco.flattenToSubCommandOptions("app"); err != nil {
		t.Errorf("conversion SuSubCommandOptions -> SubCommandOptions failed (app); \nGot the following error: %s", err.Message)
	}

	if sco, _ := ssco.flattenToSubCommandOptions("app"); equalSubCommandOptions(sco, expectedAppSubOptions) == false {
		t.Errorf("conversion SubSubCommandOptions -> SubCommandOptions failed (app); \nGot\n\t#%#v#\nExpected:\n\t#%#v#\n", sco, expectedAppSubOptions)
	}

	if _, err := ssco.flattenToSubCommandOptions("revision"); err != nil {
		t.Errorf("conversion SuSubCommandOptions -> SubCommandOptions failed (revision); \nGot the following error: %s", err.Message)
	}

	if sco, _ := ssco.flattenToSubCommandOptions("revision"); equalSubCommandOptions(sco, expectedRevisionOptions) == false {
		t.Errorf("conversion SubSubCommandOptions -> SubCommandOptions failed (revision); \nGot\n\t#%#v#\nExpected:\n\t#%#v#\n", sco, expectedRevisionOptions)
	}

	if _, err := ssco.flattenToSubCommandOptions("nonexistantsubcommand"); err.ErrorCode != UnknownSubCommand {
		t.Errorf("non existant sub command didn't throw error")
	}

	expectedAppGetEnvOptions := Options{
		"app getenv description",
		Definitions{
			{"config|c", "config file", IsConfigFile | ExampleIsDefault, "/etc/visor.conf"},
			{"server|s", "doozer server", Optional, ""},
			{"scope", "scope", IsSubCommand, ""},
			{"foo|f", "a param", Optional, ""},
			{"command", "command to execute", IsSubCommand, ""},
			{"key", "environment variable's name", IsArg | Required, ""},
		},
	}

	if _, err := ssco.flattenToOptions("app", "getenv"); err != nil {
		t.Errorf("conversion SubSubCommandOptions -> Options failed (app/getenv); \nGot the following error: %s", err.Message)
	}

	if options, _ := ssco.flattenToOptions("app", "getenv"); equalOptions(options, expectedAppGetEnvOptions) == false {
		t.Errorf("conversion SubCommandOptions -> Options failed (app/getenv); \nGot\n\t#%#v#\nExpected:\n\t#%#v#\n", options, expectedAppGetEnvOptions)
	}

}

func TestSubSubCommandScopeFinder(t *testing.T) {
	ssco := testingSubSubDefinitions()

	os.Args = []string{"prog", "app"}
	if command, _ := ssco.findScope(); command != "app" {
		t.Errorf("did not correctly find subcommand app")
	}

	os.Args = []string{"prog", "-s", "10.20.30.40", "app", "getenv", "key"}
	if command, _ := ssco.findScope(); command != "app" {
		t.Errorf("did not correctly find subcommand app (w/ other options)")
	}

	os.Args = []string{"prog"}
	if _, err := ssco.findScope(); err == nil || err.ErrorCode != NoScope {
		t.Errorf("did not throw error on unknown subcommand")
	}
}

func TestSubSubCommandSubCommand(t *testing.T) {
	ssco := testingSubSubDefinitions()

	os.Args = []string{"prog", "app", "getenv"}
	if scope, command, _ := ssco.findScopeAndSubCommand(); scope != "app" || command != "getenv" {
		t.Errorf("did not correctly find subcommand app / getenv (got: " + scope + " / " + command + ")")
	}

	os.Args = []string{"prog", "-s", "10.20.30.40", "app", "-ffoo", "getenv", "key"}
	if _, _, err := ssco.findScopeAndSubCommand(); err != nil {
		t.Errorf("did not correctly find subcommand app / getenv; Error message: " + err.Message)
	}

	if scope, command, _ := ssco.findScopeAndSubCommand(); scope != "app" || command != "getenv" {
		t.Errorf("did not correctly find subcommand app / getenv (got: " + scope + " / " + command + ")")
	}

	os.Args = []string{"prog"}
	if _, _, err := ssco.findScopeAndSubCommand(); err == nil || err.ErrorCode != NoScope {
		t.Errorf("did not throw error on missing scope")
	}

	os.Args = []string{"prog", "app"}
	if _, _, err := ssco.findScopeAndSubCommand(); err == nil || err.ErrorCode != NoSubCommand {
		t.Errorf("did not throw error on missing subcommand")
	}

	os.Args = []string{"prog", "unknownscope"}
	if _, _, err := ssco.findScopeAndSubCommand(); err == nil || err.ErrorCode != UnknownScope {
		t.Errorf("did not throw error on unknown scope")
	}

	os.Args = []string{"prog", "app", "unknowncommand"}
	if _, _, err := ssco.findScopeAndSubCommand(); err == nil || err.ErrorCode != UnknownSubCommand {
		t.Errorf("did not throw error on unknown subcommand")
	}
}

func TestSubSubCommandOptionsParser(t *testing.T) {
	ssco := testingSubSubDefinitions()

	os.Args = []string{"prog", "-sfoo.com", "app", "-ffoo", "setenv", "-p", "-aFOO", "name", "val", "--", "pass", "through"}
	scope, command, options, arguments, passThrough, err := ssco.ParseCommandLine()

	if err != nil {
		t.Errorf("Got an unexpected error while parsing SubSubCommandOptions: " + err.Message)
	}

	if scope != "app" {
		t.Errorf("SubSubCommandOptions parsing: failed to correctly parse scope: Expected: app, Got: " + scope)
	}

	if command != "setenv" {
		t.Errorf("SubSubCommandOptions parsing: failed to correctly parse command: Expected: setenv, Got: " + scope)
	}

	if options["server"].String != "foo.com" {
		t.Errorf("SubSubCommandOptions parsing: failed to correctly parse option: Expected: foo.com, Got: " + options["server"].String)
	}

	if options["alias"].String != "FOO" {
		t.Errorf("SubSubCommandOptions parsing: failed to correctly parse option: Expected:  FOO, Got: " + options["foo"].String)
	}

	if arguments[0] != "name" {
		t.Errorf("SubSubCommandOptions parsing: failed to correctly parse arg1: Expected: name, Got: " + arguments[0])
	}

	if arguments[1] != "val" {
		t.Errorf("SubSubCommandOptions parsing: failed to correctly parse arg2: Expected: val, Got: " + arguments[1])
	}

	if passThrough[0] != "pass" {
		t.Errorf("SubSubCommandOptions parsing: failed to correctly parse pass through[0]: Expected: pass, Got: " + passThrough[0])
	}

	if passThrough[1] != "through" {
		t.Errorf("SubSubCommandOptions parsing: failed to correctly parse pass through[0]: Expected: through, Got: " + passThrough[0])
	}
}

func TestErrorMessages(t *testing.T) {
	ssco := testingSubSubDefinitions()

	os.Args = []string{"prog"}
	_, _, _, _, _, err := ssco.ParseCommandLine()

	if err == nil {
		t.Errorf("NoScope error not triggered on missing")
	}

	if err.ErrorCode != NoScope {
		t.Errorf("NoScope error not triggered on missing")
	}

	os.Args = []string{"prog", "app"}
	_, _, _, _, _, err = ssco.ParseCommandLine()

	if err == nil {
		t.Errorf("NoCommand error not triggered on missing")
	}

	if err.ErrorCode != NoSubCommand {
		t.Errorf("NoCommand error not triggered on missing")
	}

	os.Args = []string{"prog", "--help"}
	_, _, parsedOptions, _, _, _ := ssco.ParseCommandLine()

	if helpOption, present := parsedOptions["help"]; !present || helpOption.String != "help" {
		t.Errorf("Wants help for global command set help option")
	}

	os.Args = []string{"prog", "app", "-h"}
	_, _, parsedOptions, _, _, _ = ssco.ParseCommandLine()

	if helpOption, present := parsedOptions["help"]; !present || helpOption.String != "usage" {
		t.Errorf("Wants usage for global command did set usage option correctly")
	}

	os.Args = []string{"prog", "app", "--help"}
	_, _, parsedOptions, _, _, _ = ssco.ParseCommandLine()

	if helpOption, present := parsedOptions["help"]; !present || helpOption.String != "help" {
		t.Errorf("Wants help for global command did not set help option correctly")
	}

}

func TestWantsHelpAndUsage(t *testing.T) {
	ssco := testingSubSubDefinitions()

	os.Args = []string{"prog", "-h"}
	_, _, parsedOptions, _, _, _ := ssco.ParseCommandLine()

	if helpOption, present := parsedOptions["help"]; !present || helpOption.String != "usage" {
		t.Errorf("Wants usage for global command did not set usage option correctly")
	}

	os.Args = []string{"prog", "--help"}
	_, _, parsedOptions, _, _, _ = ssco.ParseCommandLine()

	if helpOption, present := parsedOptions["help"]; !present || helpOption.String != "help" {
		t.Errorf("Wants help for global command did not help option correctly")
	}

	os.Args = []string{"prog", "app", "-h"}
	_, _, parsedOptions, _, _, _ = ssco.ParseCommandLine()

	if helpOption, present := parsedOptions["help"]; !present || helpOption.String != "usage" {
		t.Errorf("Wants usage for global command did not set usage option correclty")
	}

	os.Args = []string{"prog", "app", "--help"}
	_, _, parsedOptions, _, _, _ = ssco.ParseCommandLine()

	if helpOption, present := parsedOptions["help"]; !present || helpOption.String != "help" {
		t.Errorf("Wants help for global command did not set help option correctly")
	}

}

func TestErrorMessageForMissingArgsInSsco(t *testing.T) {
	ssco := testingSubSubDefinitions()

	os.Args = []string{"prog", "-sfoo.com", "app", "-ffoo", "setenv"}
	_, _, _, _, _, err := ssco.ParseCommandLine()

	if err == nil {
		t.Errorf("missing arg did not raise error")
	}

	if expected := "Missing required argument <name>"; err.Message != expected {
		t.Errorf("Error handling for missing arguments is messed up:\n\tGot     : " + err.Message + "\n\tExpected: " + expected)
	}

}

func TestSubSubCommandHelpForGlobalCommand(t *testing.T) {
	ssco := testingSubSubDefinitions()

	os.Args = []string{"prog"}

	expectedUsage := `Usage: prog [-c <config>] [-s <server>] <scope>

`
	expectedHelp := expectedUsage + `global description

Options:
    -c, --config=<config>   config file (default: /etc/visor.conf)
    -s, --server=<server>   doozer server
    -h, --help              usage (-h) / detailed help text (--help)

Available scopes:
    app                     app description
    revision                app revision description

`

	if got := ssco.Usage(); got != expectedUsage {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedUsage, " ", "_", -1) + "|\n")
	}

	if got := ssco.Help(); got != expectedHelp {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
	}
}

func TestSubSubCommandHelpForSubCommand(t *testing.T) {
	ssco := testingSubSubDefinitions()
	os.Args = []string{"prog", "app"}

	expectedUsage := `Usage: prog app [-f <foo>] <command>

`
	expectedHelp := expectedUsage + `app description

Options:
    -f, --foo=<foo>           a param
    -h, --help                usage (-h) / detailed help text (--help)

Available commands:
    getenv                    app getenv description
    setenv                    app setenv description

`

	if got := ssco.Help(); got != expectedHelp {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
	}

	if got := ssco.Usage(); got != expectedUsage {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
	}

	os.Args = []string{"prog", "app", "--help"}
	if got := ssco.Help(); got != expectedHelp {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
	}

	os.Args = []string{"prog", "app", "-h"}
	if got := ssco.Usage(); got != expectedUsage {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
	}

}

func TestSubSubCommandHelpForSubSubCommand(t *testing.T) {
	ssco := testingSubSubDefinitions()
	os.Args = []string{"prog", "app", "setenv"}

	expectedUsage := `Usage: prog app setenv [-p] [-a <alias>] <name> <key> [<value>]

`
	expectedHelp := expectedUsage + `app setenv description

Options:
    -p                    persist this
    -a, --alias=<alias>   alias name
    -h, --help            usage (-h) / detailed help text (--help)

Arguments:
    <name>                app name
    <key>                 env key
    <value>               env value

`

	if got := ssco.Help(); got != expectedHelp {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
	}

	if got := ssco.Usage(); got != expectedUsage {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
	}

}
