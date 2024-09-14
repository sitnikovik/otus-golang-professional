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
	args, flags, err := ParseInput(os.Args)
	if err != nil {
		log.Fatalf("parse input err: %v", err)
	}

	client := NewTelnetClient(args.Address, args.Port, time.Duration(flags.Timeout), os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer client.Close()

	// Создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// Ловим системные сигналы (Ctrl+C) и EOF (Ctrl+D)
	go handleSignals(cancel)

	// Запускаем горутины для отправки и получения данных
	go sendRoutine(ctx, client)
	go receiveRoutine(ctx, client)

	// Ожидаем завершения работы по сигналу
	<-ctx.Done()
	fmt.Fprintln(os.Stderr, "Connection closed")
}

func handleSignals(cancel context.CancelFunc) {
	sigsCh := make(chan os.Signal, 1)
	signal.Notify(sigsCh, syscall.SIGINT, syscall.SIGTERM)

	// Ожидаем сигнала
	<-sigsCh
	fmt.Fprintln(os.Stderr, "Received interrupt signal, shutting down...")
	cancel()
}

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
