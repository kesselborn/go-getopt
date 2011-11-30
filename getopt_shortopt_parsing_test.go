package getopt

import(
  "testing"
)

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

