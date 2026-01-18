package server

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"log"
	"net/http"
	"time"
)

// создание структуры сервера с полями для логгера (log.Logger) и http-сервера (http.Server)
type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

// http-роутер
func NewServer(logger *log.Logger) *Server {
	// регистраия хендлеров в http-роутер
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleMain)
	mux.HandleFunc("/upload", handlers.HandleUpload)

	// создание экземпляра структуры http.Server
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,              // http-роутер
		ErrorLog:     logger,           // логгер
		ReadTimeout:  5 * time.Second,  // таймаут для чтения
		WriteTimeout: 10 * time.Second, // таймаут для записи
		IdleTimeout:  15 * time.Second, // таймаут ожидания следующего запроса
	}

	// ссылка на сервер
	return &Server{
		Logger: logger,
		HTTP:   httpServer,
	}
}
