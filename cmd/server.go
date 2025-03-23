package main

import (
	"context"
	"fmt"
	"go-book/middlewares"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func isPortAvailable(port string) error {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	ln.Close()
	return nil
}

func StartServer(ctx context.Context, router *mux.Router) error {
	port := getPort()

	// Проверяем доступность порта
	if err := isPortAvailable(port); err != nil {
		return fmt.Errorf("port %s is not available: %w", port, err)
	}

	server := &http.Server{
		Addr:    port,
		Handler: middlewares.LoggingMiddleware(router),
	}

	// Канал для отслеживания ошибок запуска сервера
	serverErr := make(chan error, 1)

	go func() {
		log.Printf("HTTP server is listening on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
			serverErr <- err
		}
	}()

	// Ожидаем либо сигнал завершения, либо ошибку запуска
	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %w", err)
	case <-ctx.Done():
		log.Println("Initiating graceful shutdown...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("error during server shutdown: %w", err)
		}
		log.Println("Server has been gracefully stopped")
		return nil
	}
}
