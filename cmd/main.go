package main

import (
	"context"
	"go-book/pkg/db"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Проверяем доступность порта перед инициализацией других ресурсов
	port := getPort()
	if err := isPortAvailable(port); err != nil {
		log.Fatalf("Port %s is not available: %v", port, err)
	}

	// Создаем корневой контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Канал для обработки сигналов завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Горутина для обработки сигналов
	go func() {
		sig := <-sigChan
		log.Printf("Received shutdown signal: %v", sig)
		cancel() // Отменяем контекст при получении сигнала
	}()

	// Инициализируем базу данных
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func() {
		db.CloseDB()
		log.Println("All resources have been cleaned up")
	}()

	// Регистрируем маршруты
	mux, err := RegisterRoutes()
	if err != nil {
		log.Fatalf("Failed to register routes: %v", err)
	}

	log.Printf("Starting application on port %s", port)

	// Запускаем сервер
	if err := StartServer(ctx, mux); err != nil {
		log.Printf("Server error: %v", err)
	}
}
