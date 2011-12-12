// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import "testing"
import "fmt"

func TestShortOptionsFlagsParsing(t *testing.T) {
	options := Options{
		{"debug|d", "debug mode", Flag, ""},
		{"verbose|v", "verbose mode", Flag, ""},
		{"dryrun|D", "dry run only", Flag, ""},
	}

	if opts, _, _, _ := options.Parse([]string{"-d"}, []string{}, "", 0); opts["debug"].Bool != true {
		t.Errorf("debug flag was not set")
	}

	if opts, _, _, _ := options.Parse([]string{"-d"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != false {
		t.Errorf("verbose flag was not set to false by default")
	}

	if opts, _, _, _ := options.Parse([]string{"-d", "-v"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true {
		t.Errorf("did not recognize all flags")
	}

}

func TestShortOptionRequiredParsing(t *testing.T) {
	options := Options{
		{"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", Required, ""},
	}

	if opts, _, _, _ := options.Parse([]string{"-m", "heartbeat"}, []string{}, "", 0); opts["method"].String != "heartbeat" {
		t.Errorf("method optioned wasn't parsed correctly expected 'heartbeat', was '" + opts["method"].String + "'")
	}

	if opts, _, _, _ := options.Parse([]string{"-mheartbeat"}, []string{}, "", 0); opts["method"].String != "heartbeat" {
		t.Errorf("method optioned wasn't parsed correctly expected 'heartbeat', was '" + opts["method"].String + "'")
	}

	if _, _, _, err := options.Parse([]string{"-m"}, []string{}, "", 0); err == nil || err.ErrorCode != MissingValue {
		t.Errorf("required option without value did not raise error")
	}

	if _, _, _, err := options.Parse([]string{"-m", "-x"}, []string{}, "", 0); err == nil || err.ErrorCode != MissingValue {
		t.Errorf("required option without value did not raise error")
	}

	if _, _, _, err := options.Parse([]string{""}, []string{}, "", 0); err == nil || err.ErrorCode != MissingOption {
		t.Errorf("required option wasn't set")
	}

}

func TestConcatenatedOptionsParsingSimple(t *testing.T) {
	options := Options{
		{"debug|d", "debug mode", Flag, true},
		{"verbose|v", "verbose mode", Flag, true},
		{"dryrun|D", "dry run only", Flag, true},
		{"logfile|l", "log file", Optional, ""},
		{"mode|m", "operating mode", Required, ""},
	}

	if opts, _, _, _ := options.Parse([]string{"-dv"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true {
		t.Errorf("did not recognize all flags when concatenation options (2 options)")
	}

	if opts, _, _, _ := options.Parse([]string{"-dvD"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true || opts["dryrun"].Bool != true {
		t.Errorf("did not recognize all flags when concatenation options (3 options)")
	}

	if _, _, _, err := options.Parse([]string{"-dvD"}, []string{}, "", 0); err == nil || err.ErrorCode != MissingOption {
		t.Errorf("did not recognize a missing required option in concatenation mode")
	}

	if _, _, _, err := options.Parse([]string{"-Dl"}, []string{}, "", 0); err == nil || err.ErrorCode != MissingValue {
		t.Errorf("did not realize that I missed a value")
	}

	if _, _, _, err := options.Parse([]string{"-Dl", "-d"}, []string{}, "", 0); err == nil || err.ErrorCode != MissingValue {
		t.Errorf("did not realize that I missed a value")
	}

}

func TestConcatenatedOptionsParsingWithStringValueOptionAtTheEnd(t *testing.T) {
	options := Options{
		{"debug|d", "debug mode", Flag, true},
		{"verbose|v", "verbose mode", Flag, true},
		{"dryrun|D", "dry run only", Flag, true},
		{"logfile|l", "log file", Optional, ""},
		{"mode|m", "operating mode", Required, ""},
	}
	if opts, _, _, _ := options.Parse([]string{"-dvDl/tmp/log.txt"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["logfile"].String != "/tmp/log.txt" {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional)")
	}

	if opts, _, _, _ := options.Parse([]string{"-dvDl", "/tmp/log.txt"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["logfile"].String != "/tmp/log.txt" {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional separated by space)")
	}

	if opts, _, _, _ := options.Parse([]string{"-dvDmdaemon"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["mode"].String != "daemon" {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required)")
	}

	if opts, _, _, _ := options.Parse([]string{"-dvDm", "daemon"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["mode"].String != "daemon" {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required separated by space)")
	}

}

func TestConcatenatedOptionsParsingWithIntValueOptionAtTheEnd(t *testing.T) {
	options := Options{
		{"debug|d", "debug mode", Flag, true},
		{"verbose|v", "verbose mode", Flag, true},
		{"dryrun|D", "dry run only", Flag, true},
		{"port|p", "port", Optional, 3000},
		{"instances|i", "instances", Required, 1},
	}
	if opts, _, _, _ := options.Parse([]string{"-dvDp3000"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["port"].Int != 3000 {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int)")
	}

	if opts, _, _, _ := options.Parse([]string{"-dvDp", "3000"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["port"].Int != 3000 {
		fmt.Printf("%#v", opts)
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int separated by space)")
	}

	if opts, _, _, _ := options.Parse([]string{"-dvDp", "3000"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["port"].Int != 3000 {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int separated by space)")
	}

	if opts, _, _, _ := options.Parse([]string{"-dvDi4"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["instances"].Int != 4 {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required)")
	}

	if opts, _, _, _ := options.Parse([]string{"-dvDi", "4"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || opts["instances"].Int != 4 {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required int separated by space)")
	}

}

func TestConcatenatedOptionsParsingWithIntArrayValueOptionAtTheEnd(t *testing.T) {
	options := Options{
		{"debug|d", "debug mode", Flag, true},
		{"verbose|v", "verbose mode", Flag, true},
		{"dryrun|D", "dry run only", Flag, true},
		{"ports|p", "ports", Optional, []int{3000, 3001, 3002}},
		{"timeouts|t", "timeouts", Required, []int{1, 2, 4, 10, 30}},
	}
	if opts, _, _, _ := options.Parse([]string{"-dvDp5000,5001,5002"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || !equalIntArray(opts["ports"].IntArray, []int64{5000, 5001, 5002}) {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int array)")
	}

	if opts, _, _, _ := options.Parse([]string{"-dvDp", "5000,5001,5002"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || !equalIntArray(opts["ports"].IntArray, []int64{5000, 5001, 5002}) {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int array separated by space)")
	}

	if opts, _, _, _ := options.Parse([]string{"-dvDt10,20,30"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || !equalIntArray(opts["timeouts"].IntArray, []int64{10, 20, 30}) {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required int array)")
	}

	if opts, _, _, _ := options.Parse([]string{"-dvDt", "10,20,30"}, []string{}, "", 0); opts["debug"].Bool != true || opts["verbose"].Bool != true ||
		opts["dryrun"].Bool != true || !equalIntArray(opts["timeouts"].IntArray, []int64{10, 20, 30}) {
		t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required int array separated by space)")
	}

}

func TestEnvironmentValueParsing(t *testing.T) {
	options := Options{
		{"debug|d|DEBUG", "debug mode", Flag, true},
		{"ports|p|PORTS", "ports", Required, []int{3000, 3001, 3002}},
		{"instances||INSTANCES", "instances", Optional, 4},
		{"keys||KEYS", "keys", Optional, []string{"foo,bar,baz"}},
		{"logfile||LOGFILE", "ports", Optional | ExampleIsDefault, "/var/log/foo.log"},
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{"DEBUG=1"}, "", 0); opts["debug"].Bool != true {
		t.Errorf("did not recognize option set via ENV variable (DEBUG=1)")
	}
	if opts, _, _, _ := options.Parse([]string{}, []string{"DEBUG=TRUE"}, "", 0); opts["debug"].Bool != true {
		t.Errorf("did not recognize option set via ENV variable (DEBUG=TRUE)")
	}
	if opts, _, _, _ := options.Parse([]string{}, []string{"DEBUG=true"}, "", 0); opts["debug"].Bool != true {
		t.Errorf("did not recognize option set via ENV variable (DEBUG=true)")
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{"PORTS=4000,4001,4002"}, "", 0); !equalIntArray(opts["ports"].IntArray, []int64{4000, 4001, 4002}) {
		t.Errorf("did not recognize option set via ENV variable (PORTS=4000,4001,4002)")
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{"KEYS=faa,bor,boz"}, "", 0); !equalStringArray(opts["keys"].StrArray, []string{"faa", "bor", "boz"}) {
		t.Errorf("did not recognize option set via ENV variable (KEYS=faa,bor,boz)")
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{"LOGFILE=/tmp/logfile"}, "", 0); opts["logfile"].String != "/tmp/logfile" {
		t.Errorf("did not recognize option set via ENV variable (LOGFILE=/tmp/lofile)")
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{"INSTANCES=13"}, "", 0); opts["instances"].Int != 13 {
		t.Errorf("did not recognize option set via ENV variable (INSTANCES=13)")
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{"   LOGFILE   =  /tmp/logfile  "}, "", 0); opts["logfile"].String != "/tmp/logfile" {
		t.Errorf("did not recognize option set via ENV variable with whitespace (LOGFILE=/tmp/lofile)")
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{"INSTANCES   =    13   "}, "", 0); opts["instances"].Int != 13 {
		t.Errorf("did not recognize option set via ENV variable with whitespace (INSTANCES=13)")
	}
}

func TestDefaultValueParsing(t *testing.T) {
	options := Options{
		{"debug|d|DEBUG", "debug mode", Optional | ExampleIsDefault, true},
		{"ports|p|PORTS", "ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
		{"secondaryports|s|SECONDARY_PORTS", "secondary ports", Optional | ExampleIsDefault, []int{5000, 5001, 5002}},
		{"instances||INSTANCES", "instances", Optional | ExampleIsDefault, 4},
		{"keys||KEYS", "keys", Optional | ExampleIsDefault, []string{"foo", "bar", "baz"}},
		{"logfile||LOGFILE", "logfile", Optional | ExampleIsDefault, "/var/log/foo.log"},
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{}, "", 0); opts["debug"].Bool != true {
		t.Errorf("did not recognize default value (boolean debug)")
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{}, "", 0); opts["instances"].Int != 4 {
		t.Errorf("did not recognize default value (int instances)")
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{}, "", 0); opts["logfile"].String != "/var/log/foo.log" {
		t.Errorf("did not recognize default value (string logfile)")
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{}, "", 0); !equalIntArray(opts["ports"].IntArray, []int64{3000, 3001, 3002}) {
		t.Errorf("did not recognize default value (int64 array ports)")
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{}, "", 0); !equalIntArray(opts["secondaryports"].IntArray, []int64{5001, 5002, 5003}) {
		t.Errorf("did not recognize default value (int array ports)")
	}

	if opts, _, _, _ := options.Parse([]string{}, []string{}, "", 0); !equalStringArray(opts["keys"].StrArray, []string{"foo", "bar", "baz"}) {
		t.Errorf("did not recognize default value (string array keys)")
	}

}

func TestArgumentsParsing(t *testing.T) {
	options := Options{
		{"debug|d|DEBUG", "debug mode", Flag, true},
		{"ports|p|PORTS", "ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
	}

	if opts, arguments, _, _ := options.Parse([]string{"-d", "foobar", "barbaz"}, []string{}, "", 0); !equalStringArray(arguments, []string{"foobar", "barbaz"}) || opts["debug"].Bool != true {
		t.Errorf("did not recognize arguments (two at the end)")
	}

	if opts, arguments, _, _ := options.Parse([]string{"foobar", "-d", "barbaz"}, []string{}, "", 0); !equalStringArray(arguments, []string{"foobar", "barbaz"}) || opts["debug"].Bool != true {
		t.Errorf("did not recognize arguments separated by bool option")
	}

	if opts, arguments, _, _ := options.Parse([]string{"foobar", "barbaz", "-d"}, []string{}, "", 0); !equalStringArray(arguments, []string{"foobar", "barbaz"}) || opts["debug"].Bool != true {
		t.Errorf("did not recognize arguments before option")
	}

	if opts, arguments, _, _ := options.Parse([]string{"-d", "-p5000,6000", "foobar", "barbaz"}, []string{}, "", 0); !equalStringArray(arguments, []string{"foobar", "barbaz"}) ||
		opts["debug"].Bool != true ||
		!equalIntArray(opts["ports"].IntArray, []int64{5000, 6000}) {
		t.Errorf("parsing error of command line: '-d -p5000,6000 foobar barbaz'")
	}

	if opts, arguments, _, _ := options.Parse([]string{"-dp5000,6000", "foobar", "barbaz"}, []string{}, "", 0); !equalStringArray(arguments, []string{"foobar", "barbaz"}) ||
		opts["debug"].Bool != true ||
		!equalIntArray(opts["ports"].IntArray, []int64{5000, 6000}) {
		t.Errorf("parsing error of command line: '-dp5000,6000 foobar barbaz'")
	}

	if opts, arguments, _, _ := options.Parse([]string{"-d", "foobar", "-p5000,6000", "barbaz"}, []string{}, "", 0); !equalStringArray(arguments, []string{"foobar", "barbaz"}) ||
		opts["debug"].Bool != true ||
		!equalIntArray(opts["ports"].IntArray, []int64{5000, 6000}) {
		t.Errorf("parsing error of command line: '-d foobar -p5000,6000 barbaz'")
	}

	if opts, arguments, _, _ := options.Parse([]string{"-p5000,6000", "foobar", "-d", "barbaz"}, []string{}, "", 0); !equalStringArray(arguments, []string{"foobar", "barbaz"}) ||
		opts["debug"].Bool != true ||
		!equalIntArray(opts["ports"].IntArray, []int64{5000, 6000}) {
		t.Errorf("parsing error of command line: '-p5000,6000 foobar -d barbaz'")
	}

	if opts, arguments, _, _ := options.Parse([]string{"barbaz", "-d", "-p5000,6000", "foobar"}, []string{}, "", 0); !equalStringArray(arguments, []string{"barbaz", "foobar"}) ||
		opts["debug"].Bool != true ||
		!equalIntArray(opts["ports"].IntArray, []int64{5000, 6000}) {
		fmt.Printf("args: %#v\nopts: %#v\nports: %#v\n", arguments, opts["debug"].Bool, opts["ports"].IntArray)
		t.Errorf("parsing error of command line: 'barbaz -d -p5000,6000 foobar'")
	}

}

func TestPassThroughParsing(t *testing.T) {
	options := Options{
		{"debug|d|DEBUG", "debug mode", Flag, true},
		{"ports|p|PORTS", "ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
	}

	if _, _, passThrough, _ := options.Parse([]string{"foobar", "--", "barbaz"}, []string{}, "", 0); !equalStringArray(passThrough, []string{"foobar"}) {
		t.Errorf("simple pass through not recognized")
	}

	if _, _, passThrough, _ := options.Parse([]string{"foobar", "--", "ls", "-lah", "/tmp"}, []string{}, "", 0); !equalStringArray(passThrough, []string{"ls", "-lah", "/tmp"}) {
		t.Errorf("pass through not recognized")
	}
}
