package log

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func InitLogger(production bool) {
	// Ensure that the logger is initialized only once
	once.Do(func() {
		var (
			cfg     zap.Config
			encoder zapcore.Encoder
			level   zapcore.Level
			ws      zapcore.WriteSyncer
		)

		if production {
			// Use production configuration with JSON encoding and readable timestamps
			cfg = zap.NewProductionConfig()
			encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
			cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
			level = zap.InfoLevel
		} else {
			cfg = zap.NewDevelopmentConfig()
			encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
			cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
			level = zap.DebugLevel
		}

		// Send logs to stdout
		ws = zapcore.AddSync(os.Stdout)

		core := zapcore.NewCore(
			encoder, ws, level,
		)

		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
		zap.ReplaceGlobals(logger)

		logger.Info("Logger initialized", zap.Bool("production", production))
	})
}
