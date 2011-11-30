package getopt
import "testing"

func TestIsFlagOption(t *testing.T) {
  options := Options{
    {"verbose"            , "show verbose output"                          , Flag, ""},
    {"logfile|l"          , "log to file <logfile>"                        , Flag | Optional | NoLongOpt, ""},
    {"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", Required, ""},
  }

  if options.IsFlag("verbose") != true  { t.Errorf("flag 'verbose' not recognized") }

  if options.IsFlag("logfile") != false { t.Errorf("flag 'logfile' not recognized") } // NoLongOpt
  if options.IsFlag("l")       != true  { t.Errorf("flag 'l' not recognized") }

  if options.IsFlag("method")  != false { t.Errorf("non-flag 'method' incorrectly recognized as flag") }
  if options.IsFlag("m")       != false { t.Errorf("non-flag 'm' incorrectly recognized as flag") }
}

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

//func TestRequiredOptions(t *testing.T) {
//  option1 := Option{"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", Required, ""}
//  option2 := Option{"logfile|l"          , "log to file <logfile>"                        , Optional | NoLongOpt, ""}
//  option3 := Option{"verbose"            , "show verbose output"                          , Required, ""}
//
//  options := Options{ option1, option2, option3 }
//
//  if options.RequiredOptions() != []string{"method", "verbose"} { t.Errorf("did not find correct required options") }
//
//
//}


