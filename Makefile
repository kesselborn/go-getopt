include $(GOROOT)/src/Make.inc

TARG=github.com/kesselborn/go-getopt
GOFILES=\
					getopt.go\
					parsing_helper.go\
					option.go\
					options.go\
					option_value.go\
					option_stringifier.go\
					config_file.go\

default: all
	gofmt -s=true -w *.go
	6g -V > TESTED_GO_RELEASE

include $(GOROOT)/src/Make.pkg
