package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/saksham365/students-api/internal/config"
	"github.com/saksham365/students-api/internal/http/handlers/student"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// DB setup 
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// setup server
	server := http.Server{
		Addr: cfg.Address,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)

	// read about these, when they get triggered
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("Failed to start Server")
		}
	} ()

	<- done 

	slog.Info("shutting down the server")

	// read about context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("filed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")

}
