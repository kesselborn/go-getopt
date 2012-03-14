// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
	"os"
	"strings"
	"syscall"
	"testing"
)

func TestUsage(t *testing.T) {
	options := Options{
		"",
		Definitions{{"debug|d|DEBUG", "debug mode", Flag, true},
			{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
			{"files", "files that should be read in", IsArg, nil},
			{"secondaryports|s|SECONDARY_PORTS", "secondary ports", Optional | ExampleIsDefault, []int{5000, 5001, 5002}},
			{"instances||INSTANCES", "Instances", Required, 4},
			{"lock||LOCK", "create lock file", Flag, false},
			{"logfile||LOGFILE", "logfile", Optional | ExampleIsDefault, "/var/log/foo.log"},
			{"directories", "directories", IsArg | Optional, nil},
			{"command", "command", IsPassThrough | Required, nil},
			{"args", "command's args", IsPassThrough | Optional, nil}},
	}

	os.Args = []string{"prog"}
	expected := `Usage: prog -d [-p <ports>] <files> [-s <secondaryports>] --instances=<instances> --lock [--logfile=<logfile>] [<directories>] -- <command> [<args>]

`

	if got := options.Usage(); got != expected {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
	}

}

func TestUsageWithDifferentProgname(t *testing.T) {
	options := Options{
		{"debug|d|DEBUG", "debug mode", Flag, true},
		{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
		{"files", "files that should be read in", IsArg, nil},
		{"secondaryports|s|SECONDARY_PORTS", "secondary ports", Optional | ExampleIsDefault, []int{5000, 5001, 5002}},
		{"instances||INSTANCES", "Instances", Required, 4},
		{"lock||LOCK", "create lock file", Flag, false},
		{"logfile||LOGFILE", "logfile", Optional | ExampleIsDefault, "/var/log/foo.log"},
		{"directories", "directories", IsArg | Optional, nil},
		{"command", "command", IsPassThrough | Required, nil},
		{"args", "command's args", IsPassThrough | Optional, nil},
	}

	os.Args = []string{"prog"}
	expected := `Usage: otherprogname -d [-p <ports>] <files> [-s <secondaryports>] --instances=<instances> --lock [--logfile=<logfile>] [<directories>] -- <command> [<args>]

`

	if got := options.UsageCustomArg0("otherprogname"); got != expected {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
	}

}

func TestHelpWithDifferentProgname(t *testing.T) {
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

	os.Args = []string{"prog"}
	expected := `Usage: otherprogname -d [-p <ports>] <files> [-s <secondaryports>] --instances=<instances> --lock [--logfile=<logfile>] [<directories>] -- <pass through args>

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

	if got := options.HelpCustomArg0("this is not a program", "otherprogname"); got != expected {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
	}

}

func TestHelp(t *testing.T) {
	options := Options{
		"this is not a program",
		Definitions{{"debug|d|DEBUG", "debug mode", Flag, true},
			{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
			{"files", "Files that should be read in", IsArg, nil},
			{"secondaryports|s", "Secondary ports", Optional | ExampleIsDefault, []int{5000, 5001, 5002}},
			{"instances", "Instances", Required, 4},
			{"lock||LOCK", "create lock file", Flag, false},
			{"logfile||LOGFILE", "Logfile", Optional | ExampleIsDefault, "/var/log/foo.log"},
			{"directories", "Directories", IsArg | Optional, nil},
			{"pass through args", "arguments for subcommand", IsPassThrough, nil}},
	}

	os.Args = []string{"prog"}
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

	if got := options.Help(); got != expected {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
	}

}

func TestHelpNoOptions(t *testing.T) {
	options := Options{
		"this is not a program",
		Definitions{{"files", "Files that should be read in", IsArg, nil},
			{"directories", "Directories", IsArg | Optional, nil}},
	}

	os.Args = []string{"prog"}
	expected := `Usage: prog <files> [<directories>]

this is not a program

Arguments:
    <files>                           Files that should be read in
    <directories>                     Directories

`

	if got := options.Help(); got != expected {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
	}

}

func TestHelpNoArgs(t *testing.T) {
	options := Options{
		"this is not a program",
		Definitions{{"debug|d|DEBUG", "debug mode", Flag, true},
			{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
			{"secondaryports|s", "Secondary ports", Optional | ExampleIsDefault, []int{5000, 5001, 5002}},
			{"instances", "Instances", Required, 4},
			{"lock||LOCK", "create lock file", Flag, false},
			{"logfile||LOGFILE", "Logfile", Optional | ExampleIsDefault, "/var/log/foo.log"}},
	}

	os.Args = []string{"prog"}
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

	if got := options.Help(); got != expected {
		t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
	}

}

func TestUsageAndHelpOption(t *testing.T) {
	options := Options{
		"",
		Definitions{{"debug|d|DEBUG", "debug mode", Flag, true},
			{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}}},
	}

	os.Args = []string{"prog", "barbaz", "-d", "-h", "-p5000,6000", "foobar"}
	syscall.Clearenv()
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != WantsUsage {
		t.Errorf("usage call did not return help error")
	}

	os.Args = []string{"prog", "barbaz", "-d", "--help", "-p5000,6000", "foobar"}
	syscall.Clearenv()
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != WantsHelp {
		t.Errorf("help call did not return help error")
	}
}

func TestUsageAndHelpOptionInPassThrough(t *testing.T) {
	options := Options{
		"",
		Definitions{{"debug|d|DEBUG", "debug mode", Flag, true},
			{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}}},
	}

	os.Args = []string{"prog", "barbaz", "--", "-h"}
	syscall.Clearenv()
	if _, _, _, err := options.ParseCommandLine(); err != nil {
		t.Errorf("usage option in pass through triggered WantsUsage")
	}

	os.Args = []string{"prog", "barbaz", "--", "--help"}
	syscall.Clearenv()
	if _, _, _, err := options.ParseCommandLine(); err != nil {
		t.Errorf("help option in pass through triggered WantsUsage")
	}

}

func TestUsageAndHelpOptionWithOwnIdentifiers(t *testing.T) {
	options := Options{
		"",
		Definitions{{"chelp|c", "show usage / help", Usage | Help, nil},
			{"debug|d|DEBUG", "debug mode", Flag, true},
			{"ports|p|PORTS", "Ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}}},
	}

	os.Args = []string{"prog", "barbaz", "-d", "-c", "-p5000,6000", "foobar"}
	syscall.Clearenv()
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != WantsUsage {
		t.Errorf("usage call did not return help error with custom '-c' for usage")
	}

	os.Args = []string{"prog", "barbaz", "-d", "--chelp", "-p5000,6000", "foobar"}
	syscall.Clearenv()
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != WantsHelp {
		t.Errorf("help call did not return help error with custom '--chelp' for usage")
	}

}
