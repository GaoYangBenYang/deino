package utils

import (
	// "fmt"
	// "html/template"
	// "os"
	"os/exec"
	// "strings"
)

// 更新通知
func NoticeUpdate() {
	cmd := exec.Command("go", "version")
	cmd.Output()
}

// 模板渲染
func TemplateRendering(text string, data interface{}) {
	// output := colors.NewColorWriter(os.Stderr)

	// t := template.New("Usage").Funcs(BeeFuncMap())
	// template.Must(t.Parse(text))

	// err := t.Execute(output, data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
}

// // BeeFuncMap returns a FuncMap of functions used in different templates.
// func BeeFuncMap() template.FuncMap {
// 	return template.FuncMap{
// 		"trim":       strings.TrimSpace,
// 		"bold":       colors.Bold,
// 		"headline":   colors.MagentaBold,
// 		"foldername": colors.RedBold,
// 		"endline":    EndLine,
// 		"tmpltostr":  TmplToString,
// 	}
// }
