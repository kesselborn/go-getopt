package getopt
import "strings"

type Option struct {
  option_definition string
  description string
  flags int
  default_value interface{}
}

func (option Option) eq(other Option) bool {
  return option.option_definition == other.option_definition &&
         option.description == other.description &&
         option.flags == other.flags &&
         option.default_value == other.default_value
}

func (option Option) neq(other Option) bool {
  return !option.eq(other)
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
