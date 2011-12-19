// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt


package getopt

import "testing"

func TestConsistencyChecking(t *testing.T) {
	if _, _, _, err := (Options{{"verbose", "...", Optional | Required, ""}}.ParseCommandLine("", 0)); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags Optional & Required did not raise error!")
	}
	if _, _, _, err := (Options{{"verbose", "...", Flag | ExampleIsDefault, ""}}.ParseCommandLine("", 0)); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags Flag & ExampleIsDefault did not raise error!")
	}
	if _, _, _, err := (Options{{"verbose", "...", Required | ExampleIsDefault, ""}}.ParseCommandLine("", 0)); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags Required & ExampleIsDefault did not raise error!")
	}
	if _, _, _, err := (Options{{"verbose", "...", IsArg | Required, ""}}.ParseCommandLine("", 0)); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags IsArg & Required did not raise error!")
	}
	if _, _, _, err := (Options{{"verbose", "...", NoLongOpt, ""}}.ParseCommandLine("", 0)); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags NoLongOpt and without a short opt is not possible")
	}
	if _, _, _, err := (Options{{"verbose", "...", IsArg | Flag, ""}}.ParseCommandLine("", 0)); err == nil || err.ErrorCode != ConsistencyError {
		t.Errorf("flags IsArg & Flag")
	}
}
