include $(GOROOT)/src/Make.inc

TARG=getopt
GOFILES=\
					getopt.go\
					parsing_helper.go\
					option.go\
					options.go\
					option_value.go\


include $(GOROOT)/src/Make.pkg
