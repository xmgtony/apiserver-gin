// author: maxf
// date: 2021-03-05 13:27
// version:

package log

import (
	"apiserver-gin/internal/base/constant"
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/xtime"
	"apiserver-gin/tools/uuid"
	"context"
	"testing"
)

func init() {
	defer Sync()
	c := config.Config{
		LogConfig: config.LogConfig{
			Level:      "debug",
			FileName:   "test.log",
			TimeFormat: xtime.DateTime,
			MaxSize:    1,
			MaxBackups: 5,
			MaxAge:     2,
			Compress:   false,
			LocalTime:  true,
			Console:    true,
		},
		AppName: "zapTest",
	}
	InitLogger(&(c.LogConfig),
		WithOption("requestId", Valuer(func(ctx context.Context) any {
			return ctx.Value(constant.TraceID)
		})))
}

func TestInfo(t *testing.T) {
	Info("test info", "age", 20, "name", "小明")
}

func BenchmarkInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info("test info", "age", 20, "name", "小明")
	}
}

func TestTempLogger_Debug(t *testing.T) {
	c := context.WithValue(context.TODO(), constant.TraceID, uuid.GenUUID16())
	WithCtx(c).Debug("test log Request ID", "age", 20, "name", "小明")
	// 在包外使用时, 可以把web框架比如*gin.Context实例直接传入
	// log.WithCtx(c).Debug("test log Request ID", "age", 20, "name", "小明")
}
