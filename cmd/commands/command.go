package commands

import (
	"flag"
	"io"
	"os"

	"github.com/GaoYangBenYang/deino/pkg/utils"
)

const CMD_USAGE = `Use {{printf "bee help %s" .Name | bold}} for more information.{{endline}}`

// 已注册命令集
var AvailableCommands = []*Command{}

// command：指令结构
type Command struct {
	// Run：执行命令的操作
	// args：是命令名之后的参数。
	Run func(cmd *Command, args []string) int
	// BeforeRun 执行命令前的操作
	// args是命令名之后的参数。
	BeforeRun func(cmd *Command, args []string)
	// Name 命令的名称
	Name string
	// ShortDescription 是在'deino help'输出中显示的简短描述。
	ShortDescription string
	// FullDescription 是“deino help <this-command>”输出中显示的完整描述。
	FullDescription string
	// Flag 是一组特定于此命令的标志。
	Flag flag.FlagSet
	// CustomFlags 指示该命令将自行执行
	CustomFlags bool
	//如果在SetOutput(w)中设置，则输出写入器
	output *io.Writer
}

func (c *Command) Usage() {
	utils.TemplateRendering(CMD_USAGE, c)
	os.Exit(0)
}