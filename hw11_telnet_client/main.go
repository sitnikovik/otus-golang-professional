package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Парсим входные аргументы и флаги
	args, flags, err := ParseInput(os.Args)
	if err != nil {
		log.Fatalf("parse input err: %v", err)
	}

	// Создаем клиент и подключаемся к серверу
	addr := fmt.Sprintf("%s:%d", args.Address, args.Port)
	client := NewTelnetClient(addr, time.Duration(flags.Timeout)*time.Second, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithCancel(context.Background())

	// Ловим системные сигналы для корректного завершения работы
	go handleSignals(cancel)

	// Запускаем горутины для отправки и получения данных
	go sendRoutine(ctx, client)
	go receiveRoutine(ctx, client)

	// Ожидаем завершения работы по сигналу
	<-ctx.Done()
	fmt.Fprintln(os.Stderr, "Connection closed")
}

// handleSignals обрабатывает системные сигналы для корректного завершения работы.
func handleSignals(cancel context.CancelFunc) {
	sigsCh := make(chan os.Signal, 1)
	signal.Notify(sigsCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigsCh
	fmt.Fprintln(os.Stderr, "Received interrupt signal, shutting down...")
	cancel()
}

// sendRoutine запускает рутину для отправки данных на телнет-сервер.
func sendRoutine(ctx context.Context, client TelnetClient) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := client.Send(); err != nil {
				fmt.Fprintln(os.Stderr, "Send error:", err)
				return
			}
		}
	}
}

// receiveRoutine запускает рутину для получения данных от телнет-сервера.
func receiveRoutine(ctx context.Context, client TelnetClient) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := client.Receive(); err != nil {
				fmt.Fprintln(os.Stderr, "Receive error:", err)
				return
			}
		}
	}
}
