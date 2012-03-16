// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
	"os"
	"testing"
)

func TestConsistencyChecking(t *testing.T) {
	os.Args = []string{"prog"}
	if _, _, _, err := (Options{"", Definitions{{"verbose", "...", Optional | Required, ""}}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags Optional & Required did not raise error!")
	}
	if _, _, _, err := (Options{"", Definitions{{"verbose", "...", Flag | ExampleIsDefault, ""}}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags Flag & ExampleIsDefault did not raise error!")
	}
	if _, _, _, err := (Options{"", Definitions{{"verbose", "...", Required | ExampleIsDefault, ""}}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags Required & ExampleIsDefault did not raise error!")
	}
	if _, _, _, err := (Options{"", Definitions{{"verbose", "...", NoLongOpt, ""}}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags NoLongOpt and without a short opt is not possible")
	}
	if _, _, _, err := (Options{"", Definitions{{"verbose", "...", IsArg | Flag, ""}}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags IsArg & Flag")
	}
	if _, _, _, err := (Options{"", Definitions{{"verbose", "...", IsArg, ""}}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags IsArg without Required or Optional")
	}
	if _, _, _, err := (Options{"", Definitions{{"arg1", "...", IsArg | Optional, ""}, {"arg1", "...", IsArg | Required, ""}}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("required arg following an optional arg did not raise error")
	}
	if _, _, _, err := (Options{"", Definitions{{"arg1", "...", IsArg | Optional, ""}, {"arg1", "...", IsArg | Required, ""}, {"arg3", "...", IsArg | Optional, ""}}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("required arg following an optional arg did not raise error")
	}

	os.Args = []string{"prog", "lirum", "larum"}
	if _, _, _, err := (Options{"", Definitions{{"arg1", "...", IsArg | Required, ""}, {"arg2", "...", IsArg | Optional, ""}}}.ParseCommandLine()); err != nil {
		t.Errorf("required arg followed by an optional arg did raise error: %#v", err)
	}

	if _, _, _, err := (Options{"", Definitions{{"a", "...", Flag, ""}, {"a", "...", Flag, ""}}}).ParseCommandLine(); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("double usage of short opt not detected")
	}

	if _, _, _, err := (Options{"", Definitions{{"a|arg1", "...", Flag, ""}, {"b|arg1", "...", Flag, ""}}}).ParseCommandLine(); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("double usage of long opt not detected")
	}

	if _, _, _, err := (Options{"", Definitions{{"a|arg1|FOO", "...", Flag, ""}, {"b|arg2|FOO", "...", Flag, ""}}}).ParseCommandLine(); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("double usage of env var not detected")
	}
}
