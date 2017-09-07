package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// CliCommands is the list of all Cli commands available.
// Each command `cmd` is invoked via `Cli cmd`
var CliCommands []*Cmd

// Cmd is used by any package exposing a runnable command to gather information
// about command name, usage and flagset.
type Cmd struct {
	// Name is the name of the command. It's what comes after `Cli`.
	Name string

	// Run is the command entrypoint.
	Run func(...string)

	// UsageLine is the header of what's printed by flag.PrintDefaults.
	UsageLine string

	// Short is a one-line description of what the command does.
	Short string

	// Long is the detailed description of what the command does.
	Long string

	// Flag is the set of flags accepted by the command. This should be initialized in
	// the command's module's `init` function. The parsing of these flags is issued
	// by the main Cli entrypoint, so each command doesn't have to do it itself.
	Flag flag.FlagSet
}

// AddCommand allows packages to setup their own command. In order for them to
// be compiled, they must be imported by the main package with the "_" alias
func AddCommand(c *Cmd) {
	CliCommands = append(CliCommands, c)
}

func init() {
	AddCommand(cmdVersion)
}

func (c *Cmd) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	c.Flag.PrintDefaults()
	os.Exit(2)
}

// FindCommand takes a string and searches for a command whose name has that string as
// prefix. If more than 1 command name has that string as a prefix (and no command name
// equals that string), an error is returned. If no suitable command is found, an error
// is returned.
func FindCommand(name string) (command *Cmd, err error) {
	for _, cmd := range CliCommands {
		if cmd.Name == name {
			command = cmd
			// If there were several commands beginning with this string, but I
			// have an exact match, the error should not be returned.
			err = nil
			return
		}
		if strings.HasPrefix(cmd.Name, name) {
			if command != nil {
				err = fmt.Errorf("Ambiguous command: '%s'.", name)
			} else {
				command = cmd
			}
		}
	}
	if command == nil {
		err = fmt.Errorf("Command not found: '%s'.", name)
	}
	return
}

var defaultCommand *Cmd

func SetDefault(c *Cmd) {
	defaultCommand = c
}

func initCommands() {
	if len(os.Args) > 1 {
		//read the first argument
		directive := os.Args[1]
		if len(os.Args) > 2 {
			//shift parameters left, but keep argv[0]
			os.Args = append(os.Args[:1], os.Args[2:]...)
		} else {
			os.Args = os.Args[:1]
		}
		command, err := FindCommand(directive)
		if err == nil {
			callCommand(command)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			fmt.Fprintln(os.Stderr, "Available commands are:\n")
			for _, cmd := range CliCommands {
				fmt.Fprintln(os.Stderr, "\t"+cmd.Name+"\n\t\t"+cmd.Short)
			}
			fmt.Fprintln(os.Stderr, "\nDefault command is: ", defaultCommand.Name)
		}
	} else {
		callCommand(defaultCommand)
	}
}

func callCommand(command *Cmd) {
	command.Flag.Usage = command.Usage
	//TODO handle this error
	_ = command.Flag.Parse(os.Args[1:])
	command.Run(command.Flag.Args()...)
	return
}

var cmdVersion = &Cmd{
	Name: "version",
	Run: func(_ ...string) {
		fmt.Printf("Version: %s\nCommit: %s\nBuild%s\n", Version, Commit, Build)
	},
	UsageLine: "version",
	Short:     "print version and exit",
	Long:      "print version and exit",
}
