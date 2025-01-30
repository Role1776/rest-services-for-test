package main

import (
	"app"
	"app/internal/handler"
	"app/internal/service"
	"app/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := app.InitServer()

	service := service.NewService()
	handler := handler.NewHandler(service)

	go func() {
		if err := s.Run("8001", handler.InitRouter()); err != nil {
			logger.Log().Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	logger.Log().Warn("graceful shutdown.")
}
