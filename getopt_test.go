// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
	"fmt"
	"os"
	"testing"
)

func TestShortOptionsFlagsParsing(t *testing.T) {
	options := Options{
		"",
		Definitions{
			{"debug|d", "debug mode", Flag, ""},
			{"verbose|v", "verbose mode", Flag, ""},
			{"dryrun|D", "dry run only", Flag, ""},
		},
	}

	os.Args = []string{"prog", "-d"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true {
		t.Errorf("debug flag was not set")
	}

	os.Args = []string{"prog", "-d"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != false {
		t.Errorf("verbose flag was not set to false by default")
	}

	os.Args = []string{"prog", "-d", "-v"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true {
		t.Errorf("did not recognize all flags")
	}

}

func TestShortOptionRequiredParsing(t *testing.T) {
	options := Options{"", Definitions{{"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", Required, ""}}}

	os.Args = []string{"prog", "-m", "heartbeat"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["method"].String != "heartbeat" {
		t.Errorf("method optioned wasn't parsed correctly expected 'heartbeat', was '" + opts["method"].String + "'")
	}

	os.Args = []string{"prog", "-mheartbeat"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["method"].String != "heartbeat" {
		t.Errorf("method optioned wasn't parsed correctly expected 'heartbeat', was '" + opts["method"].String + "'")
	}

	os.Args = []string{"prog", "-m"}
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != MissingValue {
		t.Errorf("required option without value did not raise error")
	}

	os.Args = []string{"prog", "-m", "-x"}
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != MissingValue {
		t.Errorf("required option without value did not raise error")
	}

	os.Args = []string{"prog", ""}
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != MissingOption {
		t.Errorf("required option wasn't set")
	}

}

func TestRequiredArgument(t *testing.T) {
	options := Options{
		"",
		Definitions{
			{"source", "original file name", Required | IsArg, ""},
			{"destination", "destination file name", Required | IsArg, ""},
		},
	}

	os.Args = []string{"prog"}
	if _, _, _, err := options.ParseCommandLine(); err == nil {
		t.Errorf("missing required argument did not raise an error")
	}

	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != MissingArgument {
		t.Errorf("missing required argument did raise wrong error")
	}

	if _, _, _, err := options.ParseCommandLine(); err == nil || err.Error() != "Missing required argument <source>" {
		t.Errorf("wrong error message for missing required argument:\n\texpected: %s\n\tgot     : %s", "Missing required argument <source>", err.Error())
	}

	os.Args = []string{"prog", "file1"}
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.Error() != "Missing required argument <destination>" {
		t.Errorf("wrong error message for missing required argument:\n\texpected: %s\n\tgot     : %s", "Missing required argument <destination>", err.Error())
	}

	os.Args = []string{"prog", "file1", "file1.bak"}
	if _, arguments, _, err := options.ParseCommandLine(); err != nil || len(arguments) != 2 || arguments[0] != "file1" || arguments[1] != "file1.bak" {
		t.Errorf("required arguments weren't set correctly (2), got: %#v", arguments)
	}

}

func TestConcatenatedOptionsParsingSimple(t *testing.T) {
	options := Options{
		"",
		Definitions{
			{"debug|d", "debug mode", Flag, true},
			{"verbose|v", "verbose mode", Flag, true},
			{"dryrun|D", "dry run only", Flag, true},
			{"logfile|l", "log file", Optional, ""},
			{"mode|m", "operating mode", Required, ""},
		},
	}

	os.Args = []string{"prog", "-dv"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true {
		t.Errorf("did not recognize all flags when concatenation options (2 options)")
	}

	os.Args = []string{"prog", "-dvD"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true || opts["dryrun"].Bool != true {
		t.Errorf("did not recognize all flags when concatenation options (3 options)")
	}

	os.Args = []string{"prog", "-dvD"}
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != MissingOption {
		t.Errorf("did not recognize a missing required option in concatenation mode")
	}

	os.Args = []string{"prog", "-Dl"}
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != MissingValue {
		t.Errorf("did not realize that I missed a value")
	}

	os.Args = []string{"prog", "-Dl", "-d"}
	if _, _, _, err := options.ParseCommandLine(); err == nil || err.ErrorCode != MissingValue {
		t.Errorf("did not realize that I missed a value")
	}

}

func TestConcatenatedOptionsParsingWithStringValueOptionAtTheEnd(t *testing.T) {
	options := Options{
		"",
		Definitions{
			{"debug|d", "debug mode", Flag, true},
			{"verbose|v", "verbose mode", Flag, true},
			{"dryrun|D", "dry run only", Flag, true},
			{"logfile|l", "log file", Optional, ""},
			{"mode|m", "operating mode", Required, ""},
		},
	}

	os.Args = []string{"prog", "-dvDl/tmp/log.txt"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["logfile"].String != "/tmp/log.txt" {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional)")
	}

	os.Args = []string{"prog", "-dvDl", "/tmp/log.txt"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["logfile"].String != "/tmp/log.txt" {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional separated by space)")
	}

	os.Args = []string{"prog", "-dvDmdaemon"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["mode"].String != "daemon" {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required)")
	}

	os.Args = []string{"prog", "-dvDm", "daemon"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["mode"].String != "daemon" {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required separated by space)")
	}

}

func TestConcatenatedOptionsParsingWithIntValueOptionAtTheEnd(t *testing.T) {
	options := Options{
		"",
		Definitions{
			{"debug|d", "debug mode", Flag, true},
			{"verbose|v", "verbose mode", Flag, true},
			{"dryrun|D", "dry run only", Flag, true},
			{"port|p", "port", Optional, 3000},
			{"instances|i", "instances", Required, 1},
		},
	}

	os.Args = []string{"prog", "-dvDp3000"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["port"].Int != 3000 {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int)")
	}

	os.Args = []string{"prog", "-dvDp", "3000"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["port"].Int != 3000 {
		fmt.Printf("%#v", opts)
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int separated by space)")
	}

	os.Args = []string{"prog", "-dvDp", "3000"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["port"].Int != 3000 {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int separated by space)")
	}

	os.Args = []string{"prog", "-dvDi4"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["instances"].Int != 4 {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required)")
	}

	os.Args = []string{"prog", "-dvDi", "4"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["instances"].Int != 4 {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required int separated by space)")
	}

}

func TestConcatenatedOptionsParsingWithIntArrayValueOptionAtTheEnd(t *testing.T) {
	options := Options{
		"",
		Definitions{
			{"debug|d", "debug mode", Flag, true},
			{"verbose|v", "verbose mode", Flag, true},
			{"dryrun|D", "dry run only", Flag, true},
			{"ports|p", "ports", Optional, []int{3000, 3001, 3002}},
			{"timeouts|t", "timeouts", Required, []int{1, 2, 4, 10, 30}},
		},
	}

	os.Args = []string{"prog", "-dvDp5000,5001,5002"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || !equalIntArray(opts["ports"].IntArray, []int64{5000, 5001, 5002}) {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int array)")
	}

	os.Args = []string{"prog", "-dvDp", "5000,5001,5002"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || !equalIntArray(opts["ports"].IntArray, []int64{5000, 5001, 5002}) {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int array separated by space)")
	}

	os.Args = []string{"prog", "-dvDt10,20,30"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || !equalIntArray(opts["timeouts"].IntArray, []int64{10, 20, 30}) {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required int array)")
	}

	os.Args = []string{"prog", "-dvDt", "10,20,30"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || !equalIntArray(opts["timeouts"].IntArray, []int64{10, 20, 30}) {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required int array separated by space)")
	}

}

func TestEnvironmentValueParsing(t *testing.T) {
	options := Options{
		"",
		Definitions{
			{"debug|d|DEBUG", "debug mode", Flag, true},
			{"ports|p|PORTS", "ports", Required, []int{3000, 3001, 3002}},
			{"instances||INSTANCES", "instances", Optional, 4},
			{"keys||KEYS", "keys", Optional, []string{"foo,bar,baz"}},
			{"logfile||LOGFILE", "ports", Optional | ExampleIsDefault, "/var/log/foo.log"},
			{"hostname|h|HOSTNAME", "hostname", Optional | ExampleIsDefault | NoLongOpt, "/var/log/foo.log"},
		},
	}

	os.Args = []string{"prog"}
	os.Setenv("DEBUG", "1")
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true {
		t.Errorf("did not recognize option set via ENV variable (DEBUG=1)")
	}

	os.Args = []string{"prog"}
	os.Setenv("DEBUG", "TRUE")
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true {
		t.Errorf("did not recognize option set via ENV variable (DEBUG=TRUE)")
	}

	os.Args = []string{"prog"}
	os.Setenv("DEBUG", "true")
	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true {
		t.Errorf("did not recognize option set via ENV variable (DEBUG=true)")
	}

	os.Args = []string{"prog"}
	os.Setenv("PORTS", "4000,4001,4002")
	if opts, _, _, _ := options.ParseCommandLine(); !equalIntArray(opts["ports"].IntArray, []int64{4000, 4001, 4002}) {
		t.Errorf("did not recognize option set via ENV variable (PORTS=4000,4001,4002)")
	}

	os.Args = []string{"prog"}
	os.Setenv("KEYS", "faa,bor,boz")
	if opts, _, _, _ := options.ParseCommandLine(); !equalStringArray(opts["keys"].StrArray, []string{"faa", "bor", "boz"}) {
		t.Errorf("did not recognize option set via ENV variable (KEYS=faa,bor,boz)")
	}

	os.Args = []string{"prog"}
	os.Setenv("LOGFILE", "/tmp/logfile")
	if opts, _, _, _ := options.ParseCommandLine(); opts["logfile"].String != "/tmp/logfile" {
		t.Errorf("did not recognize option set via ENV variable (LOGFILE=/tmp/lofile)")
	}

	os.Args = []string{"prog"}
	os.Setenv("INSTANCES", "13")
	if opts, _, _, _ := options.ParseCommandLine(); opts["instances"].Int != 13 {
		t.Errorf("did not recognize option set via ENV variable (INSTANCES=13)")
	}

	os.Args = []string{"prog"}
	os.Setenv("LOGFILE", "  /tmp/logfile  ")
	if opts, _, _, _ := options.ParseCommandLine(); opts["logfile"].String != "/tmp/logfile" {
		t.Errorf("did not recognize option set via ENV variable with whitespace (LOGFILE=/tmp/lofile)")
	}

	os.Args = []string{"prog"}
	os.Setenv("INSTANCES", "    13   ")
	if opts, _, _, _ := options.ParseCommandLine(); opts["instances"].Int != 13 {
		t.Errorf("did not recognize option set via ENV variable with whitespace (INSTANCES=13)")
	}

	os.Args = []string{"prog"}
	os.Setenv("HOSTNAME", "eberhard")
	if opts, _, _, _ := options.ParseCommandLine(); opts["hostname"].String != "eberhard" {
		t.Errorf("did not recognize NoLongOpt option set via ENV variable")
	}
}

func TestDefaultValueParsing(t *testing.T) {
	options := Options{
		"",
		Definitions{
			{"debug|d|DEBUG", "debug mode", Optional | ExampleIsDefault, true},
			{"ports|p|PORTS", "ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
			{"secondaryports|s|SECONDARY_PORTS", "secondary ports", Optional | ExampleIsDefault, []int{5000, 5001, 5002}},
			{"instances||INSTANCES", "instances", Optional | ExampleIsDefault, 4},
			{"keys||KEYS", "keys", Optional | ExampleIsDefault, []string{"foo", "bar", "baz"}},
			{"logfile||LOGFILE", "logfile", Optional | ExampleIsDefault, "/var/log/foo.log"},
		},
	}
	os.Args = []string{"prog"}
	os.Setenv("INSTANCES", "")
	os.Setenv("KEYS", "")
	os.Setenv("LOGFILE", "")

	if opts, _, _, _ := options.ParseCommandLine(); opts["instances"].Int != 4 {
		t.Errorf("did not recognize default value (int instances)")
	}

	if opts, _, _, _ := options.ParseCommandLine(); !equalStringArray(opts["keys"].StrArray, []string{"foo", "bar", "baz"}) {
		t.Errorf("did not recognize default value (string array keys)")
	}

	if opts, _, _, _ := options.ParseCommandLine(); opts["logfile"].String != "/var/log/foo.log" {
		t.Errorf("did not recognize default value (string logfile)")
	}

	if opts, _, _, _ := options.ParseCommandLine(); opts["debug"].Bool != true {
		t.Errorf("did not recognize default value (boolean debug)")
	}

	if opts, _, _, _ := options.ParseCommandLine(); !equalIntArray(opts["ports"].IntArray, []int64{3000, 3001, 3002}) {
		t.Errorf("did not recognize default value (int64 array ports)")
	}

	if opts, _, _, _ := options.ParseCommandLine(); !equalIntArray(opts["secondaryports"].IntArray, []int64{5001, 5002, 5003}) {
		t.Errorf("did not recognize default value (int array ports)")
	}

}

func TestArgumentsParsing(t *testing.T) {
	options := Options{
		"",
		Definitions{
			{"debug|d|DEBUG", "debug mode", Flag, true},
			{"ports|p|PORTS", "ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
		},
	}

	os.Args = []string{"prog", "-d", "foobar", "barbaz"}
	if opts, arguments, _, _ := options.ParseCommandLine(); !equalStringArray(arguments, []string{"foobar", "barbaz"}) || opts["debug"].Bool != true {
		t.Errorf("did not recognize arguments (two at the end)")
	}

	os.Args = []string{"prog", "foobar", "-d", "barbaz"}
	if opts, arguments, _, _ := options.ParseCommandLine(); !equalStringArray(arguments, []string{"foobar", "barbaz"}) || opts["debug"].Bool != true {
		t.Errorf("did not recognize arguments separated by bool option")
	}

	os.Args = []string{"prog", "foobar", "barbaz", "-d"}
	if opts, arguments, _, _ := options.ParseCommandLine(); !equalStringArray(arguments, []string{"foobar", "barbaz"}) || opts["debug"].Bool != true {
		t.Errorf("did not recognize arguments before option")
	}

	os.Args = []string{"prog", "-d", "-p5000,6000", "foobar", "barbaz"}
	if opts, arguments, _, _ := options.ParseCommandLine(); !equalStringArray(arguments, []string{"foobar", "barbaz"}) ||
		opts["debug"].Bool != true ||
		!equalIntArray(opts["ports"].IntArray, []int64{5000, 6000}) {
		t.Errorf("parsing error of command line: '-d -p5000,6000 foobar barbaz'")
	}

	os.Args = []string{"prog", "-dp5000,6000", "foobar", "barbaz"}
	if opts, arguments, _, _ := options.ParseCommandLine(); !equalStringArray(arguments, []string{"foobar", "barbaz"}) ||
		opts["debug"].Bool != true ||
		!equalIntArray(opts["ports"].IntArray, []int64{5000, 6000}) {
		t.Errorf("parsing error of command line: '-dp5000,6000 foobar barbaz'")
	}

	os.Args = []string{"prog", "-d", "foobar", "-p5000,6000", "barbaz"}
	if opts, arguments, _, _ := options.ParseCommandLine(); !equalStringArray(arguments, []string{"foobar", "barbaz"}) ||
		opts["debug"].Bool != true ||
		!equalIntArray(opts["ports"].IntArray, []int64{5000, 6000}) {
		t.Errorf("parsing error of command line: '-d foobar -p5000,6000 barbaz'")
	}

	os.Args = []string{"prog", "-p5000,6000", "foobar", "-d", "barbaz"}
	if opts, arguments, _, _ := options.ParseCommandLine(); !equalStringArray(arguments, []string{"foobar", "barbaz"}) ||
		opts["debug"].Bool != true ||
		!equalIntArray(opts["ports"].IntArray, []int64{5000, 6000}) {
		t.Errorf("parsing error of command line: '-p5000,6000 foobar -d barbaz'")
	}

	os.Args = []string{"prog", "barbaz", "-d", "-p5000,6000", "foobar"}
	if opts, arguments, _, _ := options.ParseCommandLine(); !equalStringArray(arguments, []string{"barbaz", "foobar"}) ||
		opts["debug"].Bool != true ||
		!equalIntArray(opts["ports"].IntArray, []int64{5000, 6000}) {
		fmt.Printf("args: %#v\nopts: %#v\nports: %#v\n", arguments, opts["debug"].Bool, opts["ports"].IntArray)
		t.Errorf("parsing error of command line: 'barbaz -d -p5000,6000 foobar'")
	}

}

func TestPassThroughParsing(t *testing.T) {
	options := Options{
		"",
		Definitions{
			{"debug|d|DEBUG", "debug mode", Flag, true},
			{"ports|p|PORTS", "ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
			{"command args", "command args", Required | IsPassThrough, "command"},
		},
	}

	os.Args = []string{"prog"}
	if _, _, _, err := options.ParseCommandLine(); err != nil {
		t.Errorf("missing pass through rose error: %#v", err)
	}

	os.Args = []string{"prog", "foobar", "--", "barbaz"}
	expected := []string{"barbaz"}
	if _, _, passThrough, _ := options.ParseCommandLine(); !equalStringArray(passThrough, expected) {
		t.Errorf("simple pass through not recognized:\ngot:      |" + fmt.Sprintf("%#v", passThrough) + "|\nexpected: |" + fmt.Sprintf("%#v", expected) + "|\n")
	}

	os.Args = []string{"prog", "foobar", "--", "ls", "-lah", "/tmp"}
	expected = []string{"ls", "-lah", "/tmp"}
	if _, _, passThrough, _ := options.ParseCommandLine(); !equalStringArray(passThrough, expected) {
		t.Errorf("simple pass through not recognized:\ngot:      |" + fmt.Sprintf("%#v", passThrough) + "|\nexpected: |" + fmt.Sprintf("%#v", expected) + "|\n")
	}
}
