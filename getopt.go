package getopt

const Required = 1
const Optional = 2
const Flag = 4
const NoLongOpt = 8

const InvalidOption = 1
const MissingValue = 2
const InvalidValue = 4
const MissingOption = 8

const OPTIONS_SEPARATOR = "--"


type GetOptError struct {
  errorCode int
  message string
}

func (optionsDefinition Options) parse(args []string,
                                       description string,
                                       flags int) (
                                  options map[string] string,
                                  arguments []string,
                                  passThrough []string,
                                  err *GetOptError) {
  options = make(map[string] string)

  for _, flagOption := range optionsDefinition.FlagOptions() {
    options[flagOption] = "false"
  }

  for i:=0; i < len(args); i++ {
    var opt, val string
    var found bool

    token := args[i]

    if argumentsEnd(token) {
      break
    }

    if isValue(token) {
      arguments = args[i:]
      break
    }

    opt, val, found = parseShortOpt(token)

    if !found {
      opt, val, found = parseLongOpt(token)
    }

    currentOption, found := optionsDefinition.FindOption(opt)
    key := currentOption.Key()

    if !found {
      err = &GetOptError{ InvalidOption, "invalid option '" + token + "'" }
      break
    }

    if optionsDefinition.IsFlag(opt) {
      options[key] = "true"
    } else {
      if val == "" {
        if len(args) > i + 1 && isValue(args[i + 1]) {
          i = i + 1
          val = args[i]
        } else {
          err = &GetOptError{ MissingValue, "Option '" + token + "' needs a value" }
          break
        }
      }


      if !isValue(val) {
        err = &GetOptError{ InvalidValue, "Option '" + token + "' got invalid value: '" + val + "'" }
        break
      }

      options[key] = val
    }

  }

  if err == nil {
    for _, requiredOption := range optionsDefinition.RequiredOptions() {
      if options[requiredOption] == "" {
        err = &GetOptError{ MissingOption, "Option '" + requiredOption + "' is missing" }
        break
      }
    }

  }

  return options, arguments, passThrough, err
}
