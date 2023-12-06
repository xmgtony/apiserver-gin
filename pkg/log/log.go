// 作为中间层，即使底层log库更换，也不影响业务

package log

import (
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/constant"
	"context"
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

// DefaultPair 表示接收打印的键值对参数
type DefaultPair struct {
	key   string
	value interface{}
}

func Pair(key string, v interface{}) DefaultPair {
	return DefaultPair{
		key:   key,
		value: v,
	}
}

func spread(kvs ...DefaultPair) []interface{} {
	s := make([]interface{}, 0, len(kvs))
	for _, v := range kvs {
		s = append(s, v.key, v.value)
	}
	return s
}

// Debug 打印debug级别信息
func Debug(message string, kvs ...DefaultPair) {
	if !_logger.EnabledLevel(zapcore.DebugLevel) {
		return
	}
	args := spread(kvs...)
	_logger.sugar.Debugw(message, args...)
}

// Info 打印info级别信息
func Info(message string, kvs ...DefaultPair) {
	if !_logger.EnabledLevel(zapcore.InfoLevel) {
		return
	}
	args := spread(kvs...)
	_logger.sugar.Infow(message, args...)
}

// Warn 打印warn级别信息
func Warn(message string, kvs ...DefaultPair) {
	if !_logger.EnabledLevel(zapcore.WarnLevel) {
		return
	}
	args := spread(kvs...)
	_logger.sugar.Warnw(message, args...)
}

// Error 打印error级别信息
func Error(message string, kvs ...DefaultPair) {
	if !_logger.EnabledLevel(zapcore.ErrorLevel) {
		return
	}
	args := spread(kvs...)
	_logger.sugar.Errorw(message, args...)
}

// Fatal 打印错误信息，然后退出
func Fatal(message string, kvs ...DefaultPair) {
	if !_logger.EnabledLevel(zapcore.FatalLevel) {
		return
	}
	args := spread(kvs...)
	_logger.sugar.Fatalw(message, args...)
}

// tempLogger 临时的logger，作为链式调用中间变量
type tempLogger struct {
	extra []DefaultPair
}

// getPrefix 根据extra生成日志前缀，比如 "requestId:%s name:%s "
func (tl *tempLogger) getPrefix(template string, args []interface{}) ([]interface{}, string) {

	if len(tl.extra) > 0 {
		var prefix string
		tmp := make([]interface{}, 0, len(args)+len(tl.extra))
		for _, pair := range tl.extra {
			prefix += pair.key + ":%s,"
			tmp = append(tmp, pair.value)
		}
		args = append(tmp, args...)
		template = prefix + template
	}
	return args, template
}

func (tl *tempLogger) getArgs(kvs []DefaultPair) []interface{} {
	var args []interface{}
	if len(tl.extra) > 0 {
		tl.extra = append(tl.extra, kvs...)
		args = spread(tl.extra...)
	} else {
		args = spread(kvs...)
	}
	return args
}

// RID 实现rid(RequestID打印) 使用格式 log.RID(ctx).Debug(), 可以继续拓展 比如Log.RID(ctx).AppName(ctx).Debug()
func RID(ctx context.Context) *tempLogger {
	tl := &tempLogger{extra: make([]DefaultPair, 0)}
	if ctx == nil {
		return tl
	}
	if v := ctx.Value(constant.RequestId); v != nil && v != "" {
		tl.extra = append(tl.extra, Pair(constant.RequestId, v))
	}
	return tl
}

func (tl *tempLogger) Debug(message string, kvs ...DefaultPair) {
	// 这里重复写的原因是zap的log设置的SKIP是1，
	//并且使用的全局只有一个logger，不能修改SKIP，否则打印的位置不正确，后续都是重复代码
	// Debug(message, tl.extra...) 这种写法要修改SKIP
	if !_logger.EnabledLevel(zapcore.DebugLevel) {
		return
	}
	args := tl.getArgs(kvs)
	_logger.sugar.Debugw(message, args...)
}

func (tl *tempLogger) Info(message string, kvs ...DefaultPair) {
	if !_logger.EnabledLevel(zapcore.InfoLevel) {
		return
	}
	args := tl.getArgs(kvs)
	_logger.sugar.Infow(message, args...)
}

// Warn 打印warn级别信息
func (tl *tempLogger) Warn(message string, kvs ...DefaultPair) {
	if !_logger.EnabledLevel(zapcore.WarnLevel) {
		return
	}
	args := tl.getArgs(kvs)
	_logger.sugar.Warnw(message, args...)
}

// Error 打印error级别信息
func (tl *tempLogger) Error(message string, kvs ...DefaultPair) {
	if !_logger.EnabledLevel(zapcore.ErrorLevel) {
		return
	}
	args := tl.getArgs(kvs)
	_logger.sugar.Errorw(message, args...)
}

// Fatal 打印错误信息，然后退出
func (tl *tempLogger) Fatal(message string, kvs ...DefaultPair) {
	if !_logger.EnabledLevel(zapcore.FatalLevel) {
		return
	}
	args := tl.getArgs(kvs)
	_logger.sugar.Fatalw(message, args...)
}

// Sync 关闭时需要同步日志到输出
func Sync() {
	_ = _logger.sugar.Sync()
}
