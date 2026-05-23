package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"posts/internal/config"
	"posts/internal/infrastructure/postgres"
	httpTransport "posts/internal/transport/http"
	"posts/internal/transport/http/handler"
	postusecase "posts/internal/usecases/post"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoadConfig()

	dbPool, err := postgres.NewConnection(cfg)
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}
	log.Printf("connected to database")

	postRepo := postgres.NewPostRepository(dbPool)

	createUC := postusecase.NewCreateUseCase(postRepo)
	getUC := postusecase.NewGetUseCase(postRepo)
	updateUC := postusecase.NewUpdateUseCase(postRepo)
	deleteUC := postusecase.NewDeleteUseCase(postRepo)
	listUC := postusecase.NewListUseCase(postRepo)

	postHandler := handler.NewPostHandler(
		createUC,
		getUC,
		updateUC,
		deleteUC,
		listUC,
	)

	router := httpTransport.NewRouter(postHandler)

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
