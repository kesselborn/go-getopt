// Copyright (c) 2011, SoundCloud Ltd., Daniel Bornkessel
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Source code and contact info at http://github.com/kesselborn/go-getopt

package getopt

type SubCommandOptions map[string]Options

func (sco SubCommandOptions) flattenToOptions(subCommand string) (options Options, err *GetOptError) {
	genericOptions := sco["*"]

	if subCommandOptions, present := sco[subCommand]; present == true {

		for _, option := range genericOptions {
			options = append(options, option)
		}
		for _, option := range subCommandOptions {
			options = append(options, option)
		}
	} else {
		err = &GetOptError{UnknownSubcommand, "Unknown command: " + subCommand}
	}

	return
}

//func checkOptionsDefinitionConsistency(optionsDefinition Options) (err *GetOptError) {
//	consistencyErrorPrefix := "wrong getopt usage: "
//
//	foundOptionalArg := false
//	for _, option := range optionsDefinition {
//		switch {
//		case option.Flags&IsArg > 0 && option.Flags&Required == 0 && option.Flags&Optional == 0:
//			err = &GetOptError{ConsistencyError, consistencyErrorPrefix + "an argument must be explicitly set to be Optional or Required"}
//		case option.Flags&IsArg > 0 && option.Flags&Optional > 0:
//			foundOptionalArg = true
//		case option.Flags&IsArg > 0 && option.Flags&Required > 0 && foundOptionalArg:
//			err = &GetOptError{ConsistencyError, consistencyErrorPrefix + "a required argument can't come after an optional argument"}
//		case option.Flags&Optional > 0 && option.Flags&Required > 0:
//			err = &GetOptError{ConsistencyError, consistencyErrorPrefix + "an option can not be Required and Optional"}
//		case option.Flags&Flag > 0 && option.Flags&ExampleIsDefault > 0:
//			err = &GetOptError{ConsistencyError, consistencyErrorPrefix + "an option can not be a Flag and have ExampleIsDefault"}
//		case option.Flags&Required > 0 && option.Flags&ExampleIsDefault > 0:
//			err = &GetOptError{ConsistencyError, consistencyErrorPrefix + "an option can not be Required and have ExampleIsDefault"}
//		case option.Flags&NoLongOpt > 0 && !option.HasShortOpt() && option.Flags&IsArg == 0:
//			err = &GetOptError{ConsistencyError, consistencyErrorPrefix + "an option must have either NoLongOpt or a ShortOption"}
//		case option.Flags&Flag > 0 && option.Flags&IsArg > 0:
//			err = &GetOptError{ConsistencyError, consistencyErrorPrefix + "an option can not be a Flag and be an argument (IsArg)"}
//		}
//	}
//
//	return
//}
