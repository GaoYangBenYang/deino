package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/GaoYangBenYang/deino/cmd"
	"github.com/GaoYangBenYang/deino/utils"
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

	//没有参数命令提示时提示使用
	if len(args) < 1 {
		cmd.Usage()
		os.Exit(0)
		fmt.Println("执行一次")
		return
	}

	//帮助文档
	if args[0] == "help" {
		cmd.Help(args[1:])
		return
	}

	fmt.Println(args)

	fmt.Println("执行成功！")
}
