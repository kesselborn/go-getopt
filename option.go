// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt
import "strings"

const Required = 1
const Optional = 2
const Flag = 4
const NoLongOpt = 8
const ExampleIsDefault = 16
const IsArg = 32
const Argument = 64
const Usage = 128
const Help = 256
const IsPassThrough = 512

type Option struct {
  option_definition string
  description string
  flags int
  defaultValue interface{}
}

func (option Option) eq(other Option) bool {
  return option.option_definition == other.option_definition &&
         option.description == other.description &&
         option.flags == other.flags &&
         option.defaultValue == other.defaultValue
}

func (option Option) neq(other Option) bool {
  return !option.eq(other)
}

func (option Option) Key() (key string) {
  return strings.Split(option.option_definition, "|")[0]
}


func (option Option) LongOpt() (longOpt string) {
  if option.flags & NoLongOpt == 0 {
    longOpt = option.Key()
  }

  return longOpt
}

func (option Option) HasLongOpt() (result bool) {
  return option.LongOpt() != ""
}

func (option Option) ShortOpt() (shortOpt string) {
  token := strings.Split(option.option_definition, "|")

  if len(token) > 1 {
    shortOpt = token[1]
  }

  return shortOpt
}

func (option Option) HasShortOpt() (result bool) {
  return option.ShortOpt() != ""
}


func (option Option) EnvVar() (envVar string) {
  token := strings.Split(option.option_definition, "|")

  if len(token) > 2 {
    envVar = token[2]
  }

  return envVar
}

func (option Option) HasEnvVar() (result bool) {
  return option.EnvVar() != ""
}

