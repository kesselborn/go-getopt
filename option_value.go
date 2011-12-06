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

func assign(value interface{}) (returnValue OptionValue, err *GetOptError) {
  valType := reflect.TypeOf(value).String()
  var e os.Error

  // mmm ...there should be an easier way
  switch valType {
    case "string":
      returnValue.String = value.(string)
    case "bool":
      returnValue.Bool = value.(bool)
    case "int":
      returnValue.Int = int64(value.(int))
    case "int64":
      returnValue.Int = value.(int64)
    case "[]string":
      returnValue.StrArray = value.([]string)
    case "[]int":
      var ints []int = value.([]int)
      long_ints := make([]int64, len(ints))
      for i, integer := range ints {
        long_ints[i] = int64(integer)
      }
      returnValue.IntArray = long_ints
    case "[]int64":
      returnValue.IntArray = value.([]int64)
    default:
      e = os.NewError("Couldn't assign value of type '" + valType + "'")
  }

  if e == nil {
    returnValue.set = true
  } else {
    err = &GetOptError{ OptionValueError, "Conversion Error: " + e.String()}
  }

  return

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
