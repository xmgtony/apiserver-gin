// Package constant 应用常量包，放置项目所需要的常量
// 建议不同的模块常量分开放，便于维护，也可以避免不同模块开发人员修改同一文件带来冲突
// 目前只有一个项目，并未拆分模块，所以集中放在这里
package constant

const (
	// TraceID 请求id名称
	TraceID = "trace_id"
	// UserID 用户id key
	UserID = "user_id"
)
