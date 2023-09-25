package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GaoYangBenYang/deino/cmd"
	"github.com/GaoYangBenYang/deino/cmd/commands"
	"github.com/GaoYangBenYang/deino/pkg/utils"
)

func main() {
	//检查deino是否有版本更新
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
	//执行已经注册的命令命令
	for _, cmd := range commands.AvailableCommands {
		//判断命令名称是否与输入匹配以及是否绑定命令函数
		if cmd.Name == args[0] && cmd.Run != nil {
			fmt.Println("111")
			fmt.Println(cmd.Name, cmd.Run)
			cmd.Flag.Usage = func() { cmd.Usage() }
			//是否自动执行函数
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			//在执行函数之前的相关操作
			if cmd.BeforeRun != nil {
				cmd.BeforeRun(cmd, args)
			}

			// config.LoadConfig()

			//结束
			os.Exit(cmd.Run(cmd, args))
		}
	}
	//未知命令
	utils.PrintErrorAndExit("Unknown subcommand", cmd.ERROR_TEMPLATE)
}
