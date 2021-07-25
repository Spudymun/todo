package main

import (
	"os"

	"github.com/Spudymun/todo"
	"github.com/Spudymun/todo/pkg/handler"
	"github.com/Spudymun/todo/pkg/repository"
	"github.com/Spudymun/todo/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Задания формата JSON для логов
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Инициализация конфигов. Ф-ция viper.ReadInConfig() возвращает только ошибку
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// Инициализация БД с передачей всех необходимых знаяений из вайпера и переменной окружения
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	// Обьявление всех зависимостей в нужном порядке
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// Инициализация экземпляра сервера с помощью new()
	srv := new(todo.Server)
	//Запуск сервера и передача значения порта из вайпера по ключу
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while runnng httpserver: %s", err.Error())
	}
}

// Ф-ция инициализации конф. файлов
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	// Возырат ф-ции котрая считывает конфиги и записывает их во внутренний обьек вайпера
	return viper.ReadInConfig()
}
