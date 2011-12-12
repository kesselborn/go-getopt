include $(GOROOT)/src/Make.inc

TARG=github.com/kesselborn/go-getopt
GOFILES=\
					getopt.go\
					parsing_helper.go\
					option.go\
					options.go\
					option_value.go\
					option_stringifier.go\

default: all
	6g -V > TESTED_GO_RELEASE

include $(GOROOT)/src/Make.pkg
