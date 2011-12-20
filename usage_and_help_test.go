// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
	"testing"
	"strings"
	"os"
)

func TestUsage(t *testing.T) {
	options := Options{
		{"debug|d|DEBUG", "debug mode", Flag, true},
		{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
		{"files", "files that should be read in", IsArg, nil},
		{"secondaryports|s|SECONDARY_PORTS", "secondary ports", Optional | ExampleIsDefault, []int{5000, 5001, 5002}},
		{"instances||INSTANCES", "Instances", Required, 4},
		{"lock||LOCK", "create lock file", Flag, false},
		{"logfile||LOGFILE", "logfile", Optional | ExampleIsDefault, "/var/log/foo.log"},
		{"directories", "directories", IsArg | Optional, nil},
	}

	expected := `Usage: prog -d [-p <ports>] <files> [-s <secondaryports>] --instances=<instances> --lock [--logfile=<logfile>] [<directories>]

`

	if got := options.Usage(); got != expected {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
	}

}

func TestHelp(t *testing.T) {
	options := Options{
		{"debug|d|DEBUG", "debug mode", Flag, true},
		{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
		{"files", "Files that should be read in", IsArg, nil},
		{"secondaryports|s", "Secondary ports", Optional | ExampleIsDefault, []int{5000, 5001, 5002}},
		{"instances", "Instances", Required, 4},
		{"lock||LOCK", "create lock file", Flag, false},
		{"logfile||LOGFILE", "Logfile", Optional | ExampleIsDefault, "/var/log/foo.log"},
		{"directories", "Directories", IsArg | Optional, nil},
		{"pass through args", "arguments for subcommand", IsPassThrough, nil},
	}

	expected := `Usage: prog -d [-p <ports>] <files> [-s <secondaryports>] --instances=<instances> --lock [--logfile=<logfile>] [<directories>] -- <pass through args>

this is not a program

Options:
    -d, --debug                             debug mode; setable via $DEBUG
    -p, --ports=<ports>                     Ports (default: 3000,3001,3002); setable via $PORTS
    -s, --secondaryports=<secondaryports>   Secondary ports (default: 5000,5001,5002)
        --instances=<instances>             Instances (e.g. 4)
        --lock                              create lock file; setable via $LOCK
        --logfile=<logfile>                 Logfile (default: /var/log/foo.log); setable via $LOGFILE
    -h, --help                              usage (-h) / detailed help text (--help)

Arguments:
    <files>                                 Files that should be read in
    <directories>                           Directories

Pass through arguments:
    <pass through args>                     arguments for subcommand

`

	if got := options.Help("this is not a program"); got != expected {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
	}

}

func TestHelpNoOptions(t *testing.T) {
	options := Options{
		{"files", "Files that should be read in", IsArg, nil},
		{"directories", "Directories", IsArg | Optional, nil},
	}

	expected := `Usage: prog <files> [<directories>]

this is not a program

Arguments:
    <files>                           Files that should be read in
    <directories>                     Directories

`

	if got := options.Help("this is not a program"); got != expected {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
	}

}

func TestHelpNoArgs(t *testing.T) {
	options := Options{
		{"debug|d|DEBUG", "debug mode", Flag, true},
		{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
		{"secondaryports|s", "Secondary ports", Optional | ExampleIsDefault, []int{5000, 5001, 5002}},
		{"instances", "Instances", Required, 4},
		{"lock||LOCK", "create lock file", Flag, false},
		{"logfile||LOGFILE", "Logfile", Optional | ExampleIsDefault, "/var/log/foo.log"},
	}

	expected := `Usage: prog -d [-p <ports>] [-s <secondaryports>] --instances=<instances> --lock [--logfile=<logfile>]

this is not a program

Options:
    -d, --debug                             debug mode; setable via $DEBUG
    -p, --ports=<ports>                     Ports (default: 3000,3001,3002); setable via $PORTS
    -s, --secondaryports=<secondaryports>   Secondary ports (default: 5000,5001,5002)
        --instances=<instances>             Instances (e.g. 4)
        --lock                              create lock file; setable via $LOCK
        --logfile=<logfile>                 Logfile (default: /var/log/foo.log); setable via $LOGFILE
    -h, --help                              usage (-h) / detailed help text (--help)

`

	if got := options.Help("this is not a program"); got != expected {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
	}

}

func TestUsageAndHelpOption(t *testing.T) {
	options := Options{
		{"debug|d|DEBUG", "debug mode", Flag, true},
		{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
	}

	os.Args = []string{"prog", "barbaz", "-d", "-h", "-p5000,6000", "foobar"}
	os.Envs = []string{}
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != UsageOrHelp {
		t.Errorf("usage call did not return help error")
	}

	os.Args = []string{"prog", "barbaz", "-d", "--help", "-p5000,6000", "foobar"}
	os.Envs = []string{}
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != UsageOrHelp {
		t.Errorf("help call did not return help error")
	}

}

func TestUsageAndHelpOptionWithOwnIdentifiers(t *testing.T) {
	options := Options{
		{"chelp|c", "show usage / help", Usage | Help, nil},
		{"debug|d|DEBUG", "debug mode", Flag, true},
		{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
	}

	os.Args = []string{"prog", "barbaz", "-d", "-c", "-p5000,6000", "foobar"}
	os.Envs = []string{}
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != UsageOrHelp {
		t.Errorf("usage call did not return help error with custom '-c' for usage")
	}

	os.Args = []string{"prog", "barbaz", "-d", "--chelp", "-p5000,6000", "foobar"}
	os.Envs = []string{}
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != UsageOrHelp {
		t.Errorf("help call did not return help error with custom '--chelp' for usage")
	}

}
