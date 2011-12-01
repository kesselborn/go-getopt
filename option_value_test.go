package getopt

import "testing"

func equalStringArray(array1 []string, array2 []string) (equal bool) {
  if len(array1) == len(array2) {
    for i := 0; i<len(array1); i++ {
      if array1[i] != array2[i] { break }
    }
    equal = true
  }

  return
}

func equalIntArray(array1 []int64, array2 []int64) (equal bool) {
  if len(array1) == len(array2) {
    for i := 0; i<len(array1); i++ {
      if array1[i] != array2[i] { break }
    }
    equal = true
  }

  return
}

func TestStringAssignValue(t *testing.T) {
  if value, _ := assignValue("foo", "bar");      value.String      != "bar" { t.Errorf("assigning string did not work") }
  if value, _ := assignValue([]string{}, "bar"); value.StrArray[0] != "bar" { t.Errorf("assigning string to StringArr did not work") }

  if value, _ := assignValue([]string{}, "bar,baz"); !equalStringArray(value.StrArray, []string{"bar", "baz"}) {
    t.Errorf("parsing string array did not work correctly")
  }

  if _, err := assignValue(42, "bar");      err == nil   { t.Errorf("assigning string to number option did not raise error") }
  if _, err := assignValue(true, "bar");    err == nil   { t.Errorf("assigning string to bool option did not raise error") }
  if _, err := assignValue([]int{}, "bar"); err == nil   { t.Errorf("assigning string to int array option did not raise error") }

}

func TestIntAssignValue(t *testing.T) {
  if value, _ := assignValue(42, "23");      value.Int != 23 { t.Errorf("assigning int did not work") }
  if value, _ := assignValue([]int{17,4}, "23");  value.IntArray[0] != 23 { t.Errorf("assigning int to IntArray did not work") }
  if value, _ := assignValue([]int64{17,4}, "23");  value.IntArray[0] != 23 { t.Errorf("assigning int to IntArray did not work") }

  if value, _ := assignValue([]int{}, "17,4"); !equalIntArray(value.IntArray, []int64{17, 4}) {
    t.Errorf("parsing int array did not work correctly")
  }

  if value, _ := assignValue([]int64{}, "17,4"); !equalIntArray(value.IntArray, []int64{17, 4}) {
    t.Errorf("parsing int64 array did not work correctly")
  }
}

func TestBoolAssignValue(t *testing.T) {
  if value, _ := assignValue(true, "1");     value.Bool != true  { t.Errorf("assigning '1' to bool failed") }
  if value, _ := assignValue(true, "true");  value.Bool != true  { t.Errorf("assigning 'true' to bool failed") }
  if value, _ := assignValue(true, "TRUE");  value.Bool != true  { t.Errorf("assigning 'TRUE' to bool failed") }

  if value, _ := assignValue(true, "0");     value.Bool != false { t.Errorf("assigning '0' to bool failed") }
  if value, _ := assignValue(true, "false"); value.Bool != false { t.Errorf("assigning 'false' to bool failed") }
  if value, _ := assignValue(true, "FALSE"); value.Bool != false { t.Errorf("assigning 'FALSE' to bool failed") } } 
