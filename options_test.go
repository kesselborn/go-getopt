package getopt
import "testing"



func TestFindOption(t *testing.T) {
  option1 := Option{"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", Required, ""}
  option2 := Option{"logfile|l"          , "log to file <logfile>"                        , Optional | NoLongOpt, ""}
  option3 := Option{"verbose"            , "show verbose output"                          , Flag, ""}

  options := Options{ option1, option2, option3 }

  if val, _   := options.FindOption("method"); val.neq(option1) { t.Errorf("couldn't find option 'method'") }
  if _, found := options.FindOption("method"); found != true    { t.Errorf("couldn't find option 'method'") }

  if val, _   := options.FindOption("m"); val.neq(option1)      { t.Errorf("couldn't find option 'm'") }
  if _, found := options.FindOption("m"); found != true         { t.Errorf("couldn't find option 'm'") }

  if _, found := options.FindOption("logfile"); found == true   { t.Errorf("shouldn't find an option 'logfile'") }
}

