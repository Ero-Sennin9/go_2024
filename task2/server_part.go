package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"task2/Server"
	"task2/api"
	"time"
)

const shutdownTimeout = 20 * time.Second
const address = "0.0.0.0:8070"

func main() {
	serverLogic := Server.ServerLogic{}

	mux := http.NewServeMux()

	handler := api.HandlerFromMux(&serverLogic, mux)

	server := http.Server{
		Addr:    address,
		Handler: handler,
	}

	signalContext, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(signalContext)

	group.Go(func() error {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Debug(err.Error())
			return fmt.Errorf("failed to serve http server: %w", err)
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()

		contextShutdown, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(contextShutdown); err != nil {
			slog.Debug(err.Error())
			return err
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		slog.Debug(err.Error())
		return
	}
}
