package logger

import "go.uber.org/zap"

var L *zap.Logger

func InitLogger() {
	L, _ = zap.NewProduction()
}
