package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/gauravgahlot/syncroot/internal/config"
	"github.com/gauravgahlot/syncroot/internal/enqueuer"
	"github.com/gauravgahlot/syncroot/internal/handlers"
	"github.com/gauravgahlot/syncroot/internal/middleware"
)

// Command returns the command to start the HTTP server.
func Command(logger *zap.Logger, cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start the syncroot HTTP server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runE(cmd.Context(), logger, cfg)
		},
	}
}

func runE(ctx context.Context, logger *zap.Logger, cfg *config.Config) error {
	router, err := readSpecAndCreateRouter(ctx)
	if err != nil {
		return fmt.Errorf("failed to create router: %w", err)
	}

	en, err := enqueuer.NewEnqueuer(logger)
	if err != nil {
		return fmt.Errorf("failed to initialize enqueuer: %w", err)
	}

	mux := http.NewServeMux()
	contactHandler := handlers.NewContactHandler(logger, en, cfg.Forwarder.Topic)

	// register API routes
	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("POST /contacts", contactHandler.CreateContact)
	mux.HandleFunc("GET /contacts/{id}", contactHandler.GetContact)
	mux.HandleFunc("PUT /contacts/{id}", contactHandler.UpdateContact)
	mux.HandleFunc("DELETE /contacts/{id}", contactHandler.DeleteContact)

	oapiValidator := middleware.NewOpenAPIValidator(*router, logger, cfg.Environment)
	wrappedMux := middleware.Logger(
		oapiValidator.Validate(mux),
		logger,
	)

	// create an HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: wrappedMux,
	}

	// start the server
	go func() {
		logger.Info("starting HTTP server", zap.Int("port", cfg.Server.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("failed to start the server", zap.Error(err))

			os.Exit(1)
		}
	}()

	// wait for interrupt signal
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	logger.Info("shutting down server...")

	// create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	logger.Info("server exited gracefully")

	return nil
}
