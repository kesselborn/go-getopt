// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

import (
	"testing"
)

func equalOptionsArray(array1 Options, array2 Options) (equal bool) {
	if len(array1) == len(array2) {
		for i := 0; i < len(array1); i++ {
			if array1[i] != array2[i] {
				break
			}
		}
		equal = true
	}

	return
}

func TestSubcommandOptionsConverter(t *testing.T) {
	sco := SubCommandOptions{
		"*": {
			{"command", "command to execute", IsSubcommand, ""}},
		"getenv": {
			{"name", "app's name", IsArg | Required, ""},
			{"key", "environment variable's name", IsArg | Required, ""}},
		"register": {
			{"name|n", "app's name", IsArg | Required, ""},
			{"deploytype|t", "deploy type (one of mount, bazapta, lxc)", Optional | ExampleIsDefault, "lxc"}},
	}

	expectedGetenvOptions := Options{
		{"command", "command to execute", IsSubcommand, ""},
		{"name", "app's name", IsArg | Required, ""},
		{"key", "environment variable's name", IsArg | Required, ""},
	}

	expectedRegisterOptions := Options{
		{"command", "command to execute", IsSubcommand, ""},
		{"name|n", "app's name", IsArg | Required, ""},
		{"deploytype|t", "deploy type (one of mount, bazapta, lxc)", Optional | ExampleIsDefault, "lxc"},
	}

	if _, err := sco.flattenToOptions("getenv"); err != nil {
		t.Errorf("conversion SubCommandOptions -> Options failed (getenv); \nGot the following error: %s", err.Message)
	}

	if options, _ := sco.flattenToOptions("getenv"); equalOptionsArray(options, expectedGetenvOptions) == false {
		t.Errorf("conversion SubCommandOptions -> Options failed (getenv); \nGot\n\t#%#v#\nExpected:\n\t#%#v#\n", options, expectedGetenvOptions)
	}

	if _, err := sco.flattenToOptions("register"); err != nil {
		t.Errorf("conversion SubCommandOptions -> Options failed (register); \nGot the following error: %s", err.Message)
	}

	if options, _ := sco.flattenToOptions("register"); equalOptionsArray(options, expectedRegisterOptions) == false {
		t.Errorf("conversion SubCommandOptions -> Options failed (register); \nGot\n\t#%#v#\nExpected:\n\t#%#v#\n", options, expectedGetenvOptions)
	}

	if _, err := sco.flattenToOptions("nonexistantsubcommand"); err.ErrorCode != UnknownSubcommand {
		t.Errorf("non existant sub command didn't throw error")
	}

}
