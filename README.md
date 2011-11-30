gotopt -- getopt for go
=======================

This is a getopt implementation for go that is compatible with the GNU getopt
implementation

Usage example
-------------

Ok, let's do this by example. Our example binary is a monitoring app that can
monitor the health of any binary. Calling `monitor --help` would produce the
following output:

    Usage: monitor -m <method> [-l <logfile>] [--verbose] [-d] <command> [-- <command args>]
    Monitors <command> with method <method>, logs to <logfile> and passes <command args> to <command>

    Mandatory arguments to long options are mandatory for short options too.
      -m, --method               method: one of either 'heartbeat' or 'nagios' (or set with $MON_METHOD)
      -l                         log to file <logfile>
          --verbose              show verbose output
      -d, --daemon               start in daemon mode

calling `monitor -h` will produce a condensed help:

    Usage: monitor -m <method> [-l <logfile>] [-d] [--verbose] <command> [-- <command args>]

in order to parse those options, do the following in your go program:

  1.) install **getopt**

      goinstall getopt

  2.) create setup your **getopt** in your main method:


      package getopt

      func main() {
        options, arguments, passThrough, err := Options{
          {"method|m|MON_METHOD", "method: one of either 'heartbeat' or 'nagios'", Required         , ""},
          {"logfile|l"          , "log to file <logfile>",                         Optional | NoLong, nil},
          {"verbose"            , "show verbose output",                           Flag,              false},
          {"daemon|d"           , "start in daemon mode",                          Flag,              false}
        }.parse({description: "Monitors <command> with method <method>, logs to <logfile> and passes <command args> to <command>")

        if err != nil {
          //... error handling here
        }

        // ... do stuff
      }

  calling monitor in the following way:

      monitor -mheartbeat -l /tmp/log --verbose chef-client -- -j /tmp/first-run.json

  would return the following:

      options = { method:  "heartbeat",
        logfile: "/tmp/log",
        verbose: true,
        dameon: false,
      }

      arguments = {"chef-client"}

      pass-throug = {"-j", "/tmp/first-run.json"}

Options struct explained
------------------------

The options you pass to the `Options` struct have the following structure:


    {"<longopt/map key>|<shortopt>|<ENVVAR>", "<description for help text>", <options>         , <default or example value>}

  * `longopt`: the long option name that can passed to your program with
`--<long_opt>`. Furthermore, this value will be the key under which this
value is available in the options map. Long opt values need to be separated
by a whitespace: `--logfile /tmp/log.txt`. If you don't want this option to have
a long-opt style, pass `NoLong` in the options.
  * `shortopt`: short option letter ... leave it out if you only want a
long opt style for this option. Short opt values can be separated by a
whitespace: `-l /tmp/log.txt' or nothing: `-l/tmp/log.txt'
  * environment variable: if you want to let users set this option via an
environment varialbe, put the name of the env variable here. If you want
long opt style + env variable but not short opt style, pass in
`"<longopt>||ENV_VAR" `
  * the string in description will be used to create the help text
  * `<options>`:
    * **Required**: this options is required. If it is not passed in, `parse`
will return an error
    * **Optional**: can be set
    * **Flag**: this option does not have a value ... it'll toggle the default
value
  * `<default or example value>`: is a default value for optional options and
an example vaule for required options. Can be empty strings. For optional
options, if a `nil` example is set, the `options` map won't contain an entry
if this option is not passed in by a user. If set to a value different to nil,
the `options` map will contain the default value if the user does not pass in
the option.

The parse method
----------------

The parse method takes a map as a parameter which can contain the following
entries (all optional):

  * `description`: the description of this program
  * `nolong`: don't show long opt style for any option
  * `longoptEqSep`: use the **=** sign as a long opt delimiter, i.e.
    `--logfile /tmp/log.txt` would be `--logfile=/tmp/log.txt`
  * `debug`: set debug mode to true: this will print out `options`,
    `arguments` and `passThrough` values before returning

