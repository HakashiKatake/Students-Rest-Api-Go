package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HakashiKatake/Students-Rest-Api-Go/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//db setup

	//setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Student API"))
	})

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Println("server started on ", cfg.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server: ", err)
		}
	}()

	<-done

	slog.Info("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server: ", slog.String("error", err.Error()))
	} else {
		slog.Info("server shutdown gracefully")
	}

}
