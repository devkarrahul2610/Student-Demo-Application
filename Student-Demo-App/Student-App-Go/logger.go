package main

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func initLogger() {
	var err error
	Logger, err = zap.NewProduction() // or zap.NewDevelopment()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}
