package getopt

import(
  "testing"
)

func TestUsageOutput(t *testing.T) {
  if got, expected := (Option{"method", "...", IsArg | Optional, ""}.Usage()),
                      "[METHOD]"
     got != expected {
       t.Errorf("Error creating usage text (argument):\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "...", IsArg, ""}.Usage()),
                      "METHOD"
     got != expected {
       t.Errorf("Error creating usage text (argument):\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "...", Required, ""}.Usage()),
                      "-m METHOD"
     got != expected {
       t.Errorf("Error creating usage text (Required):\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "...", Required, ""}.Usage()),
                      "--method=METHOD"
     got != expected {
       t.Errorf("Error creating usage text (Required, no short):\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "...", Optional, ""}.Usage()),
                      "[-m METHOD]"
     got != expected {
       t.Errorf("Error creating usage text (Optional):\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "...", Optional, ""}.Usage()),
                      "[--method=METHOD]"
     got != expected {
       t.Errorf("Error creating usage text (Optional, no short):\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "...", Flag | Optional, ""}.Usage()),
                      "[-m]"
     got != expected {
       t.Errorf("Error creating usage text (optional flag):\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "...", Flag | Optional, ""}.Usage()),
                      "[--method]"
     got != expected {
       t.Errorf("Error creating usage text (optional flag, no short):\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "...", Flag, ""}.Usage()),
                      "-m"
     got != expected {
       t.Errorf("Error creating usage text (flag):\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "...", Flag, ""}.Usage()),
                      "--method"
     got != expected {
       t.Errorf("Error creating usage text (flag, no short):\ngot:      " + got + "\nexpected: " + expected )
  }


}

func TestBasicOutput(t *testing.T) {
  if got, expected := (Option{"method|m", "method: one of either 'heartbeat' or 'nagios'", Required | NoLongOpt, ""}.String(20)),
                      "        -m METHOD                  method: one of either 'heartbeat' or 'nagios'";
     got != expected {
       t.Errorf("Error stringifying option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "method: one of either 'heartbeat' or 'nagios'", Required, ""}.String(20)),
                      "            --method=METHOD        method: one of either 'heartbeat' or 'nagios'";
     got != expected {
       t.Errorf("Error stringifying option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "method: one of either 'heartbeat' or 'nagios'", Required, ""}.String(20)),
                      "        -m, --method=METHOD        method: one of either 'heartbeat' or 'nagios'";
     got != expected {
       t.Errorf("Error stringifying option:\ngot:      " + got + "\nexpected: " + expected )
  }
}

func TestBasicOutputWithExampleStringValue(t *testing.T) {
  if got, expected := (Option{"method|m", "method", Optional | NoLongOpt, "heartbeat"}.String(20)),
                      "        -m METHOD                  method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "method", Optional, "heartbeat"}.String(20)),
                      "            --method=METHOD        method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "method", Optional, "heartbeat"}.String(20)),
                      "        -m, --method=METHOD        method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }


  if got, expected := (Option{"method|m", "method", Required | NoLongOpt, "heartbeat"}.String(20)),
                      "        -m METHOD                  method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying required option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "method", Required, "heartbeat"}.String(20)),
                      "            --method=METHOD        method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying required option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "method", Required, "heartbeat"}.String(20)),
                      "        -m, --method=METHOD        method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying required option:\ngot:      " + got + "\nexpected: " + expected )
  }
}

func TestBasicOutputWithDefaultStringValue(t *testing.T) {
  if got, expected := (Option{"method|m", "method", Optional | ExampleIsDefault | NoLongOpt, "heartbeat"}.String(20)),
                      "        -m METHOD                  method (default: heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "method", Optional | ExampleIsDefault, "heartbeat"}.String(20)),
                      "            --method=METHOD        method (default: heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "method", Optional | ExampleIsDefault, "heartbeat"}.String(20)),
                      "        -m, --method=METHOD        method (default: heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }
}

func TestOutputWithEnvVar(t *testing.T) {
  if got, expected := (Option{"method|m|METHOD", "method", Optional | ExampleIsDefault | NoLongOpt, "heartbeat"}.String(20)),
                      "        -m METHOD                  method (default: heartbeat); setable via $METHOD";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }
}

func TestOutputArgument(t *testing.T) {
  if got, expected := (Option{"logfile", "", IsArg, ""}.String(20)),
                      "";
     got != expected {
       t.Errorf("Error stringifying argument:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"logfile", "dump log into this file", IsArg, ""}.String(20)),
                      "        LOGFILE                    dump log into this file";
     got != expected {
       t.Errorf("Error stringifying argument:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"logfile", "dump log into this file", IsArg, "/tmp/foo"}.String(20)),
                      "        LOGFILE                    dump log into this file (e.g. /tmp/foo)";
     got != expected {
       t.Errorf("Error stringifying argument:\ngot:      " + got + "\nexpected: " + expected )
  }
}
