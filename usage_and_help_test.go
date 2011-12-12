// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import(
  "testing"
  "strings"
)

func TestUsage(t *testing.T) {
  options := Options{
    {"debug|d|DEBUG", "debug mode", Flag, true},
    {"ports|p|PORTS", "ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
    {"files", "files that should be read in", IsArg, nil},
    {"secondaryports|s|SECONDARY_PORTS", "secondary ports", Optional | ExampleIsDefault, []int{5000,5001,5002}},
    {"instances||INSTANCES", "instances", Required, 4},
    {"lock||LOCK", "create lock file", Flag, false},
    {"logfile||LOGFILE", "logfile", Optional | ExampleIsDefault, "/var/log/foo.log"},
    {"directories", "directories", IsArg | Optional, nil},
  }

  expected := `

    Usage: testprogram -d [-p PORTS] FILES [-s SECONDARYPORTS] --instances=INSTANCES --lock [--logfile=LOGFILE] [DIRECTORIES]

`

  if got := options.Usage("testprogram"); got != expected {
    t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
  }

}

func TestHelp(t *testing.T) {
  options := Options{
    {"debug|d|DEBUG", "debug mode", Flag, true},
    {"ports|p|PORTS", "ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
    {"files", "files that should be read in", IsArg, nil},
    {"secondaryports|s", "secondary ports", Optional | ExampleIsDefault, []int{5000,5001,5002}},
    {"instances", "instances", Required, 4},
    {"lock||LOCK", "create lock file", Flag, false},
    {"logfile||LOGFILE", "logfile", Optional | ExampleIsDefault, "/var/log/foo.log"},
    {"directories", "directories", IsArg | Optional, nil},
    {"pass through args", "arguments for subcommand", IsPassThrough, nil},
  }

  expected := `

    Usage: testprogram -d [-p PORTS] FILES [-s SECONDARYPORTS] --instances=INSTANCES --lock [--logfile=LOGFILE] [DIRECTORIES] -- PASS THROUGH ARGS

this is not a program

Options:
    -d, --debug                           debug mode; setable via $DEBUG
    -p, --ports=PORTS                     ports (default: 3000,3001,3002); setable via $PORTS
    -s, --secondaryports=SECONDARYPORTS   secondary ports (default: 5000,5001,5002)
        --instances=INSTANCES             instances (e.g. 4)
        --lock                            create lock file; setable via $LOCK
        --logfile=LOGFILE                 logfile (default: /var/log/foo.log); setable via $LOGFILE
    -h, --help                            usage (-h) / detailed help text (--help)

Arguments:
    FILES                                 files that should be read in
    DIRECTORIES                           directories

Pass through arguments:
    PASS THROUGH ARGS                     arguments for subcommand

`

  if got := options.Help("testprogram", "this is not a program"); got != expected {
    t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
  }

}

func TestHelpNoOptions(t *testing.T) {
  options := Options{
    {"files", "files that should be read in", IsArg, nil},
    {"directories", "directories", IsArg | Optional, nil},
  }


  expected := `

    Usage: testprogram FILES [DIRECTORIES]

this is not a program

Arguments:
    FILES                           files that should be read in
    DIRECTORIES                     directories

`

  if got := options.Help("testprogram", "this is not a program"); got != expected {
    t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
  }

}

func TestHelpNoArgs(t *testing.T) {
  options := Options{
    {"debug|d|DEBUG", "debug mode", Flag, true},
    {"ports|p|PORTS", "ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
    {"secondaryports|s", "secondary ports", Optional | ExampleIsDefault, []int{5000,5001,5002}},
    {"instances", "instances", Required, 4},
    {"lock||LOCK", "create lock file", Flag, false},
    {"logfile||LOGFILE", "logfile", Optional | ExampleIsDefault, "/var/log/foo.log"},
  }

  expected := `

    Usage: testprogram -d [-p PORTS] [-s SECONDARYPORTS] --instances=INSTANCES --lock [--logfile=LOGFILE]

this is not a program

Options:
    -d, --debug                           debug mode; setable via $DEBUG
    -p, --ports=PORTS                     ports (default: 3000,3001,3002); setable via $PORTS
    -s, --secondaryports=SECONDARYPORTS   secondary ports (default: 5000,5001,5002)
        --instances=INSTANCES             instances (e.g. 4)
        --lock                            create lock file; setable via $LOCK
        --logfile=LOGFILE                 logfile (default: /var/log/foo.log); setable via $LOGFILE
    -h, --help                            usage (-h) / detailed help text (--help)

`

  if got := options.Help("testprogram", "this is not a program"); got != expected {
    t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
  }

}

func TestUsageAndHelpOption(t *testing.T) {
  options := Options{
    {"debug|d|DEBUG", "debug mode", Flag, true},
    {"ports|p|PORTS", "ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
  }

  expectedUsage := `

    Usage: 6.out -d [-p PORTS]

`

  if _, _, _, err := options.parse([]string{"barbaz", "-d", "-h", "-p5000,6000", "foobar"}, []string{}, "", 0);
    err == nil || err.errorCode != UsageOrHelp || err.message != expectedUsage  {
    t.Errorf("Usage text wasn't shown with single '-h':\ngot:      |" + strings.Replace(err.message, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedUsage, " ", "_", -1) + "|\n")
  }

  expectedHelp := `

    Usage: 6.out -d [-p PORTS]

Options:
    -d, --debug         debug mode; setable via $DEBUG
    -p, --ports=PORTS   ports (default: 3000,3001,3002); setable via $PORTS
    -h, --help          usage (-h) / detailed help text (--help)

`

  if _, _, _, err := options.parse([]string{"barbaz", "-d", "--help", "-p5000,6000", "foobar"}, []string{}, "", 0);
    err == nil || err.errorCode != UsageOrHelp || err.message != expectedHelp  {
    t.Errorf("Usage text wasn't shown with single '-h':\ngot:      |" + strings.Replace(err.message, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
  }

}

func TestUsageAndHelpOptionWithOwnIdentifiers(t *testing.T) {
  options := Options{
    {"chelp|c", "show usage / help", Usage | Help, nil},
    {"debug|d|DEBUG", "debug mode", Flag, true},
    {"ports|p|PORTS", "ports", Optional | ExampleIsDefault, []int64{3000, 3001, 3002}},
  }

  expectedUsage := `

    Usage: 6.out [-c] -d [-p PORTS]

`

  if _, _, _, err := options.parse([]string{"barbaz", "-d", "-c", "-p5000,6000", "foobar"}, []string{}, "", 0);
    err == nil || err.errorCode != UsageOrHelp || err.message != expectedUsage  {
    t.Errorf("Usage text wasn't shown with single '-h':\ngot:      |" + strings.Replace(err.message, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedUsage, " ", "_", -1) + "|\n")
  }

  expectedHelp := `

    Usage: 6.out [-c] -d [-p PORTS]

Options:
    -d, --debug         debug mode; setable via $DEBUG
    -p, --ports=PORTS   ports (default: 3000,3001,3002); setable via $PORTS
    -c, --chelp         show usage / help

`

  if _, _, _, err := options.parse([]string{"barbaz", "-d", "--chelp", "-p5000,6000", "foobar"}, []string{}, "", 0);
    err == nil || err.errorCode != UsageOrHelp || err.message != expectedHelp  {
    t.Errorf("Usage text wasn't shown with single '-h':\ngot:      |" + strings.Replace(err.message, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expectedHelp, " ", "_", -1) + "|\n")
  }

}

