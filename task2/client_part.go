package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"task2/api"
	"time"
)

const testTimeout = 15 * time.Second

func main() {
	time.Sleep(2 * time.Second)
	// sending requests by a client
	client, err := api.NewClientWithResponses("http://0.0.0.0:8070")
	if err != nil {
		slog.Debug(err.Error())
		return
	}
	ctxClient := context.Background()

	responseVersion, err := client.GetVersionWithResponse(ctxClient)
	if err != nil {
		slog.Debug(err.Error())
		return
	}
	version := responseVersion.JSON200
	if version != nil {
		fmt.Println(version.Version)
	}

	var input api.Input
	input.Input = "SGVsbG8sIHdvcmxkIQ=="
	responseDecode, err := client.DecodeWithResponse(ctxClient, func(ctx context.Context, req *http.Request) error {
		var buffer bytes.Buffer
		if err := json.NewEncoder(&buffer).Encode(&input); err != nil {
			slog.Debug(err.Error())
			return err
		}
		req.Body = io.NopCloser(&buffer)
		return nil
	})
	if err != nil {
		slog.Debug(err.Error())
		return
	}
	output := responseDecode.JSON200
	if output != nil {
		fmt.Println(output.Output)
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	responseHardOp, err := client.RandomShitWithResponse(ctxTimeout)
	if err != nil {
		slog.Debug(err.Error())
	}
	if responseHardOp == nil || responseHardOp.StatusCode() == http.StatusBadRequest {
		fmt.Print("false")
	} else {
		fmt.Print("true ", responseHardOp.StatusCode)
	}
}
