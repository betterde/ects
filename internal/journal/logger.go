package journal

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Environment string

var (
	Logger      *zap.SugaredLogger
	atomicLevel zap.AtomicLevel
)

func InitLogger() {
	atomicLevel = zap.NewAtomicLevel()

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		NameKey:        "logger",
		LevelKey:       "level",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		atomicLevel,
	), zap.WithCaller(true))

	Logger = logger.Sugar().Named("focusly")
}

func SetLevel(lvl string) error {
	level, err := zapcore.ParseLevel(lvl)
	if err != nil {
		return err
	}

	atomicLevel.SetLevel(level)

	return nil
}
