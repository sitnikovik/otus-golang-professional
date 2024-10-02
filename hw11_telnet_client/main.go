package main

import (
	"context"
	"errors"
	"fmt"
	"io"
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
	defer cancel()
	cancelCh := make(chan struct{})

	// Горутина для обработки сигнала отмены
	go func() {
		<-cancelCh
		cancel()
	}()

	// Ловим системные сигналы для корректного завершения работы
	go handleSignals(cancelCh)

	// Запускаем горутины для отправки и получения данных
	go sendRoutine(ctx, cancelCh, client)
	go receiveRoutine(ctx, cancelCh, client)

	// Ожидаем завершения работы по сигналу
	<-ctx.Done()
	fmt.Fprintln(os.Stderr, "Connection closed")
}

// handleSignals обрабатывает системные сигналы для корректного завершения работы.
func handleSignals(cancelCh chan struct{}) {
	sigsCh := make(chan os.Signal, 1)
	signal.Notify(sigsCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigsCh
	fmt.Fprintln(os.Stderr, "Received interrupt signal, shutting down...")
	cancelCh <- struct{}{}
}

// sendRoutine запускает рутину для отправки данных на телнет-сервер.
func sendRoutine(ctx context.Context, cancelCh chan struct{}, client TelnetClient) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := client.Send(); err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Fprintln(os.Stderr, "EOF detected, closing connection..")
					cancelCh <- struct{}{}
					return
				}
				fmt.Fprintln(os.Stderr, "Send error:", err)
				return
			}
		}
	}
}

// receiveRoutine запускает рутину для получения данных от телнет-сервера.
func receiveRoutine(ctx context.Context, cancelCh chan struct{}, client TelnetClient) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := client.Receive(); err != nil {
				fmt.Fprintln(os.Stderr, "Receive error:", err)
				cancelCh <- struct{}{}
				return
			}
		}
	}
}
