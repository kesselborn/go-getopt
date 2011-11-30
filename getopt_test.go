package getopt

import(
  "testing"
  "fmt"
)

func TestRequiredOptions(t *testing.T) {
  _, err := Options{
    {"port", "p", Required, "listening port", "3000"},
  }.parse([]string{"--port", "3000"})

  if err != nil {
    //t.Errorf(err.message)
  }

  foo := []string{"3", "4"}
  fmt.Printf(foo.GoString())
}

