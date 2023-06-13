package cmd

import (
	"fmt"
	"os"

	"github.com/GaoYangBenYang/deino/cmd/commands"
	"github.com/GaoYangBenYang/deino/pkg/utils"
)

// 使用方法渲染模板
const USAGE_TEMPLATE = `
{{"USAGE" | headline}}
    {{"deino command [arguments]" | bold}}

{{"AVAILABLE COMMANDS" | headline}}
	{{range .}}{{if .Runnable}}
		{{.Name | printf "%-11s" | bold}} {{.Short}}{{end}}{{end}}

Use {{"deino help [command]" | bold}} for more information about a command.
`

// 帮助文档模板渲染
const HELP_TEMPLATE = `
{{"USAGE" | headline}}
{{.UsageLine | printf "bee %s" | bold}}
{{if .Options}}{{endline}}{{"OPTIONS" | headline}}{{range $k,$v := .Options}}
{{$k | printf "-%s" | bold}}
{{$v}}
{{end}}{{end}}
{{"DESCRIPTION" | headline}}
{{tmpltostr .Long . | trim}}
`

// 错误模板渲染
const ERROR_TEMPLATE = `
deino: %s.

Use {{"deino help" | bold}} for more information.
`

func Usage() {
	//模板渲染
	utils.TemplateRendering(USAGE_TEMPLATE, commands.AvailableCommands)
}

func Help(args []string) {
	//除了help命令外无其他参数,进行使用方法渲染
	if len(args) == 0 {
		Usage()
		os.Exit(0)
	}
	//子命令过多打印错误并退出
	if len(args) != 1 {
		utils.PrintErrorAndExit("Too many arguments", ERROR_TEMPLATE)
		return
	}
	//子命令文档渲染
	for _, cmd := range commands.AvailableCommands {
		fmt.Println(args[0],cmd)
		// if cmd.GetName() == args[0] {
		// 	// utils.TemplateRendering(HELP_TEMPLATE, cmd)
		// 	fmt.Println("子帮助文档")
		// 	return
		// }
	}
	//打印错误并退出
	utils.PrintErrorAndExit("Unknown help topic", ERROR_TEMPLATE)
}
