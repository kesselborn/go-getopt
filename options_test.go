package getopt
import "testing"

func TestIsOptionalOption(t *testing.T) {
  options := Options{
    {"verbose"            , "show verbose output"                          , Optional, ""},
    {"logfile|l"          , "log to file <logfile>"                        , Optional | NoLongOpt, ""},
    {"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", 0, ""},
  }

  if options.IsOptional("verbose") != true  { t.Errorf("optional flag for 'verbose' not recognized") }

  if options.IsOptional("logfile") != false { t.Errorf("optional flag for 'logfile' not recognized") } // NoLongOpt
  if options.IsOptional("l")       != true  { t.Errorf("optional flag for 'l' not recognized") }

  if options.IsOptional("method")  != false { t.Errorf("non-optional flag for 'method' incorrectly recognized as optional flag for") }
  if options.IsOptional("m")       != false { t.Errorf("non-optional flag for 'm' incorrectly recognized as flag") }
}

func TestIsRequiredOption(t *testing.T) {
  options := Options{
    {"verbose"            , "show verbose output"                          , Required, ""},
    {"logfile|l"          , "log to file <logfile>"                        , Required | NoLongOpt, ""},
    {"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", 0, ""},
  }

  if options.IsRequired("verbose") != true  { t.Errorf("required flag for 'verbose' not recognized") }

  if options.IsRequired("logfile") != false { t.Errorf("required flag for 'logfile' not recognized") } // NoLongOpt
  if options.IsRequired("l")       != true  { t.Errorf("required flag for 'l' not recognized") }

  if options.IsRequired("method")  != false { t.Errorf("non-required flag for 'method' incorrectly recognized as required flag for") }
  if options.IsRequired("m")       != false { t.Errorf("non-required flag for 'm' incorrectly recognized as flag") }
}

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

