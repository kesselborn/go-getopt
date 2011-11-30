package getopt

import(
  "testing"
)

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
