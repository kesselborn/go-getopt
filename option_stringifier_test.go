package getopt

import(
  "testing"
)

func TestBasicOutput(t *testing.T) {
  if got, expected := (Option{"method|m", "method: one of either 'heartbeat' or 'nagios'", Required | NoLongOpt, ""}.String(20)),
                      "\t-m METHOD                  method: one of either 'heartbeat' or 'nagios'";
     got != expected {
       t.Errorf("Error stringifying option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "method: one of either 'heartbeat' or 'nagios'", Required, ""}.String(20)),
                      "\t    --method=METHOD        method: one of either 'heartbeat' or 'nagios'";
     got != expected {
       t.Errorf("Error stringifying option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "method: one of either 'heartbeat' or 'nagios'", Required, ""}.String(20)),
                      "\t-m, --method=METHOD        method: one of either 'heartbeat' or 'nagios'";
     got != expected {
       t.Errorf("Error stringifying option:\ngot:      " + got + "\nexpected: " + expected )
  }
}

func TestBasicOutputWithExampleStringValue(t *testing.T) {
  if got, expected := (Option{"method|m", "method", Optional | NoLongOpt, "heartbeat"}.String(20)),
                      "\t-m METHOD                  method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "method", Optional, "heartbeat"}.String(20)),
                      "\t    --method=METHOD        method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "method", Optional, "heartbeat"}.String(20)),
                      "\t-m, --method=METHOD        method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }


  if got, expected := (Option{"method|m", "method", Required | NoLongOpt, "heartbeat"}.String(20)),
                      "\t-m METHOD                  method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying required option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "method", Required, "heartbeat"}.String(20)),
                      "\t    --method=METHOD        method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying required option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "method", Required, "heartbeat"}.String(20)),
                      "\t-m, --method=METHOD        method (e.g. heartbeat)";
     got != expected {
       t.Errorf("Error stringifying required option:\ngot:      " + got + "\nexpected: " + expected )
  }
}

func TestBasicOutputWithDefaultStringValue(t *testing.T) {
  if got, expected := (Option{"method|m", "method", Optional | ExampleIsDefault | NoLongOpt, "heartbeat"}.String(20)),
                      "\t-m METHOD                  method (default: heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method", "method", Optional | ExampleIsDefault, "heartbeat"}.String(20)),
                      "\t    --method=METHOD        method (default: heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }

  if got, expected := (Option{"method|m", "method", Optional | ExampleIsDefault, "heartbeat"}.String(20)),
                      "\t-m, --method=METHOD        method (default: heartbeat)";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }
}

func TestOutputWithEnvVar(t *testing.T) {
  if got, expected := (Option{"method|m|METHOD", "method", Optional | ExampleIsDefault | NoLongOpt, "heartbeat"}.String(20)),
                      "\t-m METHOD                  method (default: heartbeat); setable via $METHOD";
     got != expected {
       t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected )
  }
}
