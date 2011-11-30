package getopt

import(
  "testing"
)

func TestIsLongOpt(t *testing.T) {
  if opt, _   := parseLongOpt("--longopt"); opt   != "longopt" { t.Errorf("'--longopt' option: opt not parsed correctly") }
  if _, found := parseLongOpt("--longopt"); found != true      { t.Errorf("'--longopt' option: found not set correctly") }

  if opt, _   := parseLongOpt("--a"); opt   != ""              { t.Errorf("'--a' option: opt not parsed correctly") }
  if _, found := parseLongOpt("--a"); found != false           { t.Errorf("'--a' option: found not set correctly") }

  if opt, _   := parseLongOpt("--"); opt   != ""               { t.Errorf("'--' option: opt not parsed correctly") }
  if _, found := parseLongOpt("--"); found != false            { t.Errorf("'--' option: found not set correctly") }

  if opt, _   := parseLongOpt(""); opt   != ""                 { t.Errorf("'' option: opt not parsed correctly") }
  if _, found := parseLongOpt(""); found != false              { t.Errorf("'' option: found not set correctly") }

  if opt, _   := parseLongOpt("-"); opt   != ""                { t.Errorf("'-' option: opt not parsed correctly") }
  if _, found := parseLongOpt("-"); found != false             { t.Errorf("'-' option: found not set correctly") }

  if opt, _   := parseLongOpt("-a"); opt   != ""               { t.Errorf("'-a' option: opt not parsed correctly") }
  if _, found := parseLongOpt("-a"); found != false            { t.Errorf("'-a' option: found not set correctly") }
}

