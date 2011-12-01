package getopt

import(
  "os"
  "reflect"
  "strconv"
  "strings"
)



type OptionValue struct {
  Bool bool
  String string
  Int int64
  StrArray []string
  IntArray []int64
}


func assignValue(referenceValue interface{}, value string) (returnValue OptionValue, err os.Error) {
  valType := reflect.TypeOf(referenceValue).String()

  switch valType {
    case "string":
      returnValue.String = value
    case "bool":
      returnValue.Bool, err = strconv.Atob(value)
    case "int": fallthrough
    case "int64":
      returnValue.Int, err = strconv.Atoi64(value)
    case "[]string":
      returnValue.StrArray = strings.Split(value, ",")
    case "[]int": fallthrough
    case "[]int64":
      stringArray := strings.Split(value, ",")
      returnValue.IntArray = make([]int64, len(stringArray))
      for i, value := range stringArray {
        returnValue.IntArray[i], err = strconv.Atoi64(value)
      }
    default:
      err = os.NewError("Couldn't convert '" + value + "' to type '" + valType + "'")
  }

  return
}
