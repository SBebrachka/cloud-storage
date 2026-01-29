package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sbebrachka/pet3"
	"github.com/sbebrachka/pet3/pkg/handler"
	"github.com/sbebrachka/pet3/pkg/repository"
	"github.com/sbebrachka/pet3/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal(err)
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal(err)
	}
	repos := repository.NewRepository(db) // указатель на сервисы
	services := service.NewService(repos) // конструктор с зависимостями сервиса
	handlers := handler.NewHandler(services)
	srv := new(pet3.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatal(err.Error())
	}
}

// Инициализация конфигационных файлов
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig() // считываем внутреннее значение и записываем в viper
}
