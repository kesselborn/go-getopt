package getopt

import "testing"

func TestConsistencyChecking(t *testing.T) {
  if _, _, _, err := (Options{{"verbose", "...", Optional | Required, ""}}.parse([]string{}, []string{}, "", 0));
    err == nil || err.errorCode != ConsistencyError {
    t.Errorf("flags Optional & Required did not raise error!")
  }
  if _, _, _, err := (Options{{"verbose", "...", Flag | ExampleIsDefault, ""}}.parse([]string{}, []string{}, "", 0));
    err == nil || err.errorCode != ConsistencyError {
    t.Errorf("flags Flag & ExampleIsDefault did not raise error!")
  }
  if _, _, _, err := (Options{{"verbose", "...", Required | ExampleIsDefault, ""}}.parse([]string{}, []string{}, "", 0));
    err == nil || err.errorCode != ConsistencyError {
    t.Errorf("flags Required & ExampleIsDefault did not raise error!")
  }
  if _, _, _, err := (Options{{"verbose", "...", IsArg | Required, ""}}.parse([]string{}, []string{}, "", 0));
    err == nil || err.errorCode != ConsistencyError {
    t.Errorf("flags IsArg & Required did not raise error!")
  }
  if _, _, _, err := (Options{{"verbose", "...", NoLongOpt, ""}}.parse([]string{}, []string{}, "", 0));
    err == nil || err.errorCode != ConsistencyError {
    t.Errorf("flags NoLongOpt and without a short opt is not possible")
  }
  if _, _, _, err := (Options{{"verbose", "...", IsArg | Flag, ""}}.parse([]string{}, []string{}, "", 0));
    err == nil || err.errorCode != ConsistencyError {
    t.Errorf("flags IsArg & Flag")
  }
}
