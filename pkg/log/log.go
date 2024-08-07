// 作为中间层，即使底层log库更换，也不影响业务

package log

import (
	"apiserver-gin/pkg/config"
	"apiserver-gin/pkg/xtime"
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
	cfg      *config.LogConfig
	sugar    *zap.SugaredLogger
	_level   zapcore.Level
	_ctx     context.Context
	prefixes []Option // 公共打印前缀
}

type Valuer func(ctx context.Context) any

type Option struct {
	key string
	val Valuer
}

// WithOption 设置日志打印的公共内容
func WithOption(key string, val any) Option {
	valuer, ok := val.(Valuer)
	if !ok {
		if _, ok := val.(string); !ok {
			// val只能为Valuer类型或者字符串字面量值
			panic("val can only be set to Valuer type or a string literal value")
		}
		valuer = Valuer(func(ctx context.Context) any {
			return val
		})
	}
	return Option{key: key, val: valuer}
}

func WithCtx(ctx context.Context) *logger {
	return &logger{cfg: _logger.cfg, sugar: _logger.sugar, _level: _logger._level, _ctx: ctx, prefixes: _logger.prefixes}
}

// InitLogger 初始化日志配置
func InitLogger(_cfg *config.LogConfig, options ...Option) {
	once.Do(func() {
		_logger = &logger{
			cfg:  _cfg,
			_ctx: context.Background(),
		}
		lumber := _logger.newLumber()
		writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumber))
		sugar := zap.New(_logger.newCore(writeSyncer),
			zap.ErrorOutput(writeSyncer),
			zap.AddCaller(),
			zap.AddCallerSkip(2)).Sugar()

		_logger.sugar = sugar

		if len(options) > 0 {
			_logger.prefixes = options
		}
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
		TimeKey:        "xtime",
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
		format = xtime.DateTimeMs
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

func (l *logger) enabledLevel(level zapcore.Level) bool {
	return level >= l._level
}

func (l *logger) log(level zapcore.Level, message string, kvs []any) {
	if !_logger.enabledLevel(zapcore.DebugLevel) {
		return
	}
	if hasValuers(l) {
		allKvs := make([]any, 0, len(kvs)+len(l.prefixes))
		for _, option := range l.prefixes {
			allKvs = append(allKvs, option.key, option.val(l._ctx))
		}
		kvs = append(allKvs, kvs...)
	}
	switch level {
	case zapcore.DebugLevel:
		_logger.sugar.Debugw(message, kvs...)
	case zapcore.InfoLevel:
		_logger.sugar.Infow(message, kvs...)
	case zapcore.WarnLevel:
		_logger.sugar.Warnw(message, kvs...)
	case zapcore.ErrorLevel:
		_logger.sugar.Errorw(message, kvs...)
	case zapcore.FatalLevel:
		_logger.sugar.Fatalw(message, kvs...)
	default:
		{
		}
	}
}

func hasValuers(l *logger) bool {
	return l.prefixes != nil && len(l.prefixes) > 0
}

// Debug 打印debug级别信息
func (l *logger) Debug(message string, kvs ...any) {
	l.log(zapcore.DebugLevel, message, kvs)
}

// Info 打印info级别信息
func (l *logger) Info(message string, kvs ...any) {
	l.log(zapcore.InfoLevel, message, kvs)
}

// Warn 打印warn级别信息
func (l *logger) Warn(message string, kvs ...any) {
	l.log(zapcore.WarnLevel, message, kvs)
}

// Error 打印error级别信息
func (l *logger) Error(message string, kvs ...any) {
	l.log(zapcore.ErrorLevel, message, kvs)
}

func (l *logger) Fatal(message string, kvs ...any) {
	l.log(zapcore.FatalLevel, message, kvs)
}

// 下面是一些包级别方法，使用默认的_logger。直接使用包名也能用，比如log.Debug()
// 但是当我们使用log.WithCtx(ctx)以后，就不能使用包方法了，所以logger本身也实现了对应级别的日志打印

// Debug 打印debug级别信息
func Debug(message string, kvs ...any) {
	_logger.log(zapcore.DebugLevel, message, kvs)
}

// Info 打印info级别信息
func Info(message string, kvs ...any) {
	_logger.log(zapcore.InfoLevel, message, kvs)
}

// Warn 打印warn级别信息
func Warn(message string, kvs ...any) {
	_logger.log(zapcore.WarnLevel, message, kvs)
}

// Error 打印error级别信息
func Error(message string, kvs ...any) {
	_logger.log(zapcore.ErrorLevel, message, kvs)
}

func Fatal(message string, kvs ...any) {
	_logger.log(zapcore.FatalLevel, message, kvs)
}

// Sync 关闭时需要同步日志到输出
func Sync() {
	if _logger != nil {
		_ = _logger.sugar.Sync()
	}
}
