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
  set bool
}


func assignValue(referenceValue interface{}, value string) (returnValue OptionValue, err *GetOptError) {
  valType := reflect.TypeOf(referenceValue).String()
  var e os.Error

  switch valType {
    case "string":
      returnValue.String = value
    case "bool":
      returnValue.Bool, e = strconv.Atob(value)
    case "int": fallthrough
    case "int64":
      returnValue.Int, e = strconv.Atoi64(value)
    case "[]string":
      returnValue.StrArray = strings.Split(value, ",")
    case "[]int": fallthrough
    case "[]int64":
      stringArray := strings.Split(value, ",")
      returnValue.IntArray = make([]int64, len(stringArray))
      for i, value := range stringArray {
        returnValue.IntArray[i], e = strconv.Atoi64(value)
      }
    default:
      e = os.NewError("Couldn't convert '" + value + "' to type '" + valType + "'")
  }

  if e == nil {
    returnValue.set = true
  } else {
    err = &GetOptError{ OptionValueError, "Conversion Error: " + e.String()}
  }

  return
}
