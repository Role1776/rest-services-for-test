package main

import (
	"app"
	"app/internal/handler"
	"app/internal/repository"
	"app/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := app.InitServer()
	db, err := repository.NewConn()
	if err != nil {
		logger.Log().Fatal(err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	handler := handler.NewHandler(repo)

	go func() {
		logger.Log().Fatal(s.Run("8000", handler.InitRouter()))
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	logger.Log().Warn("graceful shutdown.")
}
