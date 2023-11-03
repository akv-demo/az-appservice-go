package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log - global zap.Logged instance
var Log *zap.Logger

// Sugar - global zap.SugaredLogger instance
var Sugar *zap.SugaredLogger
var atomLevel zap.AtomicLevel

func init() {
	atomLevel = zap.NewAtomicLevel()
	cfg := zap.Config{
		Level:             atomLevel,
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "lvl",
			TimeKey:       "tm ",
			NameKey:       "nm ",
			CallerKey:     "cll",
			StacktraceKey: "stk",
			//LineEnding:     "",
			// EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.NanosDurationEncoder,
			// EncodeCaller:   zapcore.FullCallerEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	var err error
	Log, err = cfg.Build()
	if err != nil {
		panic(err)
	}
	Sugar = Log.Sugar()
}

func Setup(prod bool) {
	if prod {
		Sugar.Info("Use zap logger in prod mode: ", prod)
		atomLevel.SetLevel(zap.InfoLevel)
		Sugar.Debug("This debug message must not be visible")
	} else {
		Sugar.Info("Use zap logger in development mode")
		Sugar.Debug("This debug message must be visible")
		atomLevel.SetLevel(zap.DebugLevel)
	}
}
