package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/application"
	"github.com/leetcode-golang-classroom/golang-gin-with-dynamic-json/internal/config"
)

func main() {
	app := application.New(config.AppConfig)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer func() {
		log.Println("gin-json-http stoping")
		cancel()
	}()
	err := app.Start(ctx)
	if err != nil {
		log.Println("failed to start gin-json-http:", err)
	}
}
