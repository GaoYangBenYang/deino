package version

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"time"

	"github.com/GaoYangBenYang/deino/cmd/commands"
	"github.com/GaoYangBenYang/deino/config"
)

//完整版本介绍
const DEINO_VERSION string = ` ____         _               
|  _ \   ___ (_) _ __    ___  
| | | | / _ \| || '_ \  / _ \ 
| |_| ||  __/| || | | || (_) |
|____/  \___||_||_| |_| \___/  v{{ .DeinoVersion }}%s

├── ReactVersion      : {{ ".ReactVersion" }}
├── TypeScriptVersion : {{ ".TypeScriptVersion" }}
├── GoVersion         : {{ .GoVersion }}
├── GOOS              : {{ .GOOS }}
├── GOARCH            : {{ .GOARCH }}
├── NumCPU            : {{ .NumCPU }}
├── GOPATH            : {{ .GOPATH }}
├── GOROOT            : {{ .GOROOT }}
├── Compiler          : {{ .Compiler }}
└── Date              : {{ Now "2023-6-13" }}%s
`

type Version struct {
	GoVersion  string
	GOOS       string
	GOARCH     string
	NumCPU     int
	GOPATH     string
	GOROOT     string
	Compiler   string
	DeinoVersion string
	// Published  string
}

var versionCommand = &commands.Command{
	Run: versionInformation,
	UsageLine: "version",
	Short:     "Prints the current Deino version",
	Long: `Prints the current Bee, Beego and Go version alongside the platform information.`,
}

var outputFormat string

//初始化version命令
func init() {
	//NewFlagSet创建一个新的、名为name，采用errorHandling为错误处理策略的FlagSet。
	//FlagSet代表一个已注册的flag的集合。FlagSet零值没有名字，采用ContinueOnError错误处理策略
	versionFlagSet := flag.NewFlagSet("version",flag.ContinueOnError)
	//StringVar用指定的名称、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量。
	versionFlagSet.StringVar(&outputFormat, "version", "", "Set the output format to json")
	versionCommand.Flag = *versionFlagSet
	commands.AvailableCommands = append(commands.AvailableCommands, versionCommand)
}

//版本信息处理
func versionInformation(cmd *commands.Command, args []string) int {

	cmd.Flag.Parse(args)
	// stdout := cmd.Out()

	if outputFormat != "" {
		version := Version{
			GoVersion:  runtime.Version(),
			GOOS:       runtime.GOOS,
			GOARCH:     runtime.GOARCH,
			NumCPU:     runtime.NumCPU(),
			GOPATH:     os.Getenv("GOPATH"),
			GOROOT:     runtime.GOROOT(),
			Compiler:   runtime.Compiler,
			DeinoVersion: config.VERSION,
		}
		b, err := json.MarshalIndent(version, "", "    ")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(b))
		return 0
	}

	deinoVersion := fmt.Sprintf(DEINO_VERSION, "\x1b[35m", "\x1b[1m","\x1b[0m", "\x1b[32m", "\x1b[1m", "\x1b[0m")
	InitBanner(os.Stderr, bytes.NewBufferString(deinoVersion))
	return 0
}

// InitBanner loads the banner and prints it to output
// All errors are ignored, the application will not
// print the banner in case of error.
func InitBanner(out io.Writer, in io.Reader) {
	if in == nil {
		fmt.Println("The input is nil")
	}

	banner, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Println("Error while trying to read the banner: ", err)
	}

	show(out, string(banner))
}

func show(out io.Writer, content string) {
	t, err := template.New("banner").
		Funcs(template.FuncMap{"Now": Now}).
		Parse(content)

	if err != nil {
		fmt.Println("Cannot parse the banner template:", err)
	}

	err = t.Execute(out, Version{
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.NumCPU(),
		os.Getenv("GOPATH"),
		runtime.GOROOT(),
		runtime.Compiler,
		config.VERSION,
		// utils.GetLastPublishedTime(),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

//现在返回指定格式中的当前本地时间
func Now(layout string) string {
	return time.Now().Format(layout)
}
