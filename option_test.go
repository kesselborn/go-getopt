package getopt

import(
  "testing"
)

func TestOptionMethods(t *testing.T) {
  var option Option

  option = Option{"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", Required, ""}
  if option.LongOpt()  != "method"     { t.Errorf("'method|m|MON_METHOD': longopt parsing failed (expected 'method', got '" + option.LongOpt() + "')") }
  if option.ShortOpt() != "m"          { t.Errorf("'method|m|MON_METHOD': shortopt parsing failed (expected 'm', got '" + option.ShortOpt() + "')") }
  if option.EnvVar()   != "MON_METHOD" { t.Errorf("'method|m|MON_METHOD': env var parsing failed") }

  option = Option{"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", NoLongOpt, ""}
  if option.LongOpt()  != "" {
    t.Errorf("'method|m|MON_METHOD' with option NoLongOpt: longopt parsing failed (expected '', got '" + option.LongOpt() + "')")
  }

  option = Option{"method||MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", Required, ""}
  if option.LongOpt()  != "method"     { t.Errorf("'method||MON_METHOD': longopt parsing failed (expected 'method', got '" + option.LongOpt() + "')") }
  if option.ShortOpt() != ""           { t.Errorf("'method||MON_METHOD': shortopt parsing failed (expected '', got '" + option.ShortOpt() + "')") }
  if option.EnvVar()   != "MON_METHOD" { t.Errorf("'method||MON_METHOD': env var parsing failed") }

  option = Option{"method", "method: one of either 'heartbeat' or 'nagios'", Required, ""}
  if option.LongOpt()  != "method"     { t.Errorf("'method': longopt parsing failed (expected 'method', got '" + option.LongOpt() + "')") }
  if option.ShortOpt() != ""           { t.Errorf("'method': shortopt parsing failed (expected '', got '" + option.ShortOpt() + "')") }
  if option.EnvVar()   != ""           { t.Errorf("'method': env var parsing failed") }
}

