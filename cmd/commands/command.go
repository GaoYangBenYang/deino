package commands

import (
	"flag"
	"io"
)

// command结构体
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string) int

	// PreRun performs an operation before running the command
	PreRun func(cmd *Command, args []string)

	// UsageLine is the one-line Usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'go help' output.
	Short string

	// Long is the long message shown in the 'go help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet

	// CustomFlags indicates that the command will do its own
	// flag parsing.
	CustomFlags bool

	// output out writer if set in SetOutput(w)
	output *io.Writer
}
