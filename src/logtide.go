package main

import (
	"context"
	"log"
	"os"

	"github.com/logtide-dev/logtide-sdk-go"
)

func LogtideStart() {
	key, isPresent := os.LookupEnv("LOGTIDE_KEY")
	if !isPresent {
		log.Panic("No logtide key")
		os.Exit(0)
	}
	client, err := logtide.New(
		logtide.WithAPIKey(key),
		logtide.WithService("EDAM DEV"),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()

	// Send logs
	client.Info(ctx, "Server started", map[string]any{"port": 8080})
	//   client.Error(ctx, "Connection failed", map[string]any{"error": "timeout"})
}
