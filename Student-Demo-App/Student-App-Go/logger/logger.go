package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction() // or zap.NewDevelopment()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}
