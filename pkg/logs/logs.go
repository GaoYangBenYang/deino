package logs

//日志信息级别:
const (
	//OFF 		| 关闭：最高级别，不打印日志。
	LEVEL_OFF = iota
	//INFO 		| 信息：指明描述信息，从粗粒度上描述了应用运行过程。
	LEVEL_INFO
	//DEBUG 	| 调试：指明细致的事件信息，对调试应用最有用。
	LEVEL_DEBUG
	//WARN 		| 警告：指明可能潜在的危险状况。
	LEVEL_WARN
	//ERROR 	| 错误：指明错误事件，但应用可能还能继续运行。
	LEVEL_ERROR
	//FATAL 	| 致命：指明非常严重的可能会导致应用终止执行错误事件。
	LEVEL_FATAL
	//TRACE 	| 跟踪：指明程序运行轨迹，比DEBUG级别的粒度更细。
	LEVEL_TRACE
	//ALL 		| 所有：所有日志级别，包括定制级别。
	LEVEL_ALL
)

type log struct {
}
