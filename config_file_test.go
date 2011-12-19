package getopt

import (
	"testing"
	"fmt"
)

func TestConfigParsing(t *testing.T) {
	if _, err := readConfigFile("/i/dont/exist.conf"); err == nil || err.ErrorCode != ConfigFileNotFound {
		t.Errorf("reading a non existant config file did not return an error")
	}

	if _, err := readConfigFile("./config_sample.conf"); err != nil {
		t.Errorf("reading an existing config file failed: " + err.Message)
	}

	expected := []string{"FOO=bar", "BAR=baz"}
	if got, _ := readConfigFile("./config_sample.conf"); !equalStringArray(got, expected) {
		t.Errorf("config file was not read in correctly:\ngot:      |" + fmt.Sprintf("%#v", got) + "|\nexpected: |" + fmt.Sprintf("%#v", expected) + "|\n")
	}

	configFile, _ := readConfigFile("./config_sample.conf")
	if config := mapifyConfig(configFile); config["FOO"] != "bar" || config["BAR"] != "baz" {
		t.Errorf("config file was not mapped correctly:\ngot:      |" + fmt.Sprintf("%#v", config) + "\nexpected: |map[string] string{\"BAR\":\"baz\", \"FOO\":\"bar\"}\n")
	}
}
