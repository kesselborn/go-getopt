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
  }

  expected := `

    Usage: testprogram -d [-p PORTS] FILES [-s SECONDARYPORTS] --instances=INSTANCES --lock [--logfile=LOGFILE] [DIRECTORIES]

this is not a program

Options:
    -d, --debug                           debug mode; setable via $DEBUG
    -p, --ports=PORTS                     ports (default: 3000,3001,3002); setable via $PORTS
    -s, --secondaryports=SECONDARYPORTS   secondary ports (default: 5000,5001,5002)
        --instances=INSTANCES             instances (e.g. 4)
        --lock                            create lock file; setable via $LOCK
        --logfile=LOGFILE                 logfile (default: /var/log/foo.log); setable via $LOGFILE

Arguments:
    FILES                                 files that should be read in
    DIRECTORIES                           directories

`

  if got := options.Help("testprogram", "this is not a program"); got != expected {
    t.Errorf("Usage output not as expected:\ngot:      |" + strings.Replace(got, " ", "_", -1) + "|\nexpected: |" + strings.Replace(expected, " ", "_", -1) + "|\n")
  }

}
