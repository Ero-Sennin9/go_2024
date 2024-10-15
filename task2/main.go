package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task2/Serialization"
	"task2/Server"
	"task2/api"
	"time"
)

const shutdownTimeout = 20 * time.Second
const testTimeout = 15 * time.Second

func main() {
	// getting server info
	file, errOpen := os.Open("api/api.openapi.yaml")
	if errOpen != nil {
		slog.Debug(errOpen.Error())
		return
	}
	var servers Serialization.Servers
	if err := yaml.NewDecoder(file).Decode(&servers); err != nil {
		slog.Debug(err.Error())
		return
	}
	defer file.Close()

	serverLogic := Server.ServerLogic{}

	mux := http.NewServeMux()

	handler := api.HandlerFromMux(&serverLogic, mux)

	if len(servers.Info) == 0 {
		err := errors.New("Unable to load address from openapi specification")
		slog.Debug(err.Error())
		return
	}

	address := servers.Info[0].Url
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

	group.Go(func() error {
		time.Sleep(2 * time.Second)
		// sending requests by a client
		client, err := api.NewClient("http://" + address)
		if err != nil {
			slog.Debug(err.Error())
			return err
		}
		ctxClient := context.Background()
		responseVersion, err := client.GetVersion(ctxClient)
		if responseVersion == nil {
			err := errors.New("Unable to get version")
			slog.Debug(err.Error())
			return err
		}
		if err != nil {
			slog.Debug(err.Error())
			return err
		}
		var version api.Version
		if err := json.NewDecoder(responseVersion.Body).Decode(&version); err != nil {
			slog.Debug(err.Error())
			return err
		}
		fmt.Println(*version.Version)

		var input api.Input
		input.Input = new(string)
		*input.Input = "SGVsbG8sIHdvcmxkIQ=="
		responseDecode, err := client.Decode(ctxClient, func(ctx context.Context, req *http.Request) error {
			var buffer bytes.Buffer
			if err := json.NewEncoder(&buffer).Encode(&input); err != nil {
				slog.Debug(err.Error())
				return err
			}
			req.Body = io.NopCloser(&buffer)
			return nil
		})
		var output api.Output
		if err := json.NewDecoder(responseDecode.Body).Decode(&output); err != nil {
			slog.Debug(err.Error())
			return err
		}
		fmt.Println(*output.Output)

		ctxTimeout, cancel := context.WithTimeout(ctx, testTimeout)
		defer cancel()

		responseHardOp, err := client.RandomShit(ctxTimeout)
		if err != nil {
			slog.Debug(err.Error())
		}

		if responseHardOp == nil || responseHardOp.StatusCode == http.StatusBadRequest {
			fmt.Print("false")
		} else {
			fmt.Print("true ", responseHardOp.StatusCode)
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		slog.Debug(err.Error())
		return
	}
}
