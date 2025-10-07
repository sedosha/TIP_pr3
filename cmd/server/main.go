package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pz3-http/internal/api"
	"pz3-http/internal/storage"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	store := storage.NewMemoryStore()
	h := api.NewHandlers(store)

	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		api.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /tasks", h.ListTasks)
	mux.HandleFunc("POST /tasks", h.CreateTask)
	mux.HandleFunc("GET /tasks/", h.GetTask)
	mux.HandleFunc("PATCH /tasks/{id}", h.UpdateTask)
	mux.HandleFunc("PATCH /tasks/{id}/done", h.MarkTaskDone)
	mux.HandleFunc("DELETE /tasks/{id}", h.DeleteTask)

	handler := api.CORS(api.Logging(mux))

	server := &http.Server{
		Addr:    ":" + port,  // Используем порт из переменной
		Handler: handler,
	}

	go func() {
		log.Println("Server starting on port", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error:", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown error:", err)
	}

	log.Println("Server stopped gracefully")
}