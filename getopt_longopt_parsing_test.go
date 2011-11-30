package getopt

import(
  "testing"
)

func TestIsLongOpt(t *testing.T) {
  if opt, _, _   := parseLongOpt("--longopt");       opt   != "longopt" { t.Errorf("'--longopt' option: opt not parsed correctly") }
  if _, val, _   := parseLongOpt("--longopt");       val   != ""        { t.Errorf("'--longopt' option: val not parsed correctly") }
  if _, _, found := parseLongOpt("--longopt");       found != true      { t.Errorf("'--longopt' option: found not set correctly") }

  if opt, _, _   := parseLongOpt("--longopt=value"); opt   != "longopt" { t.Errorf("'--longopt=value' option: opt not parsed correctly") }
  if _, val, _   := parseLongOpt("--longopt=value"); val   != "value"   { t.Errorf("'--longopt=value' option: val not parsed correctly") }
  if _, _, found := parseLongOpt("--longopt=value"); found != true      { t.Errorf("'--longopt=value' option: found not set correctly") }

  if opt, _, _   := parseLongOpt("--a"); opt   != ""    { t.Errorf("'--a' option: opt not parsed correctly") }
  if _, val, _   := parseLongOpt("--a"); val   != ""    { t.Errorf("'--a' option: val not parsed correctly") }
  if _, _, found := parseLongOpt("--a"); found != false { t.Errorf("'--a' option: found not set correctly") }

  if opt, _, _   := parseLongOpt("--");  opt   != ""    { t.Errorf("'--' option: opt not parsed correctly") }
  if _, val, _   := parseLongOpt("--");  val   != ""    { t.Errorf("'--' option: val not parsed correctly") }
  if _, _, found := parseLongOpt("--");  found != false { t.Errorf("'--' option: found not set correctly") }

  if opt, _, _   := parseLongOpt("");    opt   != ""    { t.Errorf("'' option: opt not parsed correctly") }
  if _, val, _   := parseLongOpt("");    val   != ""    { t.Errorf("'' option: val not parsed correctly") }
  if _, _, found := parseLongOpt("");    found != false { t.Errorf("'' option: found not set correctly") }

  if opt, _, _   := parseLongOpt("-");   opt   != ""    { t.Errorf("'-' option: opt not parsed correctly") }
  if _, val, _   := parseLongOpt("-");   val   != ""    { t.Errorf("'-' option: val not parsed correctly") }
  if _, _, found := parseLongOpt("-");   found != false { t.Errorf("'-' option: found not set correctly") }

  if opt, _, _   := parseLongOpt("-a");  opt   != ""    { t.Errorf("'-a' option: opt not parsed correctly") }
  if _, val, _   := parseLongOpt("-a");  val   != ""    { t.Errorf("'-a' option: val not parsed correctly") }
  if _, _, found := parseLongOpt("-a");  found != false { t.Errorf("'-a' option: found not set correctly") }
}

