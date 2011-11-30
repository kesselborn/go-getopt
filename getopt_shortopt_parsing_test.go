package getopt

import(
  "testing"
)

func TestIsShortOpt(t *testing.T) {
  if opt, _, _   := isShortOpt("-a"); opt != "a"     { t.Errorf("'-a' option: opt not parsed correctly") }
  if _, val, _   := isShortOpt("-a"); val != ""      { t.Errorf("'-a' option: val not parsed correctly") }
  if _, _, found := isShortOpt("-a"); found == false { t.Errorf("'-a' option: found not set correctly") }

  if opt, _, _   := isShortOpt("-abc"); opt != "a"     { t.Errorf("'-abc' option: opt not parsed correctly") }
  if _, val, _   := isShortOpt("-abc"); val != "bc"    { t.Errorf("'-abc' option: val not parsed correctly") }
  if _, _, found := isShortOpt("-abc"); found == false { t.Errorf("'-abc' option: found not set correctly") }

  if opt, _, _   := isShortOpt("-"); opt != ""      { t.Errorf("'-' option: opt not parsed correctly") }
  if _, val, _   := isShortOpt("-"); val != ""      { t.Errorf("'-' option: val not parsed correctly") }
  if _, _, found := isShortOpt("-"); found == true  { t.Errorf("'-' option: found not set correctly") }

  if opt, _, _   := isShortOpt("--"); opt != ""      { t.Errorf("'--' option: opt not parsed correctly") }
  if _, val, _   := isShortOpt("--"); val != ""      { t.Errorf("'--' option: val not parsed correctly") }
  if _, _, found := isShortOpt("--"); found == true  { t.Errorf("'--' option: found not set correctly") }

  if opt, _, _   := isShortOpt(""); opt != ""      { t.Errorf("'' option: opt not parsed correctly") }
  if _, val, _   := isShortOpt(""); val != ""      { t.Errorf("'' option: val not parsed correctly") }
  if _, _, found := isShortOpt(""); found == true  { t.Errorf("'' option: found not set correctly") }

  if opt, _, _   := isShortOpt("a"); opt != ""      { t.Errorf("'a' option: opt not parsed correctly") }
  if _, val, _   := isShortOpt("a"); val != ""      { t.Errorf("'a' option: val not parsed correctly") }
  if _, _, found := isShortOpt("a"); found == true  { t.Errorf("'a' option: found not set correctly") }

  if opt, _, _   := isShortOpt("aaaaa"); opt != ""      { t.Errorf("'aaaaa' option: opt not parsed correctly") }
  if _, val, _   := isShortOpt("aaaaa"); val != ""      { t.Errorf("'aaaaa' option: val not parsed correctly") }
  if _, _, found := isShortOpt("aaaaa"); found == true  { t.Errorf("'aaaaa' option: found not set correctly") }
}

