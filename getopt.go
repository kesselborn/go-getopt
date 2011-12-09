package getopt
import (
  "strings"
  //"fmt"
)

const InvalidOption = 1
const MissingValue = 2
const InvalidValue = 4
const MissingOption = 8
const OptionValueError = 16
const ConsistencyError = 32

const OPTIONS_SEPARATOR = "--"


type GetOptError struct {
  errorCode int
  message string
}

func mapifyEnviron(environment []string) (envArray map[string] string) {
  envArray = make(map[string] string)

  for _, cur := range environment {
    envVar := strings.Split(cur, "=")
    if len(envVar) > 1 {
      envArray[envVar[0]] = envVar[1]
    }
  }

  return
}

func (optionsDefinition Options) setOverwrites(options map[string] OptionValue, overwrites []string) (err *GetOptError) {
  overwritesMap := mapifyEnviron(overwrites)
  acceptedEnvVars := make(map[string] Option)

  for _, opt := range optionsDefinition {
    if value := opt.EnvVar(); value != "" {
      acceptedEnvVars[value] = opt
    }
  }

  for key, acceptedEnvVar := range acceptedEnvVars {
    if value := overwritesMap[key]; value != "" {
      options[acceptedEnvVar.LongOpt()], err = assignValue(acceptedEnvVar.defaultValue, value)
      if err != nil {
        break
      }
    }
  }

  return
}

func checkOptionsDefinitionConsistency(optionsDefinition Options) (err *GetOptError) {

  for _, option := range optionsDefinition {
    switch {
      case option.flags & Optional > 0 && option.flags & Required > 0:
        err = &GetOptError{ ConsistencyError, "an option can not be Required and Optional" }
      case option.flags & Flag > 0 && option.flags & ExampleIsDefault > 0:
        err = &GetOptError{ ConsistencyError, "an option can not be a Flag and have ExampleIsDefault" }
      case option.flags & Required > 0 && option.flags & ExampleIsDefault > 0:
        err = &GetOptError{ ConsistencyError, "an option can not be Required and have ExampleIsDefault" }
      case option.flags & Required > 0 && option.flags & IsArg > 0:
        err = &GetOptError{ ConsistencyError, "an option can not be Required and be an argument (IsArg)" }
      case option.flags & NoLongOpt > 0 && !option.HasShortOpt() && option.flags & IsArg == 0 :
        err = &GetOptError{ ConsistencyError, "an option must have either NoLongOpt or a ShortOption" }
      case option.flags & Flag > 0 && option.flags & IsArg > 0:
        err = &GetOptError{ ConsistencyError, "an option can not be a Flag and be an argument (IsArg)" }
    }
  }

  return
}


func (optionsDefinition Options) parse(args []string,
                                       defaults []string,
                                       description string,
                                       flags int) (
                                  options map[string] OptionValue,
                                  arguments []string,
                                  passThrough []string,
                                  err *GetOptError) {

  if err = checkOptionsDefinitionConsistency(optionsDefinition); err == nil {
    options = make(map[string] OptionValue)
    arguments = make([]string, 0)

    for _, option := range optionsDefinition {
      switch {
        case option.flags & Flag != 0:                // all flags are false by default
          options[option.Key()], err = assignValue(false, "false")
        case option.flags & ExampleIsDefault != 0:    // set default
          options[option.Key()], err = assign(option.defaultValue)
      }
    }

    // set overwrites
    err = optionsDefinition.setOverwrites(options, defaults)

    if err == nil {
      for i:=0; i < len(args); i++ {
        var opt, val string
        var found bool

        token := args[i]

        if argumentsEnd(token) {
          passThrough = args[i:]
          break
        }

        if isValue(token) {
          arguments = append(arguments, token)
          continue
        }

        opt, val, found = parseShortOpt(token)

        if found {
          buffer := token
          for found && optionsDefinition.IsFlag(opt) && len(buffer) > 2 {
            // concatenated options ... continue parsing
            currentOption, _ := optionsDefinition.FindOption(opt)
            key := currentOption.Key()

            options[key], err = assignValue(currentOption.defaultValue, "true")

            // make it look as if we have a normal option with a '-' prefix
            buffer = "-" + buffer[2:]
            opt, val, found = parseShortOpt(buffer)
          }

        } else {
          opt, val, found = parseLongOpt(token)
        }

        currentOption, found := optionsDefinition.FindOption(opt)
        key := currentOption.Key()

        if !found {
          err = &GetOptError{ InvalidOption, "invalid option '" + token + "'" }
          break
        }

        if optionsDefinition.IsFlag(opt) {
          options[key], err = assignValue(true, "true")
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

          options[key], err = assignValue(currentOption.defaultValue, val)
        }

      }
    }

    if err == nil {
      for _, requiredOption := range optionsDefinition.RequiredOptions() {
        if options[requiredOption].set == false {
          err = &GetOptError{ MissingOption, "Option '" + requiredOption + "' is missing" }
          break
        }
      }
    }
  }

  return
}
