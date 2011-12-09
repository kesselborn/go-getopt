package getopt

import(
  "strings"
  "fmt"
)

func (option Option) HelpText(longOptLength int) (output string){
  fmtStringLongAndShort := fmt.Sprintf("    -%%-1s, --%%-%ds %%s", longOptLength) // "p", "port=PORT", "the port that should be used"
  fmtStringShort        := fmt.Sprintf("    -%%-1s%%-%ds    %%s", longOptLength)  // "p", "PORT", "the port that should be used"
  fmtStringLong         := fmt.Sprintf("        --%%-%ds %%s", longOptLength)     // "port=PORT", "the port that should be used"
  fmtArgument           := fmt.Sprintf("    %%-%ds       %%s", longOptLength)     // "port=PORT", "the port that should be used"

  if option.description != "" {
    switch {
      case option.flags & IsArg > 0:
        output = fmt.Sprintf(fmtArgument, strings.ToUpper(option.Key()), option.Description())
      case option.HasLongOpt() && option.HasShortOpt():
        output = fmt.Sprintf(fmtStringLongAndShort, option.ShortOptString(), option.LongOptString(), option.Description())
      case option.HasShortOpt():
        output = fmt.Sprintf(fmtStringShort, option.ShortOptString(), strings.ToUpper(option.Key()), option.Description())
      case option.HasLongOpt():
        output = fmt.Sprintf(fmtStringLong, option.LongOptString(), option.Description())
    }
  }

  return output
}

func (option Option) LongOptString() (longOptString string) {
  if option.HasLongOpt() {
    longOptString = option.LongOpt()

    if option.flags & Flag == 0  && option.flags & Usage == 0 && option.flags & Help == 0  {
      longOptString = longOptString + "=" + strings.ToUpper(option.LongOpt())
    }
  }

  return
}

func (option Option) ShortOptString() (shortOptString string) {
  if option.HasShortOpt() {
    shortOptString = option.ShortOpt()

    if option.flags & Flag == 0 && !option.HasLongOpt()  && option.flags & Usage == 0 && option.flags & Help == 0{
      shortOptString = shortOptString + " " + strings.ToUpper(option.LongOpt())
    }
  }

  return
}

func (option Option) Usage() (usageString string) {
  switch {
    case option.flags & IsArg > 0:
      usageString = strings.ToUpper(option.Key())
    case option.HasShortOpt():
      usageString = "-" + option.ShortOpt()
    default:
      usageString = "--" + option.LongOpt()
  }

  if option.flags & Flag == 0 && option.flags & IsArg == 0 && option.flags & Usage == 0 && option.flags & Help == 0  {
    var separator string
    if option.HasShortOpt() {
      separator = " "
    } else {
      separator = "="
    }

    usageString = usageString + separator + strings.ToUpper(option.Key())
  }

  if option.flags & Optional > 0 || option.flags & Help > 0 || option.flags & Usage > 0 {
    usageString = "[" + usageString + "]"
  }

  return
}

func (option Option) Description() (description string) {
  description = option.description

  defaultValue := fmt.Sprintf("%v", option.defaultValue)
  // %v for arrays prints something like [3 4 5] ... let's transform that to 3,4,5:
  defaultValue = strings.Replace(strings.Replace(strings.Replace(defaultValue,"[", "",-1), "]", "", -1), " ", ",", -1)

  if defaultValue != "" && option.defaultValue != nil {
    switch {
      case option.flags & Optional > 0 && option.flags & ExampleIsDefault > 0:
        description = description + " (default: " + defaultValue + ")"
      case option.flags & Required > 0 || option.flags & Optional > 0 || option.flags & IsArg > 0:
        description = description + " (e.g. " + defaultValue + ")"
    }
  }

  if option.HasEnvVar() {
    description = description + "; setable via $" + option.EnvVar()
  }

  return
}

