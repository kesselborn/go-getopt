package getopt

import(
  "strings"
  "fmt"
)

func (option Option) String(longOptLength int) (output string){
  fmtStringLongAndShort := fmt.Sprintf("\t-%%-1s, --%%-%ds %%s", longOptLength) // "p", "port=PORT", "the port that should be used"
  fmtStringShort        := fmt.Sprintf("\t-%%-1s%%-%ds    %%s", longOptLength) // "p", "PORT", "the port that should be used"
  fmtStringLong         := fmt.Sprintf("\t    --%%-%ds %%s", longOptLength)   // "port=PORT", "the port that should be used"

  switch {
    case option.HasLongOpt() && option.HasShortOpt():
      output = fmt.Sprintf(fmtStringLongAndShort, option.ShortOptString(), option.LongOptString(), option.Description())
    case option.HasShortOpt():
      output = fmt.Sprintf(fmtStringShort, option.ShortOptString(), strings.ToUpper(option.Key()), option.Description())
    case option.HasLongOpt():
      output = fmt.Sprintf(fmtStringLong, option.LongOptString(), option.Description())
  }

  return output
}

func (option Option) LongOptString() (longOptString string) {
  if option.HasLongOpt() {
    longOptString = option.LongOpt()

    if option.flags & Flag == 0 {
      longOptString = longOptString + "=" + strings.ToUpper(option.LongOpt())
    }
  }

  return
}

func (option Option) ShortOptString() (shortOptString string) {
  if option.HasShortOpt() {
    shortOptString = option.ShortOpt()

    if option.flags & Flag == 0 && !option.HasLongOpt() {
      shortOptString = shortOptString + " " + strings.ToUpper(option.LongOpt())
    }
  }

  return
}

func (option Option) Description() (description string) {
  description = option.description

  defaultValue := fmt.Sprintf("%v", option.default_value)
  // %v for arrays prints something like [3 4 5] ... let's transform that to 3,4,5:
  defaultValue = strings.Replace(strings.Replace(strings.Replace(defaultValue,"[", "",-1), "]", "", -1), " ", ",", -1)

  if defaultValue != "" {
    switch {
      case option.flags & Optional > 0 && option.flags & ExampleIsDefault > 0:
        description = description + " (default: " + defaultValue + ")"
      case option.flags & Required > 0 || option.flags & Optional > 0:
        description = description + " (e.g. " + defaultValue + ")"
    }
  }

  if option.HasEnvVar() {
    description = description + "; setable via $" + option.EnvVar()
  }

  return
}

//
//func main() {
//  longOptLength := 50
//  fmtStringLongAndShort := fmt.Sprintf("-%%s, --%%-%ds %%s\n", longOptLength)   // "p", "port=PORT", "the port that should be used"
//  fmtStringShort        := fmt.Sprintf("-%%s %%-%ds    %%s\n", longOptLength) // "p", "PORT", "the port that should be used"
//  fmtStringLong         := fmt.Sprintf("    --%%-%ds %%s\n", longOptLength)   // "port=PORT", "the port that should be used"
//
//  fmt.Printf(fmtStringLongAndShort , "p", "port=PORT", "the port that should be used")
//  fmt.Printf(fmtStringShort, "p", "PORT", "the port that should be used")
//  fmt.Printf(fmtStringLong, "port=PORT", "the port that should be used")
//  fmt.Printf(fmtStringLongAndShort , "d", "debug", "switch on debugging")
//  fmt.Printf(fmtStringShort, "d", "", "switch on debug")
//  fmt.Printf(fmtStringLong, "debug", "switch on debug")
//}

