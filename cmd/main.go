package main

import (
	"context"
	"github.com/EgorMizerov/portfolio"
	"github.com/EgorMizerov/portfolio/pkg/handler"
	"github.com/EgorMizerov/portfolio/pkg/repository"
	"github.com/EgorMizerov/portfolio/pkg/service"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Логгер
	logger := zap.NewExample()
	defer logger.Sync()

	// Конфиг
	if err := initConfig(); err != nil {
		logger.Error("error init configs", zap.String("error", err.Error()))
	}

	// База Данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		//Password: viper.GetString("db.password"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		logger.Error("error init database", zap.String("error", err.Error()))
	}

	newRepository := repository.NewRepository(db)
	services := service.NewService(newRepository)
	handlers := handler.NewHandler(services)

	server := new(portfolio.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logger.Info("Start listening server...")
	//fmt.Printf("%T\n", time.Now())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("Server shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Error("error occured on server shutting down", zap.String("error", err.Error()))
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	return viper.ReadInConfig()
}
