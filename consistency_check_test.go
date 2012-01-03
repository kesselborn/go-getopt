// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt


package getopt

import (
	"testing"
	"os"
)

func TestConsistencyChecking(t *testing.T) {
	os.Args = []string{"prog"}
	if _, _, _, err := (Options{{"verbose", "...", Optional | Required, ""}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags Optional & Required did not raise error!")
	}
	if _, _, _, err := (Options{{"verbose", "...", Flag | ExampleIsDefault, ""}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags Flag & ExampleIsDefault did not raise error!")
	}
	if _, _, _, err := (Options{{"verbose", "...", Required | ExampleIsDefault, ""}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags Required & ExampleIsDefault did not raise error!")
	}
	if _, _, _, err := (Options{{"verbose", "...", NoLongOpt, ""}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags NoLongOpt and without a short opt is not possible")
	}
	if _, _, _, err := (Options{{"verbose", "...", IsArg | Flag, ""}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags IsArg & Flag")
	}
	if _, _, _, err := (Options{{"verbose", "...", IsArg, ""}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags IsArg without Required or Optional")
	}
	if _, _, _, err := (Options{{"arg1", "...", IsArg | Optional, ""}, {"arg1", "...", IsArg | Required, ""}}.ParseCommandLine()); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("required arg following an optional arg did not raise error")
	}

	os.Args = []string{"prog", "lirum", "larum"}
	if _, _, _, err := (Options{{"arg1", "...", IsArg | Required, ""}, {"arg1", "...", IsArg | Optional, ""}}.ParseCommandLine()); err != nil {
		t.Errorf("required arg followed by an optional arg did raise error: %#v", err)
	}
}
