package log

import (
	"fmt"
	"io"
	"net/url"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	TestCustomIO = "CustomOut"
	customSuffix = "anyone"
)

type customWriter struct {
	io.Writer
}

func (cw customWriter) Close() error {
	// modify as preferred, for testing purpose this is enough:
	return nil
}
func (cw customWriter) Sync() error {
	// modify as preferred, for testing purpose this is enough:
	return nil
}

/*
Attn!
 1. if use stdOut, can leave outpathO blank: `""` (and leave `buf` nil)
 2. for any case of io.Writer output just set outpathO to TestCustomIO accordingly.
    Of cource, you can change suffix above for the scheme as you like.
*/
func StructuredInit(y string, outpathO string, buf io.Writer) *zap.Logger {
	if outpathO == TestCustomIO {
		err := zap.RegisterSink(outpathO, func(u *url.URL) (zap.Sink, error) {
			return customWriter{buf}, nil
		})
		if err != nil {
			panic(err)
		}
		outpathO = fmt.Sprintf("%s:%s", outpathO, customSuffix)
	} else {
		outpathO = "stdout"
	}

	var logCfg = zap.Config{
		Encoding:         "json",
		OutputPaths:      []string{outpathO},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,

			FunctionKey: "function",

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	level := zapcore.InfoLevel
	switch y {
	case "DEBUG":
		level = zapcore.DebugLevel
	case "WARN":
		level = zapcore.WarnLevel
	case "ERROR":
		level = zapcore.ErrorLevel
	case "PANIC":
		level = zapcore.PanicLevel
	case "DPANIC":
		level = zapcore.DPanicLevel
	case "FATAL":
		level = zapcore.FatalLevel
	default:
	}
	logCfg.Level = zap.NewAtomicLevelAt(level)
	logCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zap.Must(logCfg.Build())
}

// all downside are for OTLADKA ONLY:
type ProjectLoggerD struct {
	*zap.Logger
}

func (l ProjectLoggerD) Printf(s string, i ...interface{}) {
	var allm []zapcore.Field
	for _, m := range i {
		allm = append(allm, zapcore.Field{Type: zapcore.ReflectType, Interface: m})
	}
	l.Debug(s, allm...)
}

type ProjectLoggerI struct {
	*zap.Logger
}

func (l ProjectLoggerI) Printf(s string, i ...interface{}) {
	var allm []zapcore.Field
	for _, m := range i {
		allm = append(allm, zapcore.Field{Type: zapcore.ReflectType, Interface: m})
	}
	l.Info(s, allm...)
}

type ProjectLoggerE struct {
	*zap.Logger
}

func (l ProjectLoggerE) Printf(s string, i ...interface{}) {
	var allm []zapcore.Field
	for _, m := range i {
		allm = append(allm, zapcore.Field{Type: zapcore.ReflectType, Interface: m})
	}
	l.Error(s, allm...)
}
