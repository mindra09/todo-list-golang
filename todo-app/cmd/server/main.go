package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-app/internal/handler"
	"todo-app/internal/middleware"
	"todo-app/internal/repository"
	"todo-app/internal/service"
	"todo-app/internal/usecase"

	"github.com/gorilla/mux"
)

func main() {
	// ---------- Logger ----------
	logger := log.New(os.Stdout, "[TODO-APP] ", log.LstdFlags|log.Lshortfile)

	logger.Println("starting application...")

	// ---------- Dependency Injection ----------
	repo := repository.NewInMemoryTodoRepo()
	ai := service.NewAIClient()
	uc := usecase.NewTodoUsecase(repo, ai)
	h := handler.NewTodoHandler(uc)
	fe := handler.NewFrontEndHandler()

	// ---------- Initialize Router ----------
	r := mux.NewRouter()
	// ---------- Cors ----------
	r.Use(middleware.CORS)
	// ---------- Logging Api ----------
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("%s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	// Serve static files
	r.PathPrefix("/frontend/").Handler(http.StripPrefix("/frontend/",
		http.FileServer(http.Dir("./frontend"))))
	// ---------- Router ----------

	// ---------- Front End ----------
	r.HandleFunc("/", fe.Index).Methods("GET")

	// ---------- API ----------
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/todos", h.GetTodos).Methods("GET")
	api.HandleFunc("/todos", h.CreateTodo).Methods("POST")
	api.HandleFunc("/todos/{id}", h.DeleteTodo).Methods("DELETE")

	// ---------- HTTP Server ----------
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// ---------- Run Server ----------
	go func() {
		logger.Println("server running on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen error: %v", err)
		}
	}()

	// ---------- Graceful Shutdown ----------
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Println("shutdown signal received...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("server forced to shutdown: %v", err)
	}

	logger.Println("server exited gracefully")
}
