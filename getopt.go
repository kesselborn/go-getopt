package getopt

import(
  "strings"
)

const Required = 1
const Optional = 2
const Flag = 4
const NoLongOpt = 8

const InvalidOption = 1

const OPTIONS_SEPARATOR = "--"

type Option struct {
  option_definition string
  description string
  flags int
  default_value string
}

func (option Option) Key() (key string) {
  return strings.Split(option.option_definition, "|")[0]
}


func (option Option) LongOpt() (longOpt string) {
  if option.flags & NoLongOpt == 0 {
    longOpt = option.Key()
  }

  return longOpt
}

func (option Option) ShortOpt() (shortOpt string) {
  token := strings.Split(option.option_definition, "|")

  if len(token) > 1 {
    shortOpt = token[1]
  }

  return shortOpt
}

func (option Option) EnvVar() (envVar string) {
  token := strings.Split(option.option_definition, "|")

  if len(token) > 2 {
    envVar = token[2]
  }

  return envVar
}


type Options []Option;
type GetOptError struct {
  errorCode int
  message string
}

func parseShortOpt(option string) (opt string, val string, found bool) {
  if len(option) > 1 && option[0] == '-' && option[1] >= 'A' && option[1] <= 'z' {
    found = true
    opt = option[1:2]
    if len(option) > 2 {
      val = option[2:]
    }

  }

  return opt, val, found
}


func parseLongOpt(option string) (opt string, val string, found bool) {
  if len(option) > 3 && option[0:2] == "--" {
    found = true

    optTokens := strings.Split(option[2:], "=")

    opt = optTokens[0]

    if len(optTokens) > 1 {
      val = optTokens[1]
    }
  }

  return opt, val, found
}

func isValue(option string) bool {
  _, _, isShortOpt := parseShortOpt(option)
  _, _, isLongOpt := parseLongOpt(option)

  return !isShortOpt && !isLongOpt && !argumentsEnd(option)
}

func argumentsEnd(option string) bool {
  return option == "--"
}

func (optionsDefinition Options) parse(args []string) (map[string] string, []string, []string, *GetOptError) {
  return make(map[string] string), []string{}, []string{}, nil
}
