package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/indiealistic/bethacon/grpcapi/microservice"
	"github.com/indiealistic/bethacon/utils/logex"
)

const (
	serviceName = "BethaCon"
)

// Version may be changed during build via --ldflags parameter
var Version = "latest"

func main() {
	logger := logex.Build(serviceName, zapcore.InfoLevel)

	// Initialize service.
	microService, err := microservice.Init(&microservice.ClientOptions{
		Name:    serviceName,
		Version: Version,
		Log:     logger,
	})
	if err != nil {
		logger.Fatal("failed to initialize micro-service", zap.Error(err))
	}

	// Run service.
	if err := microService.Run(); err != nil {
		logger.Fatal("failed to run micro-service", zap.Error(err))
	}
}
