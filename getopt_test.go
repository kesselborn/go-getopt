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
    opts["debug"].Bool != true {
      t.Errorf("debug flag was not set")
  }

  if opts, _, _, _ := options.parse([]string{"-d"}, "", 0 );
    opts["debug"].Bool != true || opts["verbose"].Bool != false {
      t.Errorf("verbose flag was not set to false by default")
  }

  if opts, _, _, _ := options.parse([]string{"-d", "-v"}, "", 0 );
    opts["debug"].Bool != true || opts["verbose"].Bool != true {
      t.Errorf("did not recognize all flags")
  }

}

func TestShortOptionRequiredParsing(t *testing.T) {
  options := Options{
    {"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", Required, ""},
  }

  if opts, _, _, _ := options.parse([]string{"-m", "heartbeat"}, "", 0 );
    opts["method"].String != "heartbeat" {
      t.Errorf("method optioned wasn't parsed correctly expected 'heartbeat', was '" + opts["method"].String + "'")
  }

  if opts, _, _, _ := options.parse([]string{"-mheartbeat"}, "", 0 );
    opts["method"].String != "heartbeat" {
      t.Errorf("method optioned wasn't parsed correctly expected 'heartbeat', was '" + opts["method"].String + "'")
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

func TestConcatenatedOptionsParsingSimple(t *testing.T) {
  options := Options{
    {"debug|d", "debug mode", Flag, true},
    {"verbose|v", "verbose mode", Flag, true},
    {"dryrun|D", "dry run only", Flag, true},
    {"logfile|l", "log file", Optional, ""},
    {"mode|m", "operating mode", Required, ""},
  }

  if opts, _, _, _ := options.parse([]string{"-dv"}, "", 0 );
    opts["debug"].Bool != true || opts["verbose"].Bool != true {
      t.Errorf("did not recognize all flags when concatenation options (2 options)")
  }

  if opts, _, _, _ := options.parse([]string{"-dvD"}, "", 0 );
    opts["debug"].Bool != true || opts["verbose"].Bool != true || opts["dryrun"].Bool != true {
      t.Errorf("did not recognize all flags when concatenation options (3 options)")
  }

  if _, _, _, err := options.parse([]string{"-dvD"}, "", 0 );
    err == nil || err.errorCode != MissingOption {
      t.Errorf("did not recognize a missing required option in concatenation mode")
  }

  if _, _, _, err := options.parse([]string{"-Dl"}, "", 0 );
    err == nil || err.errorCode != MissingValue {
      t.Errorf("did not realize that I missed a value")
  }

  if _, _, _, err := options.parse([]string{"-Dl", "-d"}, "", 0 );
    err == nil || err.errorCode != MissingValue {
      t.Errorf("did not realize that I missed a value")
  }

}

func TestConcatenatedOptionsParsingWithStringValueOptionAtTheEnd(t *testing.T) {
  options := Options{
    {"debug|d", "debug mode", Flag, true},
    {"verbose|v", "verbose mode", Flag, true},
    {"dryrun|D", "dry run only", Flag, true},
    {"logfile|l", "log file", Optional, ""},
    {"mode|m", "operating mode", Required, ""},
  }
  if opts, _, _, _ := options.parse([]string{"-dvDl/tmp/log.txt"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || opts["logfile"].String != "/tmp/log.txt" {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional)")
  }

  if opts, _, _, _ := options.parse([]string{"-dvDl", "/tmp/log.txt"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || opts["logfile"].String != "/tmp/log.txt" {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional separated by space)")
  }

  if opts, _, _, _ := options.parse([]string{"-dvDmdaemon"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || opts["mode"].String    != "daemon" {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required)")
  }

  if opts, _, _, _ := options.parse([]string{"-dvDm", "daemon"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || opts["mode"].String    != "daemon" {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required separated by space)")
  }

}

func TestConcatenatedOptionsParsingWithIntValueOptionAtTheEnd(t *testing.T) {
  options := Options{
    {"debug|d", "debug mode", Flag, true},
    {"verbose|v", "verbose mode", Flag, true},
    {"dryrun|D", "dry run only", Flag, true},
    {"port|p", "port", Optional, 3000},
    {"instances|i", "instances", Required, 1},
  }
  if opts, _, _, _ := options.parse([]string{"-dvDp3000"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || opts["port"].Int != 3000 {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int)")
  }

  if opts, _, _, _ := options.parse([]string{"-dvDp", "3000"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || opts["port"].Int != 3000 {
      fmt.Printf("%#v", opts)
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int separated by space)")
  }

  if opts, _, _, _ := options.parse([]string{"-dvDp", "3000"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || opts["port"].Int != 3000 {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int separated by space)")
  }

  if opts, _, _, _ := options.parse([]string{"-dvDi4"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || opts["instances"].Int    != 4 {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required)")
  }

  if opts, _, _, _ := options.parse([]string{"-dvDi", "4"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || opts["instances"].Int    != 4 {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required int separated by space)")
  }

}

func TestConcatenatedOptionsParsingWithIntArrayValueOptionAtTheEnd(t *testing.T) {
  options := Options{
    {"debug|d", "debug mode", Flag, true},
    {"verbose|v", "verbose mode", Flag, true},
    {"dryrun|D", "dry run only", Flag, true},
    {"ports|p", "ports", Optional, []int{3000, 3001, 3002}},
    {"timeouts|t", "timeouts", Required, []int{1,2,4,10,30}},
  }
  if opts, _, _, _ := options.parse([]string{"-dvDp5000,5001,5002"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || !equalIntArray(opts["ports"].IntArray, []int64{5000, 5001, 5002}) {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int array)")
  }

  if opts, _, _, _ := options.parse([]string{"-dvDp", "5000,5001,5002"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || !equalIntArray(opts["ports"].IntArray, []int64{5000, 5001, 5002}) {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 optional int array separated by space)")
  }

  if opts, _, _, _ := options.parse([]string{"-dvDt10,20,30"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || !equalIntArray(opts["timeouts"].IntArray, []int64{10,20,30}) {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required int array)")
  }

  if opts, _, _, _ := options.parse([]string{"-dvDt", "10,20,30"}, "", 0 );
    opts["debug"].Bool   != true || opts["verbose"].Bool != true ||
    opts["dryrun"].Bool  != true || !equalIntArray(opts["timeouts"].IntArray, []int64{10,20,30}) {
      t.Errorf("did not recognize all flags when concatenation options (3 flags + 1 required int array separated by space)")
  }

}

func TestEnvVars(t *testing.T) {
  t.Errorf("env var tests missing")
}
