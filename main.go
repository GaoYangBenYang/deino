package main

import (
	"flag"
	"os"

	"github.com/GaoYangBenYang/dego/cmd"
	// "github.com/GaoYangBenYang/dego/cmd/commands"
	"github.com/GaoYangBenYang/dego/pkg/utils"
	// "github.com/GaoYangBenYang/dego/config"
)

func main() {
	//检查dego是否有版本更新
	utils.NoticeUpdate()
	//Usage打印到标准错误输出一个使用信息，记录了所有注册的flag。本函数是一个包变量，可以将其修改为指向自定义的函数。
	flag.Usage = cmd.Usage
	//从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()
	//返回解析之后剩下的非flag参数。（不包括命令名）
	args := flag.Args()
	//没有命令参数时用法提示
	if len(args) < 1 {
		cmd.Usage()
		os.Exit(0)
	}
	//帮助文档
	if args[0] == "help" {
		cmd.Help(args[1:])
		os.Exit(0)
	}
	if args[0] == "version" { 
		
		os.Exit(0)
	}
	//执行命令
	// for _, cmd := range commands.AvailableCommands {
	// 	if cmd.GetName() == args[0] && cmd.Run != nil {
	// 		cmd.Flag.Usage = func() { cmd.Usage() }
	// 		if cmd.CustomFlags {
	// 			args = args[1:]
	// 		} else {
	// 			cmd.Flag.Parse(args[1:])
	// 			args = cmd.Flag.Args()
	// 		}

	// 		if cmd.PreRun != nil {
	// 			cmd.PreRun(cmd, args)
	// 		}

	// 		config.LoadConfig()
	// 		os.Exit(cmd.Run(cmd, args))
	// 		return
	// 	}
	// }
	//未知命令
	utils.PrintErrorAndExit("Unknown subcommand", cmd.ERROR_TEMPLATE)
}
