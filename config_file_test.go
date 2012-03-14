package getopt

import (
	"fmt"
	"os"
	"testing"
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
	if config := mapifyEnvironment(configFile); config["FOO"] != "bar" || config["BAR"] != "baz" {
		t.Errorf("config file was not mapped correctly:\ngot:      |" + fmt.Sprintf("%#v", config) + "\nexpected: |map[string] string{\"BAR\":\"baz\", \"FOO\":\"bar\"}\n")
	}
}

func TestOptionCascade(t *testing.T) {
	options := Options{
		"",
		Definitions{{"foo|f|FOO", "bogus var", ExampleIsDefault, "yamalla"},
			{"config|c|CONFIG", "configuration file", Required | IsConfigFile, "/tmp/foo"}},
	}

	os.Args = []string{"prog"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["foo"].String != "yamalla" {
		t.Errorf("did not recognize default value for 'foo', expected: 'yamalla', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "-c./config_sample.conf"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["foo"].String != "bar" {
		t.Errorf("did not read value for 'foo' from config file (1), expected: 'bar', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "-c ./config_sample.conf"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["foo"].String != "bar" {
		t.Errorf("did not read value for 'foo' from config file (2), expected: 'bar', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "--config=./config_sample.conf"}
	if opts, _, _, _ := options.ParseCommandLine(); opts["foo"].String != "bar" {
		t.Errorf("did not read value for 'foo' from config file (3), expected: 'bar', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "--config=./config_sample.conf"}
	os.Setenv("FOO", "bar2")
	if opts, _, _, _ := options.ParseCommandLine(); opts["foo"].String != "bar2" {
		t.Errorf("did not recognize value for 'foo' when set via ENV var + config file, expected: 'bar2', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "--config=./config_sample.conf", "-fbar3"}
	os.Setenv("FOO", "bar2")
	if opts, _, _, _ := options.ParseCommandLine(); opts["foo"].String != "bar3" {
		t.Errorf("did not recognize value for 'foo' when set via ENV, config file and option (1), expected: 'bar3', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "--config=./config_sample.conf", "-f", "bar3"}
	os.Setenv("FOO", "bar2")
	if opts, _, _, _ := options.ParseCommandLine(); opts["foo"].String != "bar3" {
		t.Errorf("did not recognize value for 'foo' when set via ENV, config file and option (2), expected: 'bar3', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "--config=./config_sample.conf", "--foo=bar3"}
	os.Setenv("FOO", "bar2")
	if opts, _, _, _ := options.ParseCommandLine(); opts["foo"].String != "bar3" {
		t.Errorf("did not recognize value for 'foo' when set via ENV, config file and option (3), expected: 'bar3', got: '" + opts["foo"].String + "'")
	}

}

func TestDefaultConfigFileAndRequiredOption(t *testing.T) {
	options := Options{
		"",
		Definitions{{"foo|f|FOO", "bogus var", Required, "yamalla"},
			{"config|c|CONFIG", "configuration file", IsConfigFile | ExampleIsDefault, "./config_sample.conf"}},
	}

	os.Args = []string{"prog"}
	os.Setenv("FOO", "")
	if opts, _, _, _ := options.ParseCommandLine(); opts["foo"].String != "bar" {
		t.Errorf("did not read required value from default config file, expected: 'bar', got: '" + opts["foo"].String + "'")
	}
}

func TestDefaultConfigFile(t *testing.T) {
	options := Options{
		"",
		Definitions{{"foo|f|FOO", "bogus var", ExampleIsDefault, "yamalla"},
			{"config|c|CONFIG", "configuration file", IsConfigFile | ExampleIsDefault, "./config_sample.conf"}},
	}

	os.Args = []string{"prog"}
	os.Setenv("FOO", "")
	if opts, _, _, _ := options.ParseCommandLine(); opts["foo"].String != "bar" {
		t.Errorf("did not read value from default config file, expected: 'bar', got: '" + opts["foo"].String + "'")
	}
}

func TestDefaultConfigFileWithNoLongOpt(t *testing.T) {
	options := Options{
		"",
		Definitions{{"foo|f|FOO", "bogus var", ExampleIsDefault | NoLongOpt, "yamalla"},
			{"config|c|CONFIG", "configuration file", IsConfigFile | ExampleIsDefault, "./config_sample.conf"}},
	}

	os.Args = []string{"prog"}
	os.Setenv("FOO", "")
	if opts, _, _, _ := options.ParseCommandLine(); opts["foo"].String != "bar" {
		t.Errorf("did not read value from default config file for NoLongOpt option, expected: 'bar', got: '" + opts["foo"].String + "'")
	}
}

func TestConfigFileNotFoundErrors(t *testing.T) {
	options := Options{
		"",
		Definitions{{"foo|f|FOO", "bogus var", ExampleIsDefault, "yamalla"},
			{"config|c|CONFIG", "configuration file", IsConfigFile | ExampleIsDefault, "/i/dont/exist.conf"}},
	}

	os.Args = []string{"prog"}
	os.Setenv("FOO", "")
	if opts, _, _, err := options.ParseCommandLine(); opts["foo"].String != "yamalla" || err != nil {
		t.Errorf("did fail with non-existant default config file, error message: '" + err.Message + "', expected: 'yamalla', got: '" + opts["foo"].String + "'")
	}

	os.Args = []string{"prog", "-c", "/i/dont/but/was/set/explicitly.conf"}
	os.Setenv("FOO", "")
	if opts, _, _, err := options.ParseCommandLine(); opts["foo"].String != "yamalla" || err.ErrorCode != ConfigFileNotFound {
		t.Errorf("did fail with non-existant set config file, error message: '" + err.Message + "', expected: 'yamalla', got: '" + opts["foo"].String + "'")
	}
}
