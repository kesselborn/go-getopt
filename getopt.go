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

func (optionsDefinition Options) parse(args []string) (map[string] string, *GetOptError) {
  options := make(map[string] string)
  requiredValues := make([]string, 0)

  for _, option := range optionsDefinition {
    if option.arg_type == Required {
      requiredValues = append(requiredValues, option.long_opt)
    }
  }

  for i:=1; i<len(args); i++ { // args[0] is no opt
    if args[i] == "--" || args[i][0] != '-' {
      break
    }
  }

  for _, option := range requiredValues {
    if _, found := options[option]; found == false {
      return options, &GetOptError{1, "required option " + option + " is not set"}
    }
  }

  return options, nil
}
