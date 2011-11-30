package getopt

import "testing"

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

func TestIsShortOpt(t *testing.T) {
  if opt, _, _   := parseShortOpt("-a"); opt != "a"        { t.Errorf("'-a' option: opt not parsed correctly") }
  if _, val, _   := parseShortOpt("-a"); val != ""         { t.Errorf("'-a' option: val not parsed correctly") }
  if _, _, found := parseShortOpt("-a"); found == false    { t.Errorf("'-a' option: found not set correctly") }

  if opt, _, _   := parseShortOpt("-abc"); opt != "a"      { t.Errorf("'-abc' option: opt not parsed correctly") }
  if _, val, _   := parseShortOpt("-abc"); val != "bc"     { t.Errorf("'-abc' option: val not parsed correctly") }
  if _, _, found := parseShortOpt("-abc"); found == false  { t.Errorf("'-abc' option: found not set correctly") }

  if opt, _, _   := parseShortOpt("-"); opt != ""          { t.Errorf("'-' option: opt not parsed correctly") }
  if _, val, _   := parseShortOpt("-"); val != ""          { t.Errorf("'-' option: val not parsed correctly") }
  if _, _, found := parseShortOpt("-"); found == true      { t.Errorf("'-' option: found not set correctly") }

  if opt, _, _   := parseShortOpt("--"); opt != ""         { t.Errorf("'--' option: opt not parsed correctly") }
  if _, val, _   := parseShortOpt("--"); val != ""         { t.Errorf("'--' option: val not parsed correctly") }
  if _, _, found := parseShortOpt("--"); found == true     { t.Errorf("'--' option: found not set correctly") }

  if opt, _, _   := parseShortOpt(""); opt != ""           { t.Errorf("'' option: opt not parsed correctly") }
  if _, val, _   := parseShortOpt(""); val != ""           { t.Errorf("'' option: val not parsed correctly") }
  if _, _, found := parseShortOpt(""); found == true       { t.Errorf("'' option: found not set correctly") }

  if opt, _, _   := parseShortOpt("a"); opt != ""          { t.Errorf("'a' option: opt not parsed correctly") }
  if _, val, _   := parseShortOpt("a"); val != ""          { t.Errorf("'a' option: val not parsed correctly") }
  if _, _, found := parseShortOpt("a"); found == true      { t.Errorf("'a' option: found not set correctly") }

  if opt, _, _   := parseShortOpt("aaaaa"); opt != ""      { t.Errorf("'aaaaa' option: opt not parsed correctly") }
  if _, val, _   := parseShortOpt("aaaaa"); val != ""      { t.Errorf("'aaaaa' option: val not parsed correctly") }
  if _, _, found := parseShortOpt("aaaaa"); found == true  { t.Errorf("'aaaaa' option: found not set correctly") }
}

func TestIsValue(t *testing.T) {
  if !isValue("value") { t.Errorf("'value' not recognized as value") }
  if !isValue(" ") { t.Errorf("' ' not recognized as value") }
  if !isValue(" ") { t.Errorf("'  ' not recognized as value") }
  if isValue("--")     { t.Errorf("'--' identified as value") }
  if isValue("-a")     { t.Errorf("'-a' identified as value") }
  if isValue("--longopt")     { t.Errorf("'--longopt' identified as value") }
  if isValue("-abc")     { t.Errorf("'-abc' identified as value") }
  if isValue("--longopt=value")     { t.Errorf("'--longopt=value' identified as value") }
}

func TestArgumentsEnd(t *testing.T) {
  if !argumentsEnd("--") { t.Errorf("'--' not recognized as arguments end") }
}
