package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Korisss/concet-backend/internal/repository"
	"github.com/Korisss/concet-backend/internal/repository/psql"
	"github.com/Korisss/concet-backend/internal/service"
	"github.com/Korisss/concet-backend/internal/transport"
	"github.com/Korisss/concet-backend/internal/transport/handler"
	configuration "github.com/Korisss/concet-backend/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	config := configuration.Load("./configs/config.json")
	if config.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Error("error when loading .env file")
	}

	db, err := repository.NewPostgresDB(psql.Config{
		Host:     config.DBConfig.Host,
		Port:     config.DBConfig.Port,
		Username: config.DBConfig.Username,
		DBName:   config.DBConfig.DBName,
		SSLMode:  config.DBConfig.SSLMode,
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	server := new(transport.Server)

	go func() {
		if err := server.Run(strconv.Itoa(config.Port), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Concet Backend started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Concet Backend shutting down...")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Error("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Error("error occured on db connection close: %s", err.Error())
	}
}
