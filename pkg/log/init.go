package log

import (
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Log *zap.Logger

func LoggerInit() {
	Log = zap.New(newCore(),
		zap.AddCaller(),
		zap.Fields(zap.String("appName", config.Cfg.ApplicationName)))
}

func New() *zap.Logger {
	if Log == nil {
		LoggerInit()
	}
	return Log
}

func newCore() zapcore.Core {
	// rolling log
	lumber := &lumberjack.Logger{
		Filename:   config.Cfg.LogCfg.FileName,
		MaxSize:    config.Cfg.LogCfg.MaxSize,
		MaxAge:     config.Cfg.LogCfg.MaxAge,
		MaxBackups: config.Cfg.LogCfg.MaxBackups,
		LocalTime:  config.Cfg.LogCfg.LocalTime,
		Compress:   config.Cfg.LogCfg.Compress,
	}
	// log level
	atomicLevel := zap.NewAtomicLevel()
	defaultLevel := zapcore.DebugLevel
	// 会解码传递的日志级别，生成新的日志级别
	_ = (&defaultLevel).UnmarshalText([]byte(config.Cfg.LogCfg.Level))
	atomicLevel.SetLevel(defaultLevel)
	// encoder 这部分没有放到配置文件，因为一般配置一次就不会改动
	encoder := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	var writeSyncer zapcore.WriteSyncer
	if config.Cfg.LogCfg.Console {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	} else {
		// 输出到文件时，不使用彩色日志，否则会出现乱码
		encoder.EncodeLevel = zapcore.CapitalLevelEncoder
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumber))
	}
	// Tips: 如果使用zapcore.NewJSONEncoder
	// encoderConfig里面就不要配置 EncodeLevel 为zapcore.CapitalColorLevelEncoder或者是
	// zapcore.LowercaseColorLevelEncoder, 不但日志级别字段不会出现颜色，而且日志级别level字段
	// 会出现乱码，因为控制颜色的字符也被JSON编码了。
	return zapcore.NewCore(zapcore.NewConsoleEncoder(encoder),
		writeSyncer,
		atomicLevel)
}

// CustomTimeEncoder  implemented zapcore.TimeEncoder
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	timeLayout := config.Cfg.LogCfg.TimeFormat
	if len(timeLayout) <= 0 {
		timeLayout = constant.TimeLayoutMs
	}
	enc.AppendString(t.Format(timeLayout))
}
