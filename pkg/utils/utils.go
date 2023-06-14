package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"

	"github.com/GaoYangBenYang/deino/pkg/colors"
)

// 更新通知
func NoticeUpdate() {
	cmd := exec.Command("go", "version")
	cmd.Output()
}

// 模板渲染
func TemplateRendering(text string, data interface{}) {
	// output := colors.NewColorWriter(os.Stderr)
	//New方法创建一个Usage模板
	//Funcs方法向usageTemplate的函数字典里加入参数funcMap内的键值对。如果funcMap某个键值对的值不是函数类型或者返回值不符合要求会panic。但是，可以对usageTemplate函数列表的成员进行重写。方法返回t以便进行链式调用。
	usageTemplate := template.New("Usage").Funcs(FuncMap())
	//Must函数用于包装返回(*Template, error)的函数/方法调用，它会在err非nil时panic，一般用于变量初始化：
	template.Must(usageTemplate.Parse(text))
	// Execute方法将解析好的模板应用到data上，并将输出写入wr。如果执行时出现错误，会停止执行，但有可能已经写入wr部分数据。模板可以安全的并发执行
	err := usageTemplate.Execute(os.Stderr, data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// FuncMap 返回不同模板中使用的函数映射
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"trim":       strings.TrimSpace,
		"bold":       colors.Bold,
		"headline":   colors.MagentaBold,
		"foldername": colors.RedBold,
		"endline":    EndLine,
		"tmpltostr":  TemplateToString,
	}
}

// TemplateToString 解析文本模板并将结果作为字符串返回。
func TemplateToString(tmpl string, data interface{}) string {
	t := template.New("tmpl").Funcs(FuncMap())
	template.Must(t.Parse(tmpl))

	var doc bytes.Buffer
	err := t.Execute(&doc, data)
	if err != nil {
		panic(err)
	}

	return doc.String()
}

// EndLine 返回换行转义字符
func EndLine() string {
	return "\n"
}

// 打印错误信息并退出
func PrintErrorAndExit(message, errorTemplate string) {
	TemplateRendering(fmt.Sprintf(errorTemplate, message), nil)
}
