package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

// MayaCtlName will be fiiled in by the GNUMakefile using an EVN variable
//  MAYACTL
var MayaCtlName string = "maya"

func main() {
	os.Exit(Run(os.Args[1:]))
}

func Run(args []string) int {
	return RunCustom(args, Commands(nil))
}

func RunCustom(args []string, commands map[string]cli.CommandFactory) int {
	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	// Extract the commands to include in the help
	commandNames := make([]string, 0, len(commands))
	for k, _ := range commands {
		commandNames = append(commandNames, k)
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: commands,
		HelpFunc: cli.FilteredHelpFunc(commandNames, cli.BasicHelpFunc(MayaCtlName)),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
