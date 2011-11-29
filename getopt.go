package getopt

const Required = 1
const Optional = 2
const Flag = 3

const OPTIONS_SEPARATOR = "--"

type option struct {
  long_opt string
  short_opt string
  arg_type int
  description string
  default_value interface{}
}

type Options []option;
type GetOptError struct {
  errorCode int
  message string
}

func (self Options) parse(args []string) (map[string] string, *GetOptError) {
  options := make(map[string] string)

  for i:=1; i<len(args); i++ { // args[0] is no opt
    if args[i] == "--" || args[i][0] != '-' {
      break
    }

  }

  return options, nil
}
