package run

import (
	"fmt"

	"github.com/GaoYangBenYang/deino/cmd/commands"
)

var runCommand = &commands.Command{
	Run:              run,
	Name:             "run",
	ShortDescription: "Deino run",
	FullDescription:  "Deino run 1",
}

func run(cmd *commands.Command, args []string) int {

	fmt.Println("开始运行deino")

	return 0
}
