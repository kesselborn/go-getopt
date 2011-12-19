package getopt

import (
	"testing"
	"fmt"
	"os"
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

func TestOptionCascade(t *testing.T) {
	options := Options{
		{"foo|f|FOO", "bogus var", Required | ExampleIsDefault, "yamalla"},
		{"config|c|CONFIG", "configuration file", Required, "/tmp/foo"},
	}

	os.Args = []string{"prog"}
	if opts, _, _, _ := options.ParseCommandLine("", 0); opts["foo"].String != "yamalla" {
		fmt.Printf("XXX%#v\n", opts)
		t.Errorf("did not recognize default value for 'foo', expected: 'yamalla', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "-c./config_sample.conf"}
	if opts, _, _, _ := options.ParseCommandLine("", 0); opts["foo"].String != "bar" {
		t.Errorf("did not read value for 'foo' from config file (1), expected: 'bar', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "-c ./config_sample.conf"}
	if opts, _, _, _ := options.ParseCommandLine("", 0); opts["foo"].String != "bar" {
		t.Errorf("did not read value for 'foo' from config file (2), expected: 'bar', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "--config=./config_sample.conf"}
	if opts, _, _, _ := options.ParseCommandLine("", 0); opts["foo"].String != "bar" {
		t.Errorf("did not read value for 'foo' from config file (3), expected: 'bar', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "--config=./config_sample.conf"}
	os.Setenv("FOO", "bar2")
	if opts, _, _, _ := options.ParseCommandLine("", 0); opts["foo"].String != "bar2" {
		t.Errorf("did not recognize value for 'foo' when set via ENV var + config file, expected: 'bar2', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "--config=./config_sample.conf", "-fbar3"}
	os.Setenv("FOO", "bar2")
	if opts, _, _, _ := options.ParseCommandLine("", 0); opts["foo"].String != "bar3" {
		t.Errorf("did not recognize value for 'foo' when set via ENV, config file and option (1), expected: 'bar3', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "--config=./config_sample.conf", "-f bar3"}
	os.Setenv("FOO", "bar2")
	if opts, _, _, _ := options.ParseCommandLine("", 0); opts["foo"].String != "bar3" {
		t.Errorf("did not recognize value for 'foo' when set via ENV, config file and option (2), expected: 'bar3', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "--config=./config_sample.conf", "--foo=bar3"}
	os.Setenv("FOO", "bar2")
	if opts, _, _, _ := options.ParseCommandLine("", 0); opts["foo"].String != "bar3" {
		t.Errorf("did not recognize value for 'foo' when set via ENV, config file and option (1), expected: 'bar3', got: '" + opts["foo"].String + "'")
	}

}

func TestConfigFileNotFoundErrors(t *testing.T) {
	t.Errorf("plunk")
}
