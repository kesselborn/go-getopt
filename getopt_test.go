package getopt
import "testing"
import "fmt"

func printMap(_map map[string] string) {
  fmt.Println("map:")
  for key, val := range _map {
    fmt.Println(" '" + key + "': '" + val + "'")
  }
}

func TestShortOptionsFlagsParsing(t *testing.T) {
  options := Options{
    {"debug|d", "debug mode", Flag, ""},
    {"verbose|v", "verbose mode", Flag, ""},
    {"dryrun|D", "dry run only", Flag, ""},
  }

  if opts, _, _, _ := options.parse([]string{"-d"}, "", 0 );
    opts["debug"] != "true" {
      t.Errorf("debug flag was not set")
  }

  if opts, _, _, _ := options.parse([]string{"-d"}, "", 0 );
    opts["debug"] != "true" || opts["verbose"] != "false" {
      t.Errorf("verbose flag was not set to false by default")
  }

  if opts, _, _, _ := options.parse([]string{"-d", "-v"}, "", 0 );
    opts["debug"] != "true" || opts["verbose"] != "true" {
      t.Errorf("did not recognize all flags")
  }

}

func TestShortOptionRequiredParsing(t *testing.T) {
  options := Options{
    {"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", Required, ""},
  }

  if opts, _, _, _ := options.parse([]string{"-m", "heartbeat"}, "", 0 );
    opts["method"] != "heartbeat" {
      t.Errorf("method optioned wasn't parsed correctly expected 'heartbeat', was '" + opts["method"] + "'")
  }

  if opts, _, _, _ := options.parse([]string{"-mheartbeat"}, "", 0 );
    opts["method"] != "heartbeat" {
      t.Errorf("method optioned wasn't parsed correctly expected 'heartbeat', was '" + opts["method"] + "'")
  }

  if _, _, _, err := options.parse([]string{"-m"}, "", 0 );
    err == nil || err.errorCode != MissingValue {
      t.Errorf("required option without value did not raise error")
  }

  if _, _, _, err := options.parse([]string{"-m", "-x"}, "", 0 );
    err == nil || err.errorCode != MissingValue {
      t.Errorf("required option without value did not raise error")
  }

  if _, _, _, err := options.parse([]string{""}, "", 0 );
    err == nil || err.errorCode != MissingOption {
      t.Errorf("required option wasn't set")
  }

}

func TestConcatenatedOptionsParsing(t *testing.T) {
  options := Options{
    {"debug|d", "debug mode", Flag, ""},
    {"verbose|v", "verbose mode", Flag, ""},
    {"dryrun|D", "dry run only", Flag, ""},
  }

  if opts, _, _, _ := options.parse([]string{"-dv"}, "", 0 );
    opts["debug"] != "true" || opts["verbose"] != "true" {
      t.Errorf("did not recognize all flags when concatenation options")
  }

  if opts, _, _, _ := options.parse([]string{"-dvD"}, "", 0 );
    opts["debug"] != "true" || opts["verbose"] != "true" || opts["dryrun"] != "true" {
      t.Errorf("did not recognize all flags when concatenation options")
  }
}
