package getopt

const Required = 1
const Optional = 2
const Flag = 4
const NoLongOpt = 8

const InvalidOption = 1

const OPTIONS_SEPARATOR = "--"


type GetOptError struct {
  errorCode int
  message string
}
//func (optionsDefinition Options) parse(args []string) (map[string] string, []string, []string, *GetOptError) {
//  return make(map[string] string), []string{}, []string{}, nil
//}

//vim:fdm=manual
