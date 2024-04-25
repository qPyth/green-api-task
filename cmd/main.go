package main

import (
	"fmt"
	"github.com/qPyth/green-api-task/internal/transport"
	usecases "github.com/qPyth/green-api-task/internal/usecase"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// init logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))

	// init use cases and handlers
	uc := usecases.NewGreenApiUC()
	h := transport.NewHandler(uc, uc, uc, uc, logger)

	//init router
	r := http.NewServeMux()

	//init file server
	r.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui"))))

	// set routes and endpoints
	r.HandleFunc("/", h.Index)
	r.HandleFunc("/settings", h.Settings)
	r.HandleFunc("/state-instance", h.StateInstance)
	r.HandleFunc("/send-message", h.SendMessage)
	r.HandleFunc("/send-file", h.SendFile)

	addr := fmt.Sprintf("localhost:%s", port)

	slog.Info("server starting...", "host", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
