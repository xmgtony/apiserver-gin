package log

import (
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/constant"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger
var once sync.Once

// LoggerInit 初始化日志
func LoggerInit(config config.Config) {
	once.Do(func() {
		lumber := newLumber(config)
		writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumber))
		logger = zap.New(newCore(writeSyncer),
			zap.ErrorOutput(writeSyncer),
			zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.Fields(zap.String("appName", config.ApplicationName))).
			Sugar()
	})
}

func newCore(ws zapcore.WriteSyncer) zapcore.Core {
	// 默认日志级别
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
		writeSyncer = ws
	}
	// Tips: 如果使用zapcore.NewJSONEncoder
	// encoderConfig里面就不要配置 EncodeLevel 为zapcore.CapitalColorLevelEncoder或者是
	// zapcore.LowercaseColorLevelEncoder, 不但日志级别字段不会出现颜色，而且日志级别level字段
	// 会出现乱码，因为控制颜色的字符也被JSON编码了。
	return zapcore.NewCore(zapcore.NewConsoleEncoder(encoder),
		writeSyncer,
		atomicLevel)
}

// CustomTimeEncoder 实现了 zapcore.TimeEncoder
// 实现对日期格式的自定义转换
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	timeLayout := config.Cfg.LogCfg.TimeFormat
	if len(timeLayout) <= 0 {
		timeLayout = constant.TimeLayoutMs
	}
	enc.AppendString(t.Format(timeLayout))
}

func newLumber(config config.Config) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   config.LogCfg.FileName,
		MaxSize:    config.LogCfg.MaxSize,
		MaxAge:     config.LogCfg.MaxAge,
		MaxBackups: config.LogCfg.MaxBackups,
		LocalTime:  config.LogCfg.LocalTime,
		Compress:   config.LogCfg.Compress,
	}
}

// KVPair 表示接收打印的键值对参数
type KVPair struct {
	K string
	V interface{}
}

func spread(kvs ...KVPair) []interface{} {
	s := make([]interface{}, 0, len(kvs))
	for _, v := range kvs {
		s = append(s, v.K, v.V)
	}
	return s
}

// Debug 打印debug级别信息
func Debug(message string, kvs ...KVPair) {
	args := spread(kvs...)
	logger.Debugw(message, args)
}

// Info 打印info级别信息
func Info(message string, kvs ...KVPair) {
	args := spread(kvs...)
	logger.Infow(message, args...)
}

// Warn 打印warn级别信息
func Warn(message string, kvs ...KVPair) {
	args := spread(kvs...)
	logger.Warnw(message, args)
}

// Error 打印error级别信息
func Error(message string, kvs ...KVPair) {
	args := spread(kvs...)
	logger.Errorw(message, args)
}

// Panic 打印错误信息，然后panic
func Panic(message string, kvs ...KVPair) {
	args := spread(kvs...)
	logger.Panicw(message, args)
}

// Fatal 打印错误信息，然后退出
func Fatal(message string, kvs ...KVPair) {
	args := spread(kvs...)
	logger.Fatalw(message, args)
}

// Debugf 格式化输出debug级别日志
func Debugf(template string, agrs ...interface{}) {
	logger.Debugf(template, agrs)
}

// Infof 格式化输出info级别日志
func Infof(template string, agrs ...interface{}) {
	logger.Infof(template, agrs)
}

// Warnf 格式化输出warn级别日志
func Warnf(template string, agrs ...interface{}) {
	logger.Warnf(template, agrs)
}

// Errorf 格式化输出error级别日志
func Errorf(template string, agrs ...interface{}) {
	logger.Errorf(template, agrs)
}

// Panicf 格式化输出日志，并panic
func Panicf(template string, agrs ...interface{}) {
	logger.Panicf(template, agrs)
}

// Fatalf 格式化输出日志，并退出
func Fatalf(template string, agrs ...interface{}) {
	logger.Fatalf(template, agrs)
}

// Sync 关闭时需要同步日志到输出
func Sync() {
	_ = logger.Sync()
}
