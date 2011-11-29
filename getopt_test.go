package getopt

import(
  "testing"
)

func TestRequiredOptions(t *testing.T) {
  _, err := Options{
    {"port", "p", Required, "listening port", "3000"},
  }.parse([]string{"--port", "3000"})

  if err == nil {
    t.Errorf("expected getopt to fail")
  }
}

