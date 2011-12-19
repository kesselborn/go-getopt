// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
	"testing"
)

func TestUsageOutput(t *testing.T) {
	if got, expected := (Option{"method", "...", IsArg | Optional, ""}.Usage()),
		"[<method>]"; got != expected {
		t.Errorf("Error creating usage text (argument):\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method", "...", IsArg, ""}.Usage()),
		"<method>"; got != expected {
		t.Errorf("Error creating usage text (argument):\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method|m", "...", Required, ""}.Usage()),
		"-m <method>"; got != expected {
		t.Errorf("Error creating usage text (Required):\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method", "...", Required, ""}.Usage()),
		"--method=<method>"; got != expected {
		t.Errorf("Error creating usage text (Required, no short):\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method|m", "...", Optional, ""}.Usage()),
		"[-m <method>]"; got != expected {
		t.Errorf("Error creating usage text (Optional):\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method", "...", Optional, ""}.Usage()),
		"[--method=<method>]"; got != expected {
		t.Errorf("Error creating usage text (Optional, no short):\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method|m", "...", Flag | Optional, ""}.Usage()),
		"[-m]"; got != expected {
		t.Errorf("Error creating usage text (optional flag):\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method", "...", Flag | Optional, ""}.Usage()),
		"[--method]"; got != expected {
		t.Errorf("Error creating usage text (optional flag, no short):\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method|m", "...", Flag, ""}.Usage()),
		"-m"; got != expected {
		t.Errorf("Error creating usage text (flag):\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method", "...", Flag, ""}.Usage()),
		"--method"; got != expected {
		t.Errorf("Error creating usage text (flag, no short):\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"configfile|c", "method", IsConfigFile | ExampleIsDefault, "/etc/cnf.conf"}.Usage()),
		"-c <configfile>"; got != expected {
		t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected)
	}
}

func TestBasicOutput(t *testing.T) {
	if got, expected := (Option{"method|m", "method: one of either 'heartbeat' or 'nagios'", Required | NoLongOpt, ""}.HelpText(20)),
		"    -m <method>                method: one of either 'heartbeat' or 'nagios'"; got != expected {
		t.Errorf("Error stringifying option:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method", "method: one of either 'heartbeat' or 'nagios'", Required, ""}.HelpText(20)),
		"        --method=<method>      method: one of either 'heartbeat' or 'nagios'"; got != expected {
		t.Errorf("Error stringifying option:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method|m", "method: one of either 'heartbeat' or 'nagios'", Required, ""}.HelpText(20)),
		"    -m, --method=<method>      method: one of either 'heartbeat' or 'nagios'"; got != expected {
		t.Errorf("Error stringifying option:\ngot:      " + got + "\nexpected: " + expected)
	}
}

func TestBasicOutputWithExampleHelpTextValue(t *testing.T) {
	if got, expected := (Option{"method|m", "method", Optional | NoLongOpt, "heartbeat"}.HelpText(20)),
		"    -m <method>                method (e.g. heartbeat)"; got != expected {
		t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method", "method", Optional, "heartbeat"}.HelpText(20)),
		"        --method=<method>      method (e.g. heartbeat)"; got != expected {
		t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method|m", "method", Optional, "heartbeat"}.HelpText(20)),
		"    -m, --method=<method>      method (e.g. heartbeat)"; got != expected {
		t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method|m", "method", Required | NoLongOpt, "heartbeat"}.HelpText(20)),
		"    -m <method>                method (e.g. heartbeat)"; got != expected {
		t.Errorf("Error stringifying required option:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method", "method", Required, "heartbeat"}.HelpText(20)),
		"        --method=<method>      method (e.g. heartbeat)"; got != expected {
		t.Errorf("Error stringifying required option:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method|m", "method", Required, "heartbeat"}.HelpText(20)),
		"    -m, --method=<method>      method (e.g. heartbeat)"; got != expected {
		t.Errorf("Error stringifying required option:\ngot:      " + got + "\nexpected: " + expected)
	}
}

func TestBasicOutputWithDefaultHelpText(t *testing.T) {
	if got, expected := (Option{"method|m", "method", Optional | ExampleIsDefault | NoLongOpt, "heartbeat"}.HelpText(20)),
		"    -m <method>                method (default: heartbeat)"; got != expected {
		t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"configfile|c", "method", IsConfigFile | ExampleIsDefault, "/etc/cnf.conf"}.HelpText(20)),
		"    -c, --configfile=<configfile> method (default: /etc/cnf.conf)"; got != expected {
		t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method", "method", Optional | ExampleIsDefault, "heartbeat"}.HelpText(20)),
		"        --method=<method>      method (default: heartbeat)"; got != expected {
		t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"method|m", "method", Optional | ExampleIsDefault, "heartbeat"}.HelpText(20)),
		"    -m, --method=<method>      method (default: heartbeat)"; got != expected {
		t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected)
	}
}

func TestOutputWithEnvVar(t *testing.T) {
	if got, expected := (Option{"method|m|METHOD", "method", Optional | ExampleIsDefault | NoLongOpt, "heartbeat"}.HelpText(20)),
		"    -m <method>                method (default: heartbeat); setable via $METHOD"; got != expected {
		t.Errorf("Error stringifying optional option:\ngot:      " + got + "\nexpected: " + expected)
	}
}

func TestOutputArgument(t *testing.T) {
	if got, expected := (Option{"logfile", "", IsArg, ""}.HelpText(20)),
		""; got != expected {
		t.Errorf("Error stringifying argument:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"logfile", "dump log into this file", IsArg, ""}.HelpText(20)),
		"    <logfile>                  dump log into this file"; got != expected {
		t.Errorf("Error stringifying argument:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"logfile", "dump log into this file", IsArg, "/tmp/foo"}.HelpText(20)),
		"    <logfile>                  dump log into this file (e.g. /tmp/foo)"; got != expected {
		t.Errorf("Error stringifying argument:\ngot:      " + got + "\nexpected: " + expected)
	}
}

func TestOutputPassThrough(t *testing.T) {
	if got, expected := (Option{"pass through args", "pass through arguments", IsPassThrough, ""}.Usage()),
		"-- <pass through args>"; got != expected {
		t.Errorf("Error stringifying argument:\ngot:      " + got + "\nexpected: " + expected)
	}

	if got, expected := (Option{"pass through args", "pass through arguments", IsPassThrough, ""}.HelpText(20)),
		"    <pass through args>        pass through arguments"; got != expected {
		t.Errorf("Error stringifying argument:\ngot:      " + got + "\nexpected: " + expected)
	}
}
