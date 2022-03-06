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

var (
	_logger *logger
	once    sync.Once
)

type logger struct {
	cfg    *config.LogConfig
	sugar  *zap.SugaredLogger
	_level zapcore.Level
}

// InitLogger 初始化日志配置
func InitLogger(_cfg *config.LogConfig, appName string) {
	once.Do(func() {
		_logger = &logger{
			cfg: _cfg,
		}
		lumber := _logger.newLumber()
		writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumber))
		sugar := zap.New(_logger.newCore(writeSyncer),
			zap.ErrorOutput(writeSyncer),
			zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.Fields(zap.String("appName", appName))).
			Sugar()

		_logger.sugar = sugar
	})
}

func (l *logger) newCore(ws zapcore.WriteSyncer) zapcore.Core {
	// 默认日志级别
	atomicLevel := zap.NewAtomicLevel()
	defaultLevel := zapcore.DebugLevel
	// 会解码传递的日志级别，生成新的日志级别
	_ = (&defaultLevel).UnmarshalText([]byte(l.cfg.Level))
	atomicLevel.SetLevel(defaultLevel)
	l._level = defaultLevel

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
		EncodeTime:     l.customTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	var writeSyncer zapcore.WriteSyncer
	if l.cfg.Console {
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
func (l *logger) customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	format := l.cfg.TimeFormat
	if len(format) <= 0 {
		format = constant.TimeLayoutMs
	}
	enc.AppendString(t.Format(format))
}

func (l *logger) newLumber() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   l.cfg.FileName,
		MaxSize:    l.cfg.MaxSize,
		MaxAge:     l.cfg.MaxAge,
		MaxBackups: l.cfg.MaxBackups,
		LocalTime:  l.cfg.LocalTime,
		Compress:   l.cfg.Compress,
	}
}

func (l *logger) EnabledLevel(level zapcore.Level) bool {
	return level >= l._level
}

// Pair 表示接收打印的键值对参数
type Pair struct {
	key   string
	value interface{}
}

func WithPair(key string, v interface{}) Pair {
	return Pair{
		key:   key,
		value: v,
	}
}

func spread(kvs ...Pair) []interface{} {
	s := make([]interface{}, 0, len(kvs))
	for _, v := range kvs {
		s = append(s, v.key, v.value)
	}
	return s
}

// Debug 打印debug级别信息
func Debug(message string, kvs ...Pair) {
	if !_logger.EnabledLevel(zapcore.DebugLevel) {
		return
	}
	args := spread(kvs...)
	_logger.sugar.Debugw(message, args...)
}

// Info 打印info级别信息
func Info(message string, kvs ...Pair) {
	if !_logger.EnabledLevel(zapcore.InfoLevel) {
		return
	}
	args := spread(kvs...)
	_logger.sugar.Infow(message, args...)
}

// Warn 打印warn级别信息
func Warn(message string, kvs ...Pair) {
	if !_logger.EnabledLevel(zapcore.WarnLevel) {
		return
	}
	args := spread(kvs...)
	_logger.sugar.Warnw(message, args...)
}

// Error 打印error级别信息
func Error(message string, kvs ...Pair) {
	if !_logger.EnabledLevel(zapcore.ErrorLevel) {
		return
	}
	args := spread(kvs...)
	_logger.sugar.Errorw(message, args...)
}

// Panic 打印错误信息，然后panic
func Panic(message string, kvs ...Pair) {
	if !_logger.EnabledLevel(zapcore.PanicLevel) {
		return
	}
	args := spread(kvs...)
	_logger.sugar.Panicw(message, args...)
}

// Fatal 打印错误信息，然后退出
func Fatal(message string, kvs ...Pair) {
	if !_logger.EnabledLevel(zapcore.FatalLevel) {
		return
	}
	args := spread(kvs...)
	_logger.sugar.Fatalw(message, args...)
}

// Debugf 格式化输出debug级别日志
func Debugf(template string, args ...interface{}) {
	_logger.sugar.Debugf(template, args...)
}

// Infof 格式化输出info级别日志
func Infof(template string, args ...interface{}) {
	_logger.sugar.Infof(template, args...)
}

// Warnf 格式化输出warn级别日志
func Warnf(template string, args ...interface{}) {
	_logger.sugar.Warnf(template, args...)
}

// Errorf 格式化输出error级别日志
func Errorf(template string, args ...interface{}) {
	_logger.sugar.Errorf(template, args...)
}

// Panicf 格式化输出日志，并panic
func Panicf(template string, args ...interface{}) {
	_logger.sugar.Panicf(template, args...)
}

// Fatalf 格式化输出日志，并退出
func Fatalf(template string, args ...interface{}) {
	_logger.sugar.Fatalf(template, args...)
}

// Sync 关闭时需要同步日志到输出
func Sync() {
	_ = _logger.sugar.Sync()
}
