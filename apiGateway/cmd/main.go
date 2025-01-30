package main

import (
	"app"
	"app/internal/handler"
	"app/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := app.InitServer()

	handler := handler.NewHandler()

	go func() {
		logger.Log().Fatal(s.Run("7999", handler.InitRouter()))
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	logger.Log().Warn("graceful shutdown.")
}
