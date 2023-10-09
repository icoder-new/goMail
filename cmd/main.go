package main

import (
	"context"
	"fmt"
	gomail "fr33d0mz/goMail"
	"fr33d0mz/goMail/db"
	"fr33d0mz/goMail/logger"
	"fr33d0mz/goMail/pkg/handler"
	"fr33d0mz/goMail/pkg/repository"
	"fr33d0mz/goMail/pkg/service"
	"fr33d0mz/goMail/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	utils.ReadSettings()
	utils.PutAdditionalSettings()

	logger.Init()

	srv := new(gomail.Server)

	db.StartDbConnection()

	_db := db.GetDBConn()
	// if err := _db.AutoMigrate(&models.User{}); err != nil {
	// 	logger.Error.Fatal("failed to migrate tables")
	// }

	repository := repository.NewRepository(_db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	go func() {
		if err := srv.Run(utils.AppSettings.AppParams.PortRun, handler.InitRoutes()); err != nil {
			logger.Error.Fatal("Error occurred while running http server. Error is: ", err.Error())
			return
		}
	}()

	fmt.Println("SERVER IS UP [localhost:3333]")
	fmt.Println("waiting for signal [0w0]...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("Server is closed")

	db.DisconnectDB(_db)
	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error.Fatal("Error occurred on server shutting down. Error is: ", err.Error())
		return
	}
}
