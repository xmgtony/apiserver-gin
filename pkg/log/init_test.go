// author: maxf
// date: 2021-03-05 13:27
// version:

package log

import (
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/constant"
	"testing"
)

func init() {
	defer Sync()
	c := config.Config{
		LogCfg: config.LogCfg{
			Level:      "debug",
			FileName:   "test.log",
			TimeFormat: constant.TimeLayout,
			MaxSize:    1,
			MaxBackups: 5,
			MaxAge:     2,
			Compress:   false,
			LocalTime:  true,
			Console:    false,
		},
		ApplicationName: "zapTest",
	}
	LoggerInit(c)
}

func TestInfo(t *testing.T) {
	Info("test info", KVPair{
		k: "age",
		v: 20,
	}, KVPair{
		k: "name",
		v: "小明",
	})
}

func BenchmarkInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info("test info", KVPair{
			k: "age",
			v: 20,
		}, KVPair{
			k: "name",
			v: "小明",
		})
	}
}
