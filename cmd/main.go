package main

import (
	"context"
	"fmt"
	"github/MahfujulSagor/video_downloader/internal/config"
	"github/MahfujulSagor/video_downloader/internal/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//? Setup Config
	cfg := config.MustLoad()

	//? Setup Logger
	logger.Init(cfg)

	//? Setup Server
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleRoot)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: mux,
	}

	//? Start server and listen for shutdown signal
	logger.Info.Println("Server started on:", server.Addr)
	var done = make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Error.Fatal("Failed to start server", err)
		}
	}()
	<-done

	logger.Info.Println("Server shutting done...")

	//? Shutdown server gracefully within 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error.Fatal("Server forced to shutdown", err)
	}
	logger.Info.Println("Server shut down gracefully")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}
