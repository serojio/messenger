package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"users/internal/config"
	"users/internal/infrastructure/postgres"
	httpTransport "users/internal/transport/http"
	"users/internal/transport/http/handler"
	userusecase "users/internal/usecases/user"
)

func main() {
	cfg := config.MustLoadConfig()

	dbPool, err := postgres.NewConnection(cfg)
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}
	log.Printf("connected to database")

	userRepo := postgres.NewUserRepository(dbPool)

	createUC := userusecase.NewCreateUseCase(userRepo)

	userHandler := handler.NewUserHandler(
		createUC,
		nil, // getUC
		nil, // updateUC
		nil, // deleteUC
	)

	router := httpTransport.NewRouter(userHandler)

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: router,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		log.Printf("server started on :%s", cfg.HTTPPort)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		15*time.Second,
	)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown error: %v", err)
	}

	dbPool.Close()

	log.Println("server stopped")
}
