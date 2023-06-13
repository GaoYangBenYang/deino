package commands

import (
	"flag"
	"io"
	"os"
	"strings"

	"github.com/GaoYangBenYang/deino/pkg/utils"
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

//可用命令切片
var AvailableCommands = []*Command{}


var cmdUsage = `Use {{printf "bee help %s" .Name | bold}} for more information.{{endline}}`
// GetName方法返回命令的名称:Usage行中的第一个单词。
func (c *Command) GetName() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}


func (c *Command) Usage() {
	utils.TemplateRendering(cmdUsage, c)
	os.Exit(0)
}