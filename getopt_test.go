package getopt

import( "testing" )

func TestRequiredOptions(t *testing.T) {
  Options{
    {"port", "p", Required, "listening port", "3000"},
  }.parse([]string{"--port", "3000"})

  t.Errorf("EPIC FAIL")
}

