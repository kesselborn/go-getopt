package main

import (
  "fmt"
  getopt "github.com/kesselborn/go-getopt"
  "os"
)

func main() {
  sco := getopt.SubCommandOptions{
    getopt.Options{
      "global description",
      getopt.Definitions{
        {"foo|f", "some arg", getopt.Optional, ""},
        {"command", "command to execute", getopt.IsSubcommand, ""},
      },
    },
    getopt.SubCommands{
      "getenv": {
        "getenv description",
        getopt.Definitions{
          {"bar|b", "some arg", getopt.Optional, ""},
          {"name", "app's name", getopt.IsArg | getopt.Required, ""},
          {"key", "environment variable's name", getopt.IsArg | getopt.Required, ""},
        },
      },
      "register": {
        "register description",
        getopt.Definitions{
          {"deploytype|t", "deploy type (one of mount, bazapta, lxc)", getopt.NoLongOpt | getopt.Optional | getopt.ExampleIsDefault, "lxc"},
          {"name|n", "app's name", getopt.IsArg | getopt.Required, ""},
        },
      },
    },
  }

  scope, options, arguments, passThrough, e := sco.ParseCommandLine()

  if e != nil {
    exit_code := 0

    switch {
    case e.ErrorCode == getopt.WantsUsage:
      fmt.Print(sco.Usage(scope))
    case e.ErrorCode == getopt.WantsHelp:
      fmt.Print(sco.Help(scope))
    default:
      fmt.Println(sco.Help(scope), "\n", "**** Error: ", e.Message, "\n")
      exit_code = e.ErrorCode
    }
    os.Exit(exit_code)
  }

  fmt.Printf("scope:\n%s", scope)
  fmt.Printf("options:\n%#v", options)
  fmt.Printf("arguments: %#v\n", arguments)
  fmt.Printf("passThrough: %#v\n", passThrough)
}
