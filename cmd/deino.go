package cmd

import (
	"fmt"
	"log"
)

const USAGE_TEMPLATE = `Bee is a Fast and Flexible tool for managing your Beego Web Application.

You are using bee for beego v2.x. If you are working on beego v1.x, please downgrade version to bee v1.12.0

{{"USAGE" | headline}}
    {{"bee command [arguments]" | bold}}

{{"AVAILABLE COMMANDS" | headline}}
{{range .}}{{if .Runnable}}
    {{.Name | printf "%-11s" | bold}} {{.Short}}{{end}}{{end}}

Use {{"bee help [command]" | bold}} for more information about a command.

{{"ADDITIONAL HELP TOPICS" | headline}}
{{range .}}{{if not .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use {{"bee help [topic]" | bold}} for more information about that topic.
`

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

const ERROR_TEMPLATE = `
bee: %s.Use {{"bee help" | bold}} for more information.
`

func Usage() {
	fmt.Printf(USAGE_TEMPLATE)
}

func Help(args []string) {
	if len(args) == 0 {
		Usage()
		return
	}
	if len(args) != 1 {
		//设置
		log.SetFlags(log.Lmicroseconds | log.Ldate)
		log.Println("Too many arguments")
	}
	// USAGE
	// 	bee command [arguments]

	// COMMANDS

	//     version     Prints the current Bee version
	// 	update      Update Bee
	//     new         Creates a Beego application
	//     run         Run the application by starting a local development server

	// Use deino help [command] for more information about a command.

}
