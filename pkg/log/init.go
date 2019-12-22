package log

import (
	. "apidemo-gin/conf"
	"apidemo-gin/pkg/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var logger *zap.Logger

func LoggerInit() {
	logger = zap.New(newCore(),
		zap.AddCaller(),
		zap.Fields(zap.String("appName", Cfg.ApplicationName)))
}

func New() *zap.Logger {
	if logger == nil {
		LoggerInit()
	}
	return logger
}

func newCore() zapcore.Core {
	// rolling log
	lumber := &lumberjack.Logger{
		Filename:   Cfg.LogCfg.FileName,
		MaxSize:    Cfg.LogCfg.MaxSize,
		MaxAge:     Cfg.LogCfg.MaxAge,
		MaxBackups: Cfg.LogCfg.MaxBackups,
		LocalTime:  Cfg.LogCfg.LocalTime,
		Compress:   Cfg.LogCfg.Compress,
	}
	// log level
	atomicLevel := zap.NewAtomicLevel()
	defaultLevel := zapcore.DebugLevel
	// 会解码传递的日志级别，生成新的日志级别
	_ = (&defaultLevel).UnmarshalText([]byte(Cfg.LogCfg.Level))
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
	if Cfg.LogCfg.Console {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	} else {
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
	timeLayout := Cfg.LogCfg.TimeFormat
	if len(timeLayout) <= 0 {
		timeLayout = constant.TimeLayoutMs
	}
	enc.AppendString(t.Format(timeLayout))
}

// 包装下面方法，为了减少外面调用的层次(纯粹是懒)
func Debug(msg string, filed ...zap.Field) {
	logger.Debug(msg, filed...)
}

func Info(msg string, filed ...zap.Field) {
	logger.Info(msg, filed...)
}

func Warn(msg string, filed ...zap.Field) {
	logger.Warn(msg, filed...)
}

func Error(msg string, filed ...zap.Field) {
	logger.Error(msg, filed...)
}

func Fatal(msg string, filed ...zap.Field) {
	logger.Fatal(msg, filed...)
}

func Sync() {
	_ = logger.Sync()
}
