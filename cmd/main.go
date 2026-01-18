package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {
	// созданеи логгера
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	// создание сервера
	srv := server.NewServer(logger)
	logger.Println("Сервер запущен")
	logger.Fatal(srv.HTTP.ListenAndServe())
}
